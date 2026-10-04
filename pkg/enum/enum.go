package enum

import (
	"github.com/arcnadiven/GoUtils/logger"
	"os"
	"path/filepath"
)

var (
	WorkDir = ""
)

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	WorkDir = filepath.Join(homeDir, ".pictures")
	logger.Infof("WorkDir: %s", WorkDir)
}
