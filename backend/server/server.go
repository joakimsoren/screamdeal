package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"github.com/joakimsoren/screamdeal/backend/awsutils"
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
		//använd Christoffers funktion för att hämta länkar från AWS S3
		urls := awsutils.GetThemeURLs()

		c.JSON(200, gin.H{
			"urls": urls,
		})
	})
}

func setupPutPdf(router *gin.Engine) {
	router.POST("/put-pdf", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"url": "url to pdf in S3",
		})
	})
}

func setupAddThemeToPds(router *gin.Engine) {
	router.POST("/add-theme-to-pdf", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"pdf": "pdf link to themed file",
		})
	})
}
