package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"stock_Project/common"

	"github.com/gin-gonic/gin"
)

func handleExtractor(r *gin.Engine) {
	savelocation := common.Location + "/data"
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("csvfile")
		if err != nil {
			c.String(http.StatusBadRequest, "Bad request: %v", err)
			return
		}
		filePath := fmt.Sprintf("%s/%s", savelocation, file.Filename)
		fmt.Println(filePath)
		err = c.SaveUploadedFile(file, filePath)
		if err != nil {
			c.String(http.StatusInternalServerError, "Could not save file: %v", err)
			return
		}
		fileP := "data/" + file.Filename
		cmd := exec.Command("go", "run", "src/extractor/extractor.go", fileP)
		output, err := cmd.CombinedOutput()
		if err != nil {
			c.String(http.StatusInternalServerError, "Error running Extractor: %v :%s", err, fileP)
			return
		}

		c.HTML(http.StatusOK, "results.html", gin.H{
			"title":  "Extractor Results:",
			"output": string(output),
		})
	})

	err := os.MkdirAll(savelocation, os.ModePerm)
	if err != nil {
		fmt.Println("Could not create data directory:", err)
	}
}

func handlePredictor(r *gin.Engine) {
	r.GET("/run-predictor", func(c *gin.Context) {
		cmd := exec.Command("go", "run", "src/predictor/predictor.go")
		output, err := cmd.CombinedOutput()
		if err != nil {
			c.String(http.StatusInternalServerError, "Error running Predictor : %v", err)
			return
		}
		c.HTML(http.StatusOK, "results.html", gin.H{
			"title":  "Predictor Results:",
			"output": string(output),
		})
	})
}
func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Process data for stock lists",
		})
	})

	handleExtractor(r)
	handlePredictor(r)

	r.Run(":8080")
}
