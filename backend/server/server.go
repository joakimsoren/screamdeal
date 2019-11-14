package server

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joakimsoren/screamdeal/backend/awsutils"
	"github.com/unidoc/unipdf/v3/creator"
	pdf "github.com/unidoc/unipdf/v3/model"

	// "reflect"
	"bytes"
	"io"
	"io/ioutil"
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
		file, _ := fileHeader.Open()                //get the file
		fileTemp, _ := ioutil.TempFile(".", "temp") //create a temporary file
		buffer := bytes.NewBuffer(nil)              //create empty buffer
		io.Copy(buffer, file)                       //write file to buffer
		fileTemp.Write(buffer.Bytes())              //write buffer to temp file.
		// fmt.Println(reflect.TypeOf(buffer.Bytes()))
		// fmt.Println(fileTemp.)
		// fmt.Println(reflect.TypeOf(fileTemp))
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(200, gin.H{
			"url": "url to pdf in S3",
		})
	})
}

// ThemeToPdf struct for body
type ThemeToPdf struct {
	File  string `json:"file"`
	Theme string `json:"theme"`
}

// Document id - string
// image link - string
func setupAddThemeToPds(router *gin.Engine) {
	router.POST("/add-theme-to-pdf", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")

		var reqBody ThemeToPdf
		c.BindJSON(&reqBody)

		fmt.Println(reqBody)

		// bucket := "scary-bucket"

		// Download document
		// Temp
		// awsutils.DownloadFile("awsesome-o.pdf", bucket)

		// Download image

		// Send into pdf thingy

		// Upload modified pdf

		// Get signed url of new pdf

		temp := awsutils.GetSignedURL("awesome-o.pdf", "scary-bucket")

		c.JSON(200, gin.H{
			"pdf": temp,
		})
	})
}

func addWatermarkImage(inputPath string, outputPath string, watermarkPath string) error {
	c := creator.New()

	watermarkImg, err := c.NewImageFromFile(watermarkPath)
	if err != nil {
		return err
	}

	// Read the input pdf file.
	f, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	pdfReader, err := pdf.NewPdfReader(f)
	if err != nil {
		return err
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return err
	}

	for i := 0; i < numPages; i++ {
		pageNum := i + 1

		// Read the page.
		page, err := pdfReader.GetPage(pageNum)
		if err != nil {
			return err
		}

		// Add to creator.
		c.AddPage(page)

		watermarkImg.ScaleToWidth(c.Context().PageWidth)
		watermarkImg.SetPos(0, (c.Context().PageHeight-watermarkImg.Height())/2)
		watermarkImg.SetOpacity(0.5)
		_ = c.Draw(watermarkImg)
	}

	// Add reader outline tree to the creator.
	c.SetOutlineTree(pdfReader.GetOutlineTree())

	// Add reader AcroForm to the creator.
	c.SetForms(pdfReader.AcroForm)

	err = c.WriteToFile(outputPath)
	return err
}
