package main

import (
	"example.com/go-walkthrough/controller"
	"example.com/go-walkthrough/service"
	"github.com/gin-gonic/gin"
)


var (
	videoService service.VideoService = service.New()
	videoController controller.VideoController = controller.New(videoService)
)
func main(){
	app := gin.Default()



	app.GET("/", func (ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"message" : "Working!",
		})
	})

	app.GET("/posts", func (ctx *gin.Context)  {
		ctx.JSON(200, videoController.FindAll())
	})

	app.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})
	app.Run(":8080")

}