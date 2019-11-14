package awsutils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
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
