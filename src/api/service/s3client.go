package service

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

func createS3Session() *session.Session{
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String(endpoints.UsEast2RegionID),
		Credentials: credentials.NewStaticCredentials(*getAccessKey(), *getSecretKey(), ""),
	})
	return sess
}

func Upload(filename *string, path *string) (error){
	uploader := s3manager.NewUploader(createS3Session())
	f, err  := os.Open(*path)
	if err != nil {
		return fmt.Errorf("failed to open file %q, %v", filename, err)
	}

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("imagerekognitionai"),
		Key:    aws.String(*filename),
		Body:   f,
	})
	if err != nil {
		return fmt.Errorf("failed to upload file, %v", err)
	}
	fmt.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location))
	return nil

}


