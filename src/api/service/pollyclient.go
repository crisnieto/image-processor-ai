package service

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/polly"
	"io"
	"os"
)

func createPollySession() *session.Session{
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String(endpoints.UsEast2RegionID),
		Credentials: credentials.NewStaticCredentials(*getAccessKey(), *getSecretKey(), ""),
	})
	return sess
}

func Synthesize(text *string, filename *string) (*string ,error){
	svc := polly.New(createPollySession())

	textWithBreak :="<speak>" + *text + " <break time='3000ms'/></speak>"

	textType := "ssml"

	input := &polly.SynthesizeSpeechInput{OutputFormat: aws.String("mp3"), Text: aws.String(textWithBreak), VoiceId: aws.String("Mia"), LanguageCode: aws.String("es-MX")}
	input.TextType = &textType
	output, err := svc.SynthesizeSpeech(input)
	if err != nil {
		fmt.Println("Got error calling SynthesizeSpeech:")
		fmt.Print(err.Error())
		return nil, err
	}

	mp3File := *filename + ".mp3"

	outFile, err2 := os.Create("./tmp/"+mp3File)

	if err2 != nil {
		fmt.Println("Got error calling SynthesizeSpeech:")
		fmt.Print(err2.Error())
		return nil, err2
	}


	defer outFile.Close()
	_, err3 := io.Copy(outFile, output.AudioStream)
	if err3 != nil {
		fmt.Println("Got error saving MP3:")
		fmt.Print(err.Error())
		return nil, err3
	}
	return &mp3File, nil
}



