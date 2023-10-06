package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func uploadMultipleFile(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	files := form.File["images"]
	filePaths := []string{}
	for _, file := range files {
		fileExt := filepath.Ext(file.Filename)
		originalFileName := strings.TrimSuffix(filepath.Base(file.Filename), filepath.Ext(file.Filename))
		now := time.Now()
		filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt
		filePath := "/images/" + filename

		filePaths = append(filePaths, filePath)
		out, err := os.Create("./uploads/" + filename)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		readerFile, _ := file.Open()
		_, err = io.Copy(out, readerFile)
		if err != nil {
			log.Fatal(err)
		}
	}
	ctx.HTML(http.StatusOK, "index.html", gin.H{"done": true, "msg": "Fotos Guardadas", "filepaths": filePaths})
}

func init() {
	if _, err := os.Stat("uploads"); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll("uploads", os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("public/templates/*")
	router.Static("/assets", "public/assets")
	// router.
	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "success", "message": "How to Upload Single and Multiple Files in Golang"})
	})

	router.POST("/upload", uploadMultipleFile)
	router.StaticFS("/images", http.Dir("uploads"))
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	router.Run(":8000")
}
