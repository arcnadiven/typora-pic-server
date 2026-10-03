package enum

import (
	"os"
	"path/filepath"
)

var (
	WorkDir = filepath.Join(os.Getenv("HOME"), ".pictures")
)
