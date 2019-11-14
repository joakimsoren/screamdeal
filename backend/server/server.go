package server

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joakimsoren/screamdeal/backend/awsutils"
	// "reflect"
	// "io"
	// "io/ioutil"
	// "bytes"
	"github.com/gin-contrib/cors"
	// "unsafe"
)

func StartServer() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8081"}
	// config.AllowOrigins == []string{"http://google.com", "http://facebook.com"}

	router.Use(cors.New(config))
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
		// file, _ := fileHeader.Open() //get the file
		// // fileTemp, _ := ioutil.TempFile(".", "temp") //create a temporary file
		// buffer := bytes.NewBuffer(nil) //create empty buffer
		// io.Copy(buffer, file) //write file to buffer
		// message := []byte("Hello, Gophers!")
		// error := ioutil.WriteFile("testdata/hello.pdf", message, 0644)
		// fmt.Println(error)
		// // fileTemp.Write(buffer.Bytes())//write buffer to temp file.
		// // stat, _ := fileTemp.Stat()
		// // fmt.Println(stat.Name(), stat.Size())
		// // fmt.Println(reflect.TypeOf(buffer.Bytes()))
		// // fmt.Println(unsafe.Sizeof(buffer.Bytes()))
		// // fmt.Println(fileTemp.Stat().Size())
		// // fmt.Println(reflect.TypeOf(fileTemp))
		// // awsutils.UploadFile(fileTemp) 
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
