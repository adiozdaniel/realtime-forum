package shared

import (
	"errors"
	"io"
	"net/http"
	"os"
)

// SaveImage handles file uploads and saves the image to a specified directory.
// It returns the image path or an error.
func (s *SharedConfig) SaveImage(r *http.Request, fileName string) (string, error) {
	err := r.ParseMultipartForm(20 << 20) // 20MB max file size
	if err != nil {
		return "", errors.New("failed to parse form data")
	}

	var imagePath string

	// Handle file upload
	file, _, err := r.FormFile("image")
	if err == nil {
		defer file.Close()

		// Save the file to a directory profile
		imagePath = "./static/profiles/" + fileName
		dst, err := os.Create(imagePath)
		if err != nil {
			return "", errors.New("failed to save image")
		}
		defer dst.Close()

		// Copy the uploaded file to the destination
		_, err = io.Copy(dst, file)
		if err != nil {
			return "", errors.New("failed to save image")
		}
	} else if !errors.Is(err, http.ErrMissingFile) {
		return "", errors.New("image upload failed")
	}

	return imagePath[1:], nil
}

// DeletePostImage removes the image associated with a post.
func (s *SharedConfig) DeletePostImage(image string) error {
	if err := os.Remove("./static/profiles/" + image); err != nil {
		return errors.New("failed to delete post image")
	}

	return nil
}
