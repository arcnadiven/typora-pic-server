package utils

import (
	"archive/tar"
	"compress/gzip"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func TarGzFile(dir string) (string, error) {
	dirName := filepath.Base(dir)
	tarGzName := filepath.Join(dir, dirName+".tar.gz")

	tarGzFile, err := os.Create(tarGzName)
	if err != nil {
		return "", errors.WithStack(err)
	}
	defer tarGzFile.Close()

	gw := gzip.NewWriter(tarGzFile)
	defer gw.Close()
	tarGzWriter := tar.NewWriter(gw)
	defer tarGzWriter.Close()

	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		return "", errors.WithStack(err)
	}

	for _, info := range fileInfos {
		if info.IsDir() {
			continue
		}
		if info.Name() == dirName+".tar.gz" {
			continue
		}
		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return "", errors.WithStack(err)
		}
		err = tarGzWriter.WriteHeader(header)
		if err != nil {
			return "", errors.WithStack(err)
		}
		f, err := os.Open(filepath.Join(dir, info.Name()))
		if err != nil {
			return "", errors.WithStack(err)
		}
		_, err = io.Copy(tarGzWriter, f)
		if err != nil {
			f.Close()
			return "", errors.WithStack(err)
		}
		f.Close()
	}
	return tarGzName, nil
}
