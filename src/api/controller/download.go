package controller

import (
	"fmt"
	"github.com/crisnieto/image-processor-ai/src/api/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func DownloadAudio(c *gin.Context){
	name := c.Param("name")

	fileName, err := service.Download(&name)
	if err != nil{
		fmt.Println(err)
	}

	http.ServeFile(c.Writer, c.Request, *fileName)

	os.Remove(name)
}

