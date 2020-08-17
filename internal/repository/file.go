package repository

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
)

func (r *repository) UploadCardFile(fileString string) (string, error) {
	log.Println(fileString)
	var result string
	fileType := fileString[strings.IndexByte(fileString, ':')+1 : strings.IndexByte(fileString, ';')]
	unbased, err := base64.StdEncoding.DecodeString(fileString[strings.IndexByte(fileString, ',')+1:])
	if err != nil {
		return "", err
	}
	fileName := sha1.Sum([]byte(unbased))
	switch fileType {
	case "image/png":
		result, err = r.s3Upload(bytes.NewReader(unbased), fmt.Sprintf("%x.png", fileName))
	case "image/jpeg":
		result, err = r.s3Upload(bytes.NewReader(unbased), fmt.Sprintf("%x.jpeg", fileName))
	default:
		return "", errors.New("invalid file type")
	}
	if err != nil {
		return "", err
	}
	return result, nil

}

func (r *repository) s3Upload(data io.ReadSeeker, fileName string) (string, error) {
	_, err := r.fileClient.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET_NAME")),
		Key:    aws.String(fileName),
		Body:   data,
		ACL:    aws.String(s3.BucketCannedACLPublicRead),
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s/"+fileName, os.Getenv("S3_ENDPOINT_EXTERNAL"), os.Getenv("S3_BUCKET_NAME")), nil

}
