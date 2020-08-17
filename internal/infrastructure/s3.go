package infrastructure

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func NewS3Client() *s3.S3 {
	return s3.New(session.New(), &aws.Config{
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("S3_ACCESS_KEY"),
			os.Getenv("S3_SECRET_KEY"),
			"",
		),
		Region:   aws.String(os.Getenv("S3_REGION")),
		Endpoint: aws.String(os.Getenv("S3_ENDPOINT")),
	})

}
