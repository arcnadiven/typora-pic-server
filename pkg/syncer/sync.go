package syncer

import (
	"encoding/json"
	"github.com/arcnadiven/GoUtils/logger"
	"github.com/arcnadiven/typora-pic-server/pkg/enum"
	"github.com/arcnadiven/typora-pic-server/pkg/syncer/zspace"
	"github.com/arcnadiven/typora-pic-server/pkg/utils"
	"github.com/dromara/carbon/v2"
	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Run() {
	loc, err := time.LoadLocation(carbon.Shanghai)
	if err != nil {
		panic(err)
	}

	c := cron.New(cron.WithLocation(loc), cron.WithSeconds())
	c.AddFunc("0 0 18 * * *", sync)
	c.AddFunc("0 0 7 * * *", sync)

	c.Start()
}

func sync() {
	if err := doSync(); err != nil {
		logger.Errorf("sync err: %+v", err)
	}
}

func doSync() error {
	pathList := []string{
		enum.WorkDir,
	}

	for _, path := range pathList {
		if err := zspace.Sync(path, filepath.Base(path)); err != nil {
			return err
		}
	}
	return nil
}

func GetModTime(dir string) (time.Time, error) {
	dirEntry, err := os.ReadDir(dir)
	if err != nil {
		return time.Time{}, errors.WithStack(err)
	}
	latestModTime := time.Time{}
	for _, entry := range dirEntry {
		if entry.IsDir() {
			modTime, err := GetModTime(filepath.Join(dir, entry.Name()))
			if err != nil {
				return time.Time{}, err
			}
			if modTime.After(latestModTime) {
				latestModTime = modTime
			}
			continue
		}

		info, err := entry.Info()
		if err != nil {
			return time.Time{}, errors.WithStack(err)
		}
		if info.ModTime().After(latestModTime) {
			latestModTime = info.ModTime()
		}
	}
	return latestModTime, nil
}

func CompressAndSave(vuex *zspace.ZSpaceVuex, localPath, remotePath string) error {
	tarGzFile, err := utils.TarGzFile(localPath)
	if err != nil {
		return err
	}
	defer os.Remove(tarGzFile)
	createResult, err := zspace.FileCreate(vuex, tarGzFile, remotePath)
	if err != nil {
		return err
	}
	data, _ := json.Marshal(createResult)
	logger.Infof("Create result: %s", string(data))
	return nil
}

func DownloadAndUnCompress(vuex *zspace.ZSpaceVuex, localPath, remotePath string) error {
	if err := zspace.FileDownload(vuex, localPath, remotePath); err != nil {
		return err
	}
	defer os.Remove(localPath)
	localDir := strings.TrimSuffix(localPath, filepath.Ext(localPath))
	if err := os.Rename(localDir, localDir+".bak"); err != nil {
		return err
	}
	_, err := utils.UnTarGzFile(localPath)
	if err != nil {
		return err
	}
	return nil
}
