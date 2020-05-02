package constants

import (
	"os"
)

var (
	// Port - Constant variable
	Port = getEnvironmentVariable("IMAGE_UPLOAD_PORT", ":8021")
	// AwsAccessKey - aws
	AwsAccessKey = getEnvironmentVariable("AWS_ACCESS_KEY", "")
	// AwsSecretKey - aws
	AwsSecretKey = getEnvironmentVariable("AWS_SECRET_KEY", "")
	// AwsBucketLink - aws
	AwsBucketLink = getEnvironmentVariable("AWS_BUCKET_LINK", "")
	// AwsBucketRegion - aws
	AwsBucketRegion = getEnvironmentVariable("AWS_BUCKET_REGION", "")
	// PostgreSQLDb - database
	PostgreSQLDb = getEnvironmentVariable("POSTGRES_SQL_DB", "")
	// SanicServer - server
	SanicServer = "http://sanic-server:8000"
	// ImageLinkS3 - s3
	ImageLinkS3 = getEnvironmentVariable("S3_LINK", "")
)

func getEnvironmentVariable(EnvName string, defaultValue string) string {
	value := os.Getenv(EnvName)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
