package helpers

import (
	"backend/config"
	"context"
	"errors"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"mime/multipart"
	"path/filepath"
)

// UploadPhoto uploads a file to Google Drive and returns the URL to access it
func UploadPhoto(file *multipart.FileHeader, folderID string) (string, error) {
	// Create a unique filename for the upload
	fileExt := filepath.Ext(file.Filename)
	filename := uuid.New().String() + fileExt

	// Get the file content
	src, err := file.Open()
	if err != nil {
		log.Error("Error opening upload file: ", err)
		return "", errors.New("could not open uploaded file")
	}
	defer src.Close()

	// Create a Google Drive client
	client := config.ServiceAccount("config/client_secret.json")
	srv, err := drive.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Error("Unable to create Drive service: ", err)
		return "", errors.New("failed to connect to storage service")
	}

	// Create file metadata
	f := &drive.File{
		Name:     filename,
		MimeType: file.Header.Get("Content-Type"),
		Parents:  []string{folderID},
	}

	// Upload the file
	uploadedFile, err := srv.Files.Create(f).Media(src).Do()
	if err != nil {
		log.Error("Unable to upload file to Drive: ", err)
		return "", errors.New("failed to upload the image")
	}

	// Set permission to make the file publicly accessible
	permission := &drive.Permission{
		Type: "anyone",
		Role: "reader",
	}
	_, err = srv.Permissions.Create(uploadedFile.Id, permission).Do()
	if err != nil {
		log.Error("Unable to set file permission: ", err)
		return "", errors.New("failed to set file permissions")
	}

	// Construct and return the file's public URL
	fileURL := "https://drive.google.com/uc?export=view&id=" + uploadedFile.Id
	return fileURL, nil
}
