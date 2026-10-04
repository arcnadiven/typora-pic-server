package zspace

import (
	"github.com/arcnadiven/GoUtils/logger"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func Sync(localPath, remotePath string) error {
	remoteDir := "/sata11/my/data/note/" + remotePath
	vuex, err := LoadCredentials()
	if err != nil {
		return err
	}
	zspaceFileList, err := FileList(vuex, remoteDir)
	if err != nil {
		return err
	}
	if !zspaceFileList.Exist() {
		fileInfo, err := os.Stat(localPath)
		if err != nil {
			return errors.WithStack(err)
		}
		if fileInfo.IsDir() {
			remoteDirItem := strings.Split(remoteDir, "/")
			parentDir := strings.Join(remoteDirItem[:len(remoteDirItem)-1], "/")
			dirName := remoteDirItem[len(remoteDirItem)-1]
			logger.Infof("mkdir: parent %s dirname %s", parentDir, dirName)
			if err := FileNewDir(vuex, parentDir, dirName); err != nil {
				return err
			}
			subPaths, err := os.ReadDir(localPath)
			if err != nil {
				return errors.WithStack(err)
			}
			for _, subPath := range subPaths {
				// 忽略掉 MacOS 下的 .DS_Store 文件夹
				if runtime.GOOS == "darwin" {
					if subPath.Name() == ".DS_Store" {
						continue
					}
				}
				logger.Infof("path: %s, subPath: %s", localPath, subPath.Name())
				if err := Sync(filepath.Join(localPath, subPath.Name()), filepath.Join(remotePath, subPath.Name())); err != nil {
					return err
				}
			}
		} else {
			logger.Infof("create file: %s to %s", localPath, remoteDir)
			if _, err := FileCreate(vuex, localPath, remoteDir); err != nil {
				return err
			}
			time.Sleep(5 * time.Second)
		}
	}
	return nil
}
