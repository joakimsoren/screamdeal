package server

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joakimsoren/screamdeal/backend/awsutils"
	// "reflect"
	// "io/ioutil"
)

func StartServer() {
	router := gin.Default()
	setupThemes(router)
	setupPutPdf(router)
	setupAddThemeToPds(router)
	fmt.Println("Listen in port ", os.Getenv("PORT"))
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func setupThemes(router *gin.Engine) {
	router.GET("/themes", func(c *gin.Context) {
		urls := awsutils.GetThemeURLs()
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(200, gin.H{
			"urls": urls,
		})
	})
}

func setupPutPdf(router *gin.Engine) {
	router.POST("/put-pdf", func(c *gin.Context) {
		// fileHeader, _ := c.FormFile("filepdf")
		// file, _ := fileHeader.Open()
		// fmt.Println(reflect.TypeOf(fileHeader))
		// fmt.Println(reflect.TypeOf(&file))
		// fmt.Println(&file)
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(200, gin.H{
			"url": "url to pdf in S3",
		})
	})
}

func setupAddThemeToPds(router *gin.Engine) {
	router.POST("/add-theme-to-pdf", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")

		temp := awsutils.GetSignedURL("awesome-o.pdf", "scary-bucket")

		c.JSON(200, gin.H{
			"pdf": temp,
		})
	})
}
