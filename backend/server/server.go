package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"github.com/joakimsoren/screamdeal/backend/awsutils"
	// "reflect"
	"io"
	"io/ioutil"
	"bytes"
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
		fileHeader, _ := c.FormFile("filepdf")
		file, _ := fileHeader.Open() //get the file
		fileTemp, _ := ioutil.TempFile(".", "temp") //create a temporary file
		buffer := bytes.NewBuffer(nil) //create empty buffer
		io.Copy(buffer, file) //write file to buffer
		fileTemp.Write(buffer.Bytes())//write buffer to temp file.
		// fmt.Println(reflect.TypeOf(buffer.Bytes()))
		// fmt.Println(fileTemp.)
		// fmt.Println(reflect.TypeOf(fileTemp))
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(200, gin.H{
			"url": "url to pdf in S3",
		})
	})
}

func setupAddThemeToPds(router *gin.Engine) {
	router.POST("/add-theme-to-pdf", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(200, gin.H{
			"pdf": "pdf link to themed file",
		})
	})
}
