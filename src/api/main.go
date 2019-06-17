package main

import (
	"github.com/crisnieto/image-processor-ai/src/api/controller"
	"github.com/gin-gonic/gin"
)
func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", controller.GetPing)
		v1.GET("/download/:name", controller.DownloadAudio)
		v1.POST("/upload", controller.ReceiveImage)
	}
	router.Run()
}