package utils

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

//func TarGzFile(path string) (string, error) {
//	dir := filepath.Dir(path)
//	cleanName := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
//	tarGzPath := filepath.Join(dir, cleanName+".tar.gz")
//
//	tarGzFile, err := os.Create(tarGzPath)
//	if err != nil {
//		return "", errors.WithStack(err)
//	}
//	defer tarGzFile.Close()
//
//	gw := gzip.NewWriter(tarGzFile)
//	defer gw.Close()
//	tarGzWriter := tar.NewWriter(gw)
//	defer tarGzWriter.Close()
//
//	fileInfos, err := ioutil.ReadDir(dir)
//	if err != nil {
//		return "", errors.WithStack(err)
//	}
//
//	for _, info := range fileInfos {
//		if info.IsDir() {
//			continue
//		}
//		if info.Name() == cleanName+".tar.gz" {
//			continue
//		}
//		header, err := tar.FileInfoHeader(info, "")
//		if err != nil {
//			return "", errors.WithStack(err)
//		}
//		err = tarGzWriter.WriteHeader(header)
//		if err != nil {
//			return "", errors.WithStack(err)
//		}
//		f, err := os.Open(filepath.Join(dir, info.Name()))
//		if err != nil {
//			return "", errors.WithStack(err)
//		}
//		_, err = io.Copy(tarGzWriter, f)
//		if err != nil {
//			f.Close()
//			return "", errors.WithStack(err)
//		}
//		f.Close()
//	}
//	return tarGzPath, nil
//}

func TarGzFile(path string) (string, error) {
	dir := filepath.Dir(path)
	cleanName := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	tarGzPath := filepath.Join(dir, cleanName+".tar.gz")
	err := ExecCmd("tar -zcvf %s -C %s %s", tarGzPath, filepath.Dir(path), filepath.Base(path))
	return tarGzPath, err
}

func UnTarGzFile(path string) (string, error) {
	cleanName := strings.TrimSuffix(filepath.Base(path), ".tar.gz")
	err := ExecCmd("tar -zxvf %s -C %s", path, cleanName)
	return cleanName, err
}

func ExecCmd(format string, args ...any) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Minute)
	defer cancel()
	realCmd := fmt.Sprintf(format, args...)

	cmd := exec.CommandContext(ctx, "sh", "-ce", realCmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return errors.Errorf("error executing %s: %s", cmd, err)
	}
	return nil
}

func ExecCmdTo(w io.Writer, format string, args ...any) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Minute)
	defer cancel()
	realCmd := fmt.Sprintf(format, args...)

	cmd := exec.CommandContext(ctx, "sh", "-ce", realCmd)
	cmd.Stdout = w
	cmd.Stderr = w
	err := cmd.Run()
	if err != nil {
		return errors.Errorf("error executing %s: %s", cmd, err)
	}
	return nil
}
