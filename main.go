package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"stock_Project/common"

	"github.com/gin-gonic/gin"
)

func handleTask1(r *gin.Engine) {
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
		cmd := exec.Command("go", "run", "src/t1/task1.go", fileP)
		output, err := cmd.CombinedOutput()
		if err != nil {
			c.String(http.StatusInternalServerError, "Error running task1: %v :%s", err, fileP)
			return
		}

		c.HTML(http.StatusOK, "results.html", gin.H{
			"title":  "Task1 Results:",
			"output": string(output),
		})
	})

	err := os.MkdirAll(savelocation, os.ModePerm)
	if err != nil {
		fmt.Println("Could not create data directory:", err)
	}
}

func handleTask2(r *gin.Engine) {
	r.GET("/run-task2", func(c *gin.Context) {
		cmd := exec.Command("go", "run", "src/t2/task2.go")
		output, err := cmd.CombinedOutput()
		if err != nil {
			c.String(http.StatusInternalServerError, "Error running task2 : %v", err)
			return
		}
		c.HTML(http.StatusOK, "results.html", gin.H{
			"title":  "Task2 Results:",
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

	handleTask1(r)
	handleTask2(r)

	r.Run(":8080")
}
