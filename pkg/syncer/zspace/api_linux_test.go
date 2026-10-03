package zspace

import (
	"encoding/json"
	"testing"
	"time"
)

func Test_fileList(t *testing.T) {
	vuex, err := LoadCredentials()
	if err != nil {
		t.Fatal(err)
	}
	list, err := FileList(vuex, "/sata11/my/data/.pictures")
	if err != nil {
		t.Fatal(err)
	}
	data, _ := json.Marshal(list)
	t.Log(string(data))
	for _, item := range list.Data.List {
		t.Log(item.Name)
	}
}

func Test_fileCreate(t *testing.T) {
	vuex, _ := LoadCredentials()
	localPath := "/mnt/c/Users/14028/Desktop/Wallpaper/134083761886552430.jpg"
	_, err := FileCreate(vuex, localPath, "/sata11/my/data/note/134083761886552430.jpg")
	if err != nil {
		t.Errorf("%+v", err)
	}
}

func Test_fileRemove(t *testing.T) {
	vuex, _ := LoadCredentials()
	_, err := doFileRemove(vuex, "/sata11/my/data/note/134083761886552430.jpg")
	if err != nil {
		t.Errorf("%+v", err)
	}
}

func Test_fileDownload(t *testing.T) {
	vuex, _ := LoadCredentials()
	localPath := "/mnt/c/Users/14028/Desktop/Wallpaper/134083761886552430.jpg.bak"
	err := FileDownload(vuex, "/sata11/my/data/note/134083761886552430.jpg", localPath)
	t.Log(err)
}

func Test_FileNewDir(t *testing.T) {
	vuex, _ := LoadCredentials()
	FileNewDir(vuex, "/sata11/my/data/note/", "Wallpaper")
}

func Test_timeUnix(t *testing.T) {
	t.Log(time.Unix(1787219824, 0).Format(time.DateTime))
}
