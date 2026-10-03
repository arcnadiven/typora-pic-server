package zspace

import (
	"testing"
	"time"
)

func Test_fileList(t *testing.T) {
	vuex, _ := LoadCredentials()
	FileList(vuex, "/sata11/my/data")
}

func Test_fileCreate(t *testing.T) {
	vuex, _ := LoadCredentials()
	localPath := "/Users/macintosh/Pictures/pap.er/2153816083373490176.heic"
	_, err := FileCreate(vuex, localPath, "/sata11/my/data/note/2153815980835340288.heic")
	if err != nil {
		t.Errorf("%+v", err)
	}
}

func Test_fileRemove(t *testing.T) {
	vuex, _ := LoadCredentials()
	_, err := doFileRemove(vuex, "/sata11/my/data/note/2153815980835340288.heic")
	if err != nil {
		t.Errorf("%+v", err)
	}
}

func Test_fileDownload(t *testing.T) {
	vuex, _ := LoadCredentials()
	localPath := "/Users/macintosh/Pictures/pap.er/2153815980835340288.heic.bak"
	err := FileDownload(vuex, "/sata11/my/data/note/2153815980835340288.heic", localPath)
	t.Log(err)
}

func Test_timeUnix(t *testing.T) {
	t.Log(time.Unix(1787219824, 0).Format(time.DateTime))
}
