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
