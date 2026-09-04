package utils

import (
	utils "github.com/arcnadiven/GoUtils"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"os"
)

func Download(method string, rawUrl string, headers map[string]string, body io.Reader, dstFile string) ([]byte, int, http.Header, error) {

	if utils.FileExists(dstFile) {
		return nil, 0, nil, errors.Errorf("file %s already exist", dstFile)
	}

	file, err := os.Create(dstFile)
	if err != nil {
		return nil, 0, nil, errors.WithStack(err)
	}
	defer file.Close()

	req, err := http.NewRequest(method, rawUrl, body)
	if err != nil {
		return nil, 0, nil, errors.WithStack(err)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	//trans := &http.Transport{
	//	Proxy: http.ProxyFromEnvironment,
	//}
	httpClient := &http.Client{
		//Transport: trans,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, 0, nil, errors.WithStack(err)
	}
	defer resp.Body.Close()

	// 非 200 的 Code 交给外层业务去判断
	if resp.StatusCode != http.StatusOK {
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, 0, nil, errors.WithStack(err)
		}
		return data, resp.StatusCode, resp.Header, nil
	}

	if _, err := io.Copy(file, resp.Body); err != nil {
		return nil, 0, nil, errors.WithStack(err)
	}
	return nil, resp.StatusCode, resp.Header, nil
}
