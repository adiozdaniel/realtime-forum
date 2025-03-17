package shared

import (
	"crypto/rand"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

// GenerateUUID creates a cryptographically secure random token
func (s *SharedConfig) GenerateUUID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	// Format it as a UUID-like string
	token := fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])

	return token, nil
}

// Helper function to convert string to sql.NullString
func (s *SharedConfig) ToNullString(str string) sql.NullString {
	if str == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: str, Valid: true}
}

// Helper function to trims spaces
func (s *SharedConfig) SanitizeString(value string) string {
	return strings.TrimSpace(value)
}

// Helper function that sanitizes all fields in a struct recursively
func (s *SharedConfig) SanitizeInput(input any) (any, error) {
	v := reflect.ValueOf(input)

	// Ensure input is a pointer to a struct
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return nil, errors.New("bad request")
	}

	v = v.Elem()
	t := v.Type()

	// Iterate over struct fields
	for i := range v.NumField() {
		field := v.Field(i)
		fieldType := t.Field(i)

		// Skip unexported fields
		if !field.CanSet() {
			continue
		}

		switch field.Kind() {
		case reflect.String:
			field.SetString(s.SanitizeString(field.String()))

		case reflect.Struct:
			// Special case for sql.NullString and similar types
			if fieldType.Type == reflect.TypeOf(sql.NullString{}) {
				ns := field.Interface().(sql.NullString)
				ns.String = s.SanitizeString(ns.String)
				field.Set(reflect.ValueOf(ns))
				continue
			}

			// Recursively sanitize nested structs
			newStruct, err := s.SanitizeInput(field.Addr().Interface())
			if err != nil {
				return nil, err
			}
			field.Set(reflect.ValueOf(newStruct).Elem())

		case reflect.Slice:
			elemType := field.Type().Elem()
			if elemType.Kind() == reflect.Ptr && elemType.Elem().Kind() == reflect.Struct {
				for j := range field.Len() {
					elem := field.Index(j)
					if elem.IsNil() {
						continue
					}
					newStruct, err := s.SanitizeInput(elem.Interface())
					if err != nil {
						return nil, err
					}
					elem.Set(reflect.ValueOf(newStruct))
				}
			}

		case reflect.Ptr:
			if field.IsNil() {
				continue
			}
			if field.Elem().Kind() == reflect.Struct {
				newStruct, err := s.SanitizeInput(field.Interface())
				if err != nil {
					return nil, err
				}
				field.Set(reflect.ValueOf(newStruct))
			}
		}
	}
	return input, nil
}

// Helper function to sanitize name
func (s *SharedConfig) CleanUsername(name string) string {
	name = strings.TrimSpace(name)
	re := regexp.MustCompile(`[^a-zA-Z]`)
	return s.SanitizeString(re.ReplaceAllString(name, " "))
}
