package common_utils

import (
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func UploadToS3(file *multipart.FileHeader, bucketName, objectKey string) (string, error) {
	// Read AWS credentials from environment variables
	awsAccessKeyID := os.Getenv("AWS_S3_KEY")
	awsSecretAccessKey := os.Getenv("AWS_S3_SECRET")
	awsRegion := os.Getenv("AWS_S3_REGION")

	// Initialize AWS session and S3 client
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: credentials.NewStaticCredentials(awsAccessKeyID, awsSecretAccessKey, ""),
	})

	if err != nil {
		return "", err
	}

	svc := s3.New(sess)

	// Open the uploaded file
	fileData, err := file.Open()
	if err != nil {
		return "", err
	}
	defer fileData.Close()

	// Create S3 uploader input parameters
	params := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   fileData,
	}

	// Upload the file to S3
	_, err = svc.PutObject(params)
	if err != nil {
		return "", err
	}

	// Generate and return the public URL of the uploaded file
	publicURL := "https://s3.amazonaws.com/" + bucketName + "/" + objectKey
	return publicURL, nil
}
