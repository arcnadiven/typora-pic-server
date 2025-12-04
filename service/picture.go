package service

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/arcnadiven/atalanta/xtools"
	"github.com/astaxie/beego/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var (
	workDir = filepath.Join(os.Getenv("HOME"), ".picture")
	//workDir = filepath.Join(".picture")
)

func init() {
	xtools.UseDefaultLogrus()
}

func Upload(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		logrus.Errorln(err)
		return
	}

	// 初始化 workDir
	month := time.Now().Format("200601")
	dir := filepath.Join(workDir, month)
	if !utils.FileExists(dir) {
		if err := os.MkdirAll(dir, os.ModeDir|0755); err != nil {
			logrus.Errorln(err)
			return
		}
	}

	// 计算hash值，据此生成文件名并存放
	hSHA1 := sha1.New()
	_, err = hSHA1.Write(body)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	fileName := hex.EncodeToString(hSHA1.Sum(nil)) + ".png"

	err = os.WriteFile(filepath.Join(dir, fileName), body, 0644)
	if err != nil {
		logrus.Errorln(err)
		return
	}

	resp := fmt.Sprintf("http://%s/v1/images/%s/%s", ctx.Request.Host, month, fileName) + "\n"
	ctx.String(http.StatusOK, resp)
}

func Images(ctx *gin.Context) {
	dir := ctx.Param("dir")
	name := ctx.Param("name")
	file, err := os.Open(filepath.Join(workDir, dir, name))
	if err != nil {
		logrus.Errorln(err)
		return
	}
	defer file.Close()
	_, err = io.Copy(ctx.Writer, file)
	if err != nil {
		logrus.Errorln(err)
		return
	}
}
