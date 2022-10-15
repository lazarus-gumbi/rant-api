package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lazarus-gumbi/sample_api/controllers"
	"github.com/lazarus-gumbi/sample_api/initializers"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/rant", controllers.RantCreate)
	r.GET("/rants", controllers.RantsIndex)
	r.GET("rants/:company", controllers.RantsShow)
	r.DELETE("/rants/:id", controllers.RantDelete)
	r.Run() // listen and serve on 0.0.0.0:8080
}
