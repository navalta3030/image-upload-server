package controllers

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"image_upload_server/constants"
	"image_upload_server/models"
)

// ImageUploadDigitalOcean - image upload
func ImageUploadDigitalOcean(files []*multipart.FileHeader, userForm string) int {
	var userModel models.User

	// Initialize a client using spaces
	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(constants.AwsAccessKey, constants.AwsSecretKey, ""),
		Endpoint:    aws.String(constants.AwsBucketLink),
		Region:      aws.String(constants.AwsBucketRegion), // This is counter intuitive, but it will fail with a non-AWS region name.
	}
	newSession := session.New(s3Config)
	s3Client := s3.New(newSession)

	// Parse user
	errParse := json.Unmarshal([]byte(userForm), &userModel)
	if errParse != nil {
		logError.Fatal(errParse)
	}

	// Upload a file to the Space
	for _, file := range files {

		logMessage := fmt.Sprintf("Uploading File to digital ocean : %s", file.Filename)
		logger.Println(logMessage)

		// open and upload
		f, _ := file.Open()
		contentType := "image/" + filepath.Ext(file.Filename)[1:]
		object := s3.PutObjectInput{
			Body:        f,
			Key:         aws.String("image_public/" + userModel.Email + "_" + file.Filename),
			Bucket:      aws.String("xrayimages"),
			ContentType: aws.String(contentType),
			ACL:         aws.String("public-read"),
		}

		// check error
		_, err := s3Client.PutObject(&object)
		if err != nil {
			logError.Println("Unable to upload file.")
			logError.Println(err.Error())
			return 400
		}
	}

	return 200
}
