package syncer

import (
	"encoding/json"
	"github.com/arcnadiven/GoUtils/logger"
	"github.com/arcnadiven/typora-pic-server/pkg/syncer/zspace"
	"github.com/arcnadiven/typora-pic-server/pkg/utils"
	"github.com/dromara/carbon/v2"
	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
	"os"
	"path/filepath"
	"strconv"
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
		"/Users/macintosh/Desktop/doc",
		//"/Users/macintosh/.picture",
	}

	remoteDir := "/sata11/my/data/note"

	vuex, err := zspace.LoadCredentials()
	if err != nil {
		return err
	}
	listRets, err := zspace.FileList(vuex, remoteDir)
	if err != nil {
		return err
	}
	dict1 := map[string]zspace.ZSpaceFile{}
	for _, ret := range listRets.Data.List {
		dict1[ret.Name] = ret
	}

	for _, path := range pathList {
		remoteFile := filepath.Base(path) + ".tar.gz"
		localModTime, err := GetModTime(path)
		if err != nil {
			return err
		}
		zspaceFile, ok := dict1[remoteFile]

		if !ok {
			if err := CompressAndSave(vuex, path, remoteDir+"/"+remoteFile); err != nil {
				return err
			}
		} else {
			remoteModTime, err := strconv.ParseInt(zspaceFile.ModifyTime, 10, 64)
			if err != nil {
				return err
			}
			if localModTime.After(time.Unix(remoteModTime, 0)) {
				if err := CompressAndSave(vuex, path, remoteDir+"/"+remoteFile); err != nil {
					return err
				}
			} else if localModTime.Before(time.Unix(remoteModTime, 0)) {
				// TODO: DownloadAndUnCompress
				if err := DownloadAndUnCompress(vuex, path+".tar.gz", remoteDir+"/"+remoteFile); err != nil {
					return err
				}
			}
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
