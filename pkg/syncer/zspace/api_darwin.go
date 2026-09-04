package zspace

import (
	"encoding/json"
	"fmt"
	utilshttp "github.com/arcnadiven/GoUtils/http"
	"github.com/arcnadiven/GoUtils/logger"
	"github.com/arcnadiven/typora-pic-server/pkg/utils"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func LoadCredentials() (*ZSpaceVuex, error) {
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

func EncodeCookie(vuex *ZSpaceVuex) string {

	dict := map[string]string{
		"zenithtoken":    vuex.State.User.Token,
		"token":          vuex.State.User.Token,
		"cloudPubKey":    url.PathEscape(vuex.State.Nas.CloudPubKey),
		"cloudPubKeyId":  vuex.State.Nas.CloudPubKeyID,
		"device_id":      vuex.State.App.DeviceID,
		"username":       vuex.State.User.Username,
		"qcname":         vuex.State.User.Qcname,
		"nas_id":         vuex.State.Nas.NasID,
		"deviceColor":    vuex.State.Nas.Color,
		"devicePdt":      vuex.State.Nas.DevicePdt,
		"deviceMode":     vuex.State.Nas.DeviceMode,
		"isMaster":       strconv.Itoa(vuex.State.User.IsMaster),
		"_l":             vuex.State.Nas.Locale,
		"st":             strconv.FormatInt(time.Now().UnixMilli(), 10),
		"plat":           "web",
		"app":            "file",
		"nickname":       "",
		"userid":         "1",
		"isLocal":        "0",
		"nas_name":       vuex.State.Nas.NasName,
		"sign":           vuex.State.Nas.Sign,
		"nasPubKey":      url.PathEscape(vuex.State.Nas.NasPubKey),
		"isAllFlash":     "0",
		"netMode":        "wan",
		"device":         "Mac",
		"clientPublicIp": vuex.State.Nas.ClientPublicIP,
		"publicSwitch":   "true",
		"version":        vuex.State.App.Version,
	}

	cookies := []string{}
	for k, v := range dict {
		cookies = append(cookies, fmt.Sprintf("%s=%s", k, v))
	}

	return strings.Join(cookies, "; ")

}

func FileList(vuex *ZSpaceVuex, path string) (*ZSpaceFileList, error) {

	rawUrl := url.URL{
		Scheme: "http",
		Host:   "127.0.0.1" + ":" + strconv.Itoa(vuex.State.App.LocalPort),
		Path:   "/v2/file/list",
	}

	headers := map[string]string{
		"Accept":          "*/*",
		"Accept-Language": "zh-CN",
		"Content-Type":    "application/x-www-form-urlencoded",
		"Connection":      "keep-alive",
		"User-Agent":      vuex.State.App.Ua,
		"Cookie":          EncodeCookie(vuex),
	}

	vals := url.Values{}
	vals.Set("start", "0")
	vals.Set("num", "100")
	vals.Set("sortby", "mtime")
	vals.Set("order", "desc")
	vals.Set("path", path)
	vals.Set("with_fields", strings.Join([]string{"encrypted", "encrypt_icon", "duration", "nshare", "ori", "ext", "height", "weight", "type", "is_sys", "dw", "labels"}, ","))
	vals.Set("show_hidden", "1")
	vals.Set("dup", "0")

	resp, code, _, err := utilshttp.DoRequest(http.MethodPost, rawUrl.String(), headers, strings.NewReader(vals.Encode()))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if code != http.StatusOK {
		return nil, errors.Errorf("code: %d, resp: %s", code, string(resp))
	}
	fl := &ZSpaceFileList{}
	if err := json.Unmarshal(resp, fl); err != nil {
		return nil, errors.WithStack(err)
	}
	return fl, nil
}

func FileCreate(vuex *ZSpaceVuex, localPath, targetPath string) (*ZSpaceFileCreate, error) {
	dir := filepath.Dir(targetPath)
	fl, err := FileList(vuex, dir)
	if err != nil {
		return nil, err
	}

	isExist := false
	for _, zf := range fl.Data.List {
		if zf.Name == filepath.Base(targetPath) {
			isExist = true
		}
		break
	}
	defer logger.Infof("doFileCreate cost: %+v", time.Since(time.Now()))
	if !isExist {
		return doFileCreate(vuex, localPath, targetPath)
	}

	if _, err = doFileRemove(vuex, targetPath); err != nil {
		return nil, err
	}
	return doFileCreate(vuex, localPath, targetPath)
}

func doFileCreate(vuex *ZSpaceVuex, localPath, targetPath string) (*ZSpaceFileCreate, error) {
	rawUrl := url.URL{
		Scheme: "http",
		Host:   "127.0.0.1" + ":" + strconv.Itoa(vuex.State.App.LocalPort),
		Path:   "/v2/file/create",
	}
	file, err := os.Open(localPath)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	headers := map[string]string{
		"Accept":          "*/*",
		"Accept-Language": "zh-CN",
		"Content-Type":    "application/octet-stream",
		"Connection":      "keep-alive",
		"User-Agent":      vuex.State.App.Ua,
		"Cookie":          EncodeCookie(vuex),
		"path":            targetPath,
	}

	resp, code, _, err := utilshttp.DoRequest(http.MethodPost, rawUrl.String(), headers, file)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if code != http.StatusOK {
		return nil, errors.Errorf("code: %d, resp: %s", code, string(resp))
	}
	//fmt.Println(string(resp))
	fc := &ZSpaceFileCreate{}
	if err := json.Unmarshal(resp, fc); err != nil {
		return nil, errors.WithStack(err)
	}
	return fc, nil
}

func doFileRemove(vuex *ZSpaceVuex, path string) (*ZSpaceFileRemove, error) {
	rawUrl := url.URL{
		Scheme: "http",
		Host:   "127.0.0.1" + ":" + strconv.Itoa(vuex.State.App.LocalPort),
		Path:   "/v2/file/remove",
	}

	headers := map[string]string{
		"Accept":          "*/*",
		"Accept-Language": "zh-CN",
		"Content-Type":    "application/x-www-form-urlencoded",
		"Connection":      "keep-alive",
		"User-Agent":      vuex.State.App.Ua,
		"Cookie":          EncodeCookie(vuex),
	}

	vals := url.Values{}
	vals.Add("paths[]", path)
	vals.Set("skip_bin", "0")
	vals.Set("show_hidden", "1")
	vals.Set("plat", "web")
	vals.Set("version", vuex.State.App.Version)
	vals.Set("device_id", vuex.State.App.DeviceID)
	vals.Set("device", "Mac")
	vals.Set("_l", vuex.State.Nas.Locale)
	vals.Set("token", vuex.State.User.Token)
	vals.Set("nasid", vuex.State.Nas.NasID)

	resp, code, _, err := utilshttp.DoRequest(http.MethodPost, rawUrl.String(), headers, strings.NewReader(vals.Encode()))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if code != http.StatusOK {
		return nil, errors.Errorf("code: %d, resp: %s", code, string(resp))
	}
	//fmt.Println(string(resp))
	fr := ZSpaceFileRemove{}
	if err := json.Unmarshal(resp, &fr); err != nil {
		return nil, errors.WithStack(err)
	}
	return &fr, nil
}

func FileDownload(vuex *ZSpaceVuex, remotePath, localPath string) error {
	vals := url.Values{}
	vals.Set("path", remotePath)

	rawUrl := url.URL{
		Scheme:   "http",
		Host:     "127.0.0.1" + ":" + strconv.Itoa(vuex.State.App.LocalPort),
		Path:     "/v2/file/download",
		RawQuery: vals.Encode(),
	}

	headers := map[string]string{
		"Accept":          "*/*",
		"Accept-Language": "zh-CN",
		"Content-Type":    "application/octet-stream",
		"Connection":      "keep-alive",
		"User-Agent":      vuex.State.App.Ua,
		"Cookie":          EncodeCookie(vuex),
	}

	resp, code, _, err := utils.Download(http.MethodGet, rawUrl.String(), headers, nil, localPath)
	if err != nil {
		return err
	}
	if code != http.StatusOK {
		return errors.Errorf("code: %d, resp: %s", code, string(resp))
	}
	return nil
}
