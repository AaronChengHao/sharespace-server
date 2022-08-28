package tool

import (
	"crypto/md5"
	"encoding/hex"
	"mime/multipart"
)

func Md5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

func ExtByFileName(path string) string {
	for i := len(path) - 1; i >= 0 && path[i] != '/'; i-- {
		if path[i] == '.' {
			return path[i:]
		}
	}
	return ""
}

func GenerateUploadPath(file *multipart.FileHeader) string {
	//nanoStr := strconv.Itoa(int(time.Now().UnixNano()))
	return "uploads/" + file.Filename
}
