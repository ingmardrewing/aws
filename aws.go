// Simplify uploads to AWS
package aws

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// UploadFile uploads a file found at the given path to the given bucket
// with the given key and returns the location of the uploaded file
func UploadFile(path string, bucket string, key string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln("Unable to open file for upload", err)
	}
	defer file.Close()

	sess, err := session.NewSession()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("AWS session initiated:")
	log.Printf("Going to upload %s to bucket %s with key %s\n", path, bucket, key)

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
