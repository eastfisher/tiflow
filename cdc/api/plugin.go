package api

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pingcap/tiflow/cdc/model"
	"github.com/pingcap/tiflow/pkg/config"
)

const DefaultPluginPath = "/home/ubuntu/cdc_wasm_plugins"

func (h *openAPI) UploadWasmPlugin(c *gin.Context) {
	var cfg model.WasmPluginConfig
	if err := c.BindJSON(&cfg); err != nil {
		_ = c.Error(err)
		return
	}

	ret, err := base64.URLEncoding.DecodeString(cfg.Binary)
	if err != nil {
		_ = c.Error(err)
		return
	}

	if err := saveWasmPlugin(cfg.Name, ret); err != nil {
		_ = c.Error(err)
		return
	}
	c.Status(http.StatusOK)
}

func (h *openAPI) ListWasmPlugins(c *gin.Context) {
	withBinary := c.Param("with-binary")
	binary, _ := strconv.ParseBool(withBinary)
	cfgs, err := listWasmPlugins(binary)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.IndentedJSON(http.StatusOK, cfgs)
}

func saveWasmPlugin(name string, binary []byte) error {
	fileName := getWasmPluginPath(name)
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	_, err = f.Write(binary)
	return err
}

func listWasmPlugins(withBinary bool) ([]*model.WasmPluginConfig, error) {
	cfg := config.GetGlobalServerConfig()
	infos, err := ioutil.ReadDir(cfg.WasmPluginDir)
	if err != nil {
		return nil, err
	}
	var rets []*model.WasmPluginConfig
	for _, info := range infos {
		ret := &model.WasmPluginConfig{
			Name: info.Name(),
		}
		if withBinary {
			binary, err := ioutil.ReadFile(getWasmPluginPath(info.Name()))
			if err != nil {
				return nil, err
			}
			ret.Binary = string(binary)
		}
		rets = append(rets, ret)
	}
	return rets, nil
}

func getWasmPluginPath(name string) string {
	cfg := config.GetGlobalServerConfig()
	return path.Join(cfg.WasmPluginDir, name)
}
