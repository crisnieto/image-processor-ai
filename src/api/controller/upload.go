package controller

import (
	"fmt"
	"github.com/crisnieto/image-processor-ai/src/api/service"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func ReceiveImage(c *gin.Context){
	file, header , err := c.Request.FormFile("upload")
	filename := strconv.FormatInt(time.Now().Unix(),10)
	fmt.Println(header.Filename)
	path := "./tmp/"+filename
	out, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}

	err2 := service.Upload(&filename, &path)

	if(err2 != nil){
		fmt.Println(err2)
		return
	}

	textConcat :=  ""

	for _, element := range service.Rekognize(&filename).TextDetections {
		if (*element.Type == "WORD"){
			textConcat = textConcat + " " + *element.DetectedText
		}
	}

	textConcat = textConcat

	fmt.Println("VOY A DEVOLVER")
	fmt.Println(textConcat)

	c.JSON(200, textConcat)
}
