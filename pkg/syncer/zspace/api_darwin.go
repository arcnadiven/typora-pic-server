package zspace

import (
	"encoding/json"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
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
