package shared

import (
	"errors"
	"io"
	"net/http"
	"os"
)

// SaveImage handles file uploads and saves the image to a specified directory.
// It returns the image path or an error.
func (s *SharedConfig) SaveImage(r *http.Request, fieldName, uploadDir string) (string, error) {
	err := r.ParseMultipartForm(20 << 20) // 20MB max file size
	if err != nil {
		return "", errors.New("failed to parse form data")
	}

	file, handler, err := r.FormFile(fieldName)
	if err != nil {
		if errors.Is(err, http.ErrMissingFile) {
			return "", nil // No file uploaded, not an error
		}
		return "", errors.New("image upload failed")
	}
	defer file.Close()

	// Ensure the upload directory exists
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return "", errors.New("failed to create upload directory")
	}

	// Construct the file path
	imagePath := uploadDir + handler.Filename
	dst, err := os.Create(imagePath)
	if err != nil {
		return "", errors.New("failed to save image")
	}
	defer dst.Close()

	// Copy the uploaded file to the destination
	if _, err = io.Copy(dst, file); err != nil {
		return "", errors.New("failed to save image")
	}

	return imagePath[1:], nil // Remove the leading './' for cleaner storage
}
