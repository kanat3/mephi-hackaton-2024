package api

import (
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func InitHandlers(r *gin.Engine) {
	r.GET("/status", Status)
	r.POST("/video", uploadVideo)
}

func Status(c *gin.Context) {
	c.JSON(200, gin.H{"name": "hackaton-mephi-2024_backend", "status": "working"})
}

func uploadVideo(c *gin.Context) {
	const op = "uploadVideo"

	file, err := c.FormFile("files")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "from": op})
		return
	}

	fileExtension := strings.ToLower(file.Filename[strings.LastIndex(file.Filename, ".")+1:])

	if fileExtension != "mp4" && fileExtension != "mp3" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Use only mp3, mp4 or wav extention"})
		return
	}

	fileName := "res_vid_" + strconv.FormatInt(time.Now().Unix(), 10)

	err = c.SaveUploadedFile(file, fileName+"."+fileExtension)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "from": op})
		return
	}

	uploadedFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "from": op})
		return
	}

	realFile, err := os.CreateTemp("", "tempfile-*.txt")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "from": op})
		return
	}
	defer realFile.Close()

	_, err = io.Copy(realFile, uploadedFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "from": op})
		return
	}

	fileinfo, err := realFile.Stat()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "from": op})
		return
	}

	fileBuffer := make([]byte, fileinfo.Size())

	_, err = realFile.Read(fileBuffer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "from": op})
		return
	}

	switch fileExtension {
	case "mp3":
		c.Data(http.StatusOK, "audio/mp3", fileBuffer)
	case "mp4":
		c.Data(http.StatusOK, "video/mp4", fileBuffer)
	}
}
