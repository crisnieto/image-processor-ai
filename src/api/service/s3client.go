package service

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
	"strings"
)

func createS3Session() *session.Session{
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String(endpoints.UsEast2RegionID),
		Credentials: credentials.NewStaticCredentials(*getAccessKey(), *getSecretKey(), ""),
	})
	return sess
}

func Upload(filename *string, path *string) (error){
	fmt.Println("Path del file que voy a subir: " + *path)
	uploader := s3manager.NewUploader(createS3Session())
	fmt.Println("CREÉ LA SESIÓN")
	f, err  := os.Open(*path)
	if err != nil {
		return fmt.Errorf("failed to open file %q, %v", filename, err)
	}
	fmt.Println("ENTRE 2")


	// Upload the file to S3.
	s3folder := strings.Split(*filename, ".")
	fmt.Println("ENTRE 3")

	fmt.Println("Saving filename: " + *filename)

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("imagerekognitionai"),
		Key:    aws.String(s3folder[0]+"/"+*filename),
		Body:   f,
	})
	if err != nil {
		return fmt.Errorf("failed to upload file, %v", err)
	}
	fmt.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location))
	return nil

}

func Download(filename *string) (*string,error) {
	downloader := s3manager.NewDownloader(createS3Session())
	f, err := os.Create(*filename)
	if err != nil {
		return nil, fmt.Errorf("failed to create file %q, %v", filename, err)
	}

	// Write the contents of S3 Object to the file

	s3folder := strings.Split(*filename, ".")

	n, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String("imagerekognitionai"),
		Key:    aws.String(s3folder[0]+"/"+*filename),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to download file, %v", err)
	}
	fmt.Printf("file downloaded, %d bytes\n", n)

	return filename, nil

}


