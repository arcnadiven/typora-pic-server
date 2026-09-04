package utils

import "testing"

func TestTarGzFile(t *testing.T) {
	tarGzFile, err := TarGzFile("/Users/macintosh/Desktop/doc")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tarGzFile)
}

func TestUnTarGzFile(t *testing.T) {
	tarGzFile, err := UnTarGzFile("/Users/macintosh/Desktop/excel/doc.tar.gz")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tarGzFile)
}
