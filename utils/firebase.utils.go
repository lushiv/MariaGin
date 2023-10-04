package common_utils

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var FirebaseConnection *firebase.App

// InitializeFirebaseApp initializes the Firebase Admin SDK app.
func InitializeFirebaseApp(serviceAccountFile string) error {
	ctx := context.Background()
	opt := option.WithCredentialsFile(serviceAccountFile)
	var err error
	FirebaseConnection, err = firebase.NewApp(ctx, nil, opt)
	return err
}

// UploadToFirebase uploads a file to Firebase Storage using an existing Firebase app.
func UploadToFirebase(app *firebase.App, file *multipart.FileHeader, objectKey string) (string, error) {
	fmt.Println(objectKey)
	ctx := context.Background()
	client, err := app.Storage(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to initialize Firebase Storage client: %v", err)
	}

	bucketName := "eco-friendly-tourism.appspot.com" // Specify the bucket name here

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return "", fmt.Errorf("failed to get Firebase Storage bucket: %v", err)
	}

	// // Check if the object already exists in the bucket
	// if _, err := bucket.Object(objectKey).Attrs(ctx); err != nil {
	// 	if err == storage.ErrObjectNotExist {
	// 		return "", fmt.Errorf("object does not exist: %v", err)
	// 	} else {
	// 		return "", fmt.Errorf("failed to get object attributes: %v", err)
	// 	}
	// }

	wc := bucket.Object(objectKey).NewWriter(ctx)
	defer wc.Close()

	fileData, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer fileData.Close()

	if _, err = io.Copy(wc, fileData); err != nil {
		return "", fmt.Errorf("failed to copy file to Firebase Storage: %v", err)
	}

	// Set the public access permissions for the uploaded file
	if err := bucket.Object(objectKey).ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return "", fmt.Errorf("failed to set ACL for the uploaded file: %v", err)
	}

	// Generate and return the public URL of the uploaded file
	publicURL := "https://firebasestorage.googleapis.com/v0/b/" + bucketName + "/o/" + objectKey + "?alt=media"
	return publicURL, nil
}
