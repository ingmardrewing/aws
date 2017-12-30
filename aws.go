package aws

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadFile(path string, bucket string, key string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln("Unable to open file for upload", err)
	}
	defer file.Close()
	log.Printf("File found at %s, starting aws communication\n", path)
	log.Printf("with bucket %s and key %s\n", bucket, key)

	sess, _ := session.NewSession()
	log.Println("AWS session initiated")
	uploader := s3manager.NewUploader(sess)
	uploadInput := &s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   file,
		ACL:    aws.String("public-read"),
	}

	result, err := uploader.Upload(uploadInput)

	if err != nil {
		log.Fatalln("Unable to upload file", err)
	}

	log.Printf("Successfully uploaded %q to %q\n", path, result.Location)
	return result.Location
}
