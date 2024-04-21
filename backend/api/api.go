package api

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var wg sync.WaitGroup

func InitHandlers(r *gin.Engine) {
	r.GET("/status", Status)
	r.POST("/video", uploadVideo)
}

func Status(c *gin.Context) {
	c.JSON(200, gin.H{"name": "hackaton-mephi-2024_backend", "status": "working"})
}

func uploadVideo(c *gin.Context) {
	const op = "uploadVideo"

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "from": op})
		return
	}

	fileExtension := strings.ToLower(file.Filename[strings.LastIndex(file.Filename, ".")+1:])

	if fileExtension != "mp4" && fileExtension != "mp3" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Use only mp3, mp4 or wav extention"})
		return
	}

	fileName := "records/res_vid_" + strconv.FormatInt(time.Now().Unix(), 10)

	err = c.SaveUploadedFile(file, fileName+"."+fileExtension)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "from": op})
		return
	}

	realFile, err := os.Open(fileName + "." + fileExtension)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "from": op})
		return
	}
	realFileBuffer := make([]byte, file.Size)
	_, err = realFile.Read(realFileBuffer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "from": op})
		return
	}

	/* error if wait response */
	//wg.Add(1)
	startVideoProcess(fileName, c)
	//wg.Wait()

	switch fileExtension {
	case "mp3":
		c.Data(http.StatusOK, "blob", realFileBuffer)
	case "mp4":
		c.Data(http.StatusOK, "video/mp4", realFileBuffer)
	}
}

func startVideoProcess(fileName string, c *gin.Context) {
	const op = "startVideoProcess"
	outPath := "./internal/python-nn/cache"
	pyPath := "./internal/python-nn/predict-by-video.py"
	cmd := exec.Command("python3", pyPath, fileName+".mp4", "True", outPath)
	log.Printf("%s", cmd.String())
	out, err := cmd.Output() /* here get output */
	log.Printf("%s", out)
	if err != nil {
		wg.Done()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "from": op})
	}
	wg.Done()
}

func startAudioProcess(fileName string, c *gin.Context) {
	const op = "startAudioProcess"
	pyPath := "./internal/python-nn/predict-by-voice.py"
	cmd := exec.Command("python3", pyPath, fileName+".mp3")
	log.Printf("%s", cmd.String())
	out, err := cmd.Output() /* here get output */
	log.Printf("%s", out)
	if err != nil {
		wg.Done()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "from": op})
	}
	wg.Done()
}
