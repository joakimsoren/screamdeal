package awsutils

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
)

func getS3Session(region string) *session.Session {
	return session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewEnvCredentials(),
	}))
}

// GetThemeURLs Get all theme links from s3
func GetThemeURLs() []string {
	themeURLs := []string{
		"https://scary-bucket.s3-eu-west-1.amazonaws.com/images/71-512.png",
		"https://scary-bucket.s3-eu-west-1.amazonaws.com/images/84-840138_alien-isolation-alien-figure-hd-png-download.png",
		"https://scary-bucket.s3-eu-west-1.amazonaws.com/images/Ghost-Download-PNG.png",
		"https://scary-bucket.s3-eu-west-1.amazonaws.com/images/death-2026516__340.png",
		"https://scary-bucket.s3-eu-west-1.amazonaws.com/images/scream-face-png-1.png",
		"https://scary-bucket.s3-eu-west-1.amazonaws.com/images/zombie_PNG56.png",
	}
	return themeURLs
}

// UploadFile Upload a document and get its id
func UploadFile(file *os.File) string {
	documentID := uuid.New().String()

	bucket := "scary-bucket"

	key := fmt.Sprintf("/original-documents/%s.pdf", documentID)

	sess := getS3Session("eu-west-1")
	uploader := s3manager.NewUploader(sess)

	upParams := &s3manager.UploadInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   file,
	}

	_, uploadErr := uploader.Upload(upParams)
	if uploadErr != nil {
		return "Error uploading file"
	}

	return documentID
}

// DownloadFile Download a file from s3, returns local file path
func DownloadFile(filename, bucket string) string {
	sess := getS3Session("eu-west-1")
	downloader := s3manager.NewDownloader(sess)

	key := fmt.Sprintf("/original-documents/%s", filename)
	localFilePath := fmt.Sprintf("./%s", filename)

	file, err := os.Create(localFilePath)

	_, err = downloader.Download(file, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		fmt.Printf("Failed to download file, %v", err)
	}
	return localFilePath
}
