package zspace

import (
	"encoding/json"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

// LoadCredentials implement for OS WSL
func LoadCredentials() (*ZSpaceVuex, error) {
	data, err := os.ReadFile(filepath.Join("/", "mnt", "c", "Users", "14028", "AppData", "Roaming", "zspace", "vuex.json"))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	vuex := &ZSpaceVuex{}
	if err := json.Unmarshal(data, vuex); err != nil {
		return nil, errors.WithStack(err)
	}
	return vuex, nil
}
