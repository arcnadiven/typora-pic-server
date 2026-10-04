package zspace

import (
	"github.com/arcnadiven/typora-pic-server/pkg/enum"
	"path/filepath"
	"testing"
)

func Test_Sync(t *testing.T) {
	if err := Sync(enum.WorkDir, filepath.Base(enum.WorkDir)); err != nil {
		t.Fatal(err)
	}
}
