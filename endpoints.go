package rest

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Endpoints() {

	router := gin.Default()
	//router.GET("/", func(ctx *gin.Context) {
	//	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
	//		"title": "Hello from Go and Gin running on Azure App Service",
	//		"link":  "/json",
	//	})
	//})
	router.GET("/pet", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{

			"id":   "100",
			"name": "cats",
			"tag":  "c11",
		})
	})

	router.GET("/CheckCertificationExpiry", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "This a Test for Get Method",
		})
	})

	router.POST("/GetDetails/entities", func(ctx *gin.Context) {
		entity := ctx.PostForm("entity")

		ctx.String(http.StatusOK, "This is your enity value = %s.", entity)
	})

	// Azure App Service sets the port as an Environment Variable
	// This can be random, so needs to be loaded at startup
	port := os.Getenv("HTTP_PLATFORM_PORT")

	// default back to 8080 for local dev
	if port == "" {
		port = "8080"
	}

	router.Run("127.0.0.1:" + port)
}
