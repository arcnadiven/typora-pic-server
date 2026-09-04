package syncer

import (
	"encoding/json"
	utilshttp "github.com/arcnadiven/GoUtils/http"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

func loadCredentials() (*ZSpaceVuex, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	data, err := os.ReadFile(filepath.Join(homeDir, "Library", "Application Support", "zspace", "vuex.json"))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	vuex := &ZSpaceVuex{}
	if err := json.Unmarshal(data, vuex); err != nil {
		return nil, errors.WithStack(err)
	}
	return vuex, nil
}

func fileList(vuex *ZSpaceVuex) (map[string]string, error) {

	path := "/v2/file/list"
	rawUrl := "http://127.0.0.1:" + strconv.Itoa(vuex.State.App.LocalPort) + path

	headers := map[string]string{
		"Accept":          "*/*",
		"Accept-Language": "zh-CN",
		"Content-Type":    "application/x-www-form-urlencoded",
		"Connection":      "keep-alive",
		"User-Agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) zspace/2.44.2026080401 Netscape Electron/31.0.1",
	}

	url.Values{}

	utilshttp.DoRequest(http.MethodPost, rawUrl, headers)

}
