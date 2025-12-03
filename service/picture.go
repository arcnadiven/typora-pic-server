package service

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/arcnadiven/atalanta/xtools"
	"github.com/astaxie/beego/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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
	if !utils.FileExists(workDir) {
		if err := os.MkdirAll(workDir, os.ModeDir|0755); err != nil {
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
	err = os.WriteFile(filepath.Join(workDir, fileName), body, 0644)
	if err != nil {
		logrus.Errorln(err)
		return
	}

	resp := strings.Join([]string{
		"Upload Success:",
		"http://" + ctx.Request.Host + "/v1/images/" + fileName,
	}, "\n") + "\n"
	ctx.String(http.StatusOK, resp)
}

func Images(ctx *gin.Context) {
	name := ctx.Param("name")
	file, err := os.Open(filepath.Join(workDir, name))
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
