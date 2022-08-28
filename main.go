package main

import (
	"github.com/AaronChengHao/sharespace-server/tool"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var uploadsPath = ""

func init() {
	pwdPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	_, err = os.Stat(pwdPath + "/uploads")
	if err != nil {
		err = os.Mkdir("uploads", os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	uploadsPath = pwdPath + "/uploads"
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// 上传文件
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusInternalServerError, "上传图片出错"+err.Error())
			return
		}

		// 保存文件
		err = c.SaveUploadedFile(file, tool.GenerateUploadPath(file))
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.String(200, "success")
	})

	r.GET("/list", func(c *gin.Context) {
		entries, err := os.ReadDir(uploadsPath)
		if err != nil {
			c.JSON(200, err)
			return
		}
		//infos := make([]fs.FileInfo, 0, len(entries))
		files := make([]SFile, 0, 10)
		for _, entry := range entries {
			info, err := entry.Info()
			if err != nil {
				c.JSON(200, err)
				return
			}
			file := SFile{
				Name:  info.Name(),
				Size:  info.Size(),
				IsDir: info.IsDir(),
			}
			files = append(files, file)
		}

		c.JSON(200, files)
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

type SFile struct {
	Name  string `json:"name"`
	Size  int64  `json:"size"`
	IsDir bool   `json:"isDir"`
	//Name() string       // base name of the file
	//Size() int64        // length in bytes for regular files; system-dependent for others
	//Mode() FileMode     // file mode bits
	//ModTime() time.Time // modification time
	//IsDir() bool        // abbreviation for Mode().IsDir()
	//Sys() any           // underlying data source (can return nil)
}
