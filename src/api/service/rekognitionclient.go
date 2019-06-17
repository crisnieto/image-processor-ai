package service

import(
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
)



func createSession() *session.Session{
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String(endpoints.UsEast2RegionID),
		Credentials: credentials.NewStaticCredentials(*getAccessKey(), *getSecretKey(), ""),
	})

	return sess
}

func input(filename *string) *rekognition.DetectTextInput{

	fmt.Println("Received File to detect: " + *filename)

	bucket := "imagerekognitionai"
	testimage := filename


	detectTextInput := rekognition.DetectTextInput{}
	image := rekognition.Image{}

	s3Path := *testimage + "/" + *testimage

	s3Object := rekognition.S3Object{}

	s3Object.Bucket= &bucket
	s3Object.Name= &s3Path

	image.S3Object = &s3Object

	detectTextInput.SetImage(&image)
	fmt.Print(detectTextInput)

	return &detectTextInput
}

func Rekognize(filename *string) *rekognition.DetectTextOutput{
	reko := rekognition.New(createSession())
	output, err := reko.DetectText(input(filename))
	fmt.Print(output)

	if (err != nil){
		fmt.Println(err)
	}


	return output
}

