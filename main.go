package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"os"
	"path/filepath"
)

const uploadPath = "./uploads"

func main() {
	// 创建上传目录
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		os.Mkdir(uploadPath, os.ModePerm)
	}

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	// 首页
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// 文件上传
	router.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
			return
		}

		// 生成唯一文件名
		filename := uuid.New().String() + "_" + file.Filename
		filePath := filepath.Join(uploadPath, filename)

		// 保存文件
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save the file"})
			return
		}

		downloadURL := c.Request.Host + "/download/" + filename
		c.JSON(http.StatusCreated, gin.H{"url": downloadURL})
	})

	// 文件下载
	router.GET("/download/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		filePath := filepath.Join(uploadPath, filename)

		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
			return
		}

		c.FileAttachment(filePath, filename)
	})

	router.Run(":8080")
}
