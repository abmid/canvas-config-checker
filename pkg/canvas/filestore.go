package canvas

import (
	"io/ioutil"

	"github.com/abmid/canvas-config-checker/internal/checker"
	"github.com/abmid/canvas-config-checker/internal/message"
	yaml "gopkg.in/yaml.v2"
)

type CheckerCanvasFS struct {
	Storage    string
	PathPrefix string
}

type CanvasFS struct {
	Storage    string
	PathPrefix string `yaml:"path_prefix"`
}

type CanvasFSConfig struct {
	Development CanvasFS
	Production  CanvasFS
}

// GetCanvasConfig is function to get configuration database from Canvas
func (c *CheckerCanvas) GetCanvasFSConfig() (*CanvasFSConfig, error) {

	m := message.New("Canvas")
	m.Name = "file store file"
	m.File = "file_store.yml"
	m.StartGroup()
	contentDB, err := ioutil.ReadFile(c.CanvasPathConfig + "/file_store.yml")
	if err != nil {
		m.StopFailureNotExists()
		return nil, err
	}

	fsCanvasConfig := CanvasFSConfig{}

	err = yaml.Unmarshal(contentDB, &fsCanvasConfig)

	if err != nil {
		m.StopFailureNotExists()
		return nil, err
	}

	m.StopSuccess()
	return &fsCanvasConfig, nil
}

func (c *CheckerCanvas) RunFS() (notEqual []checker.CheckerNotEqual, err error) {

	fsCanvasConfig, err := c.GetCanvasFSConfig()
	if err != nil {
		return nil, err
	}

	m := message.New("Canvas")

	m.Name = "file_store:storage"
	m.Start()
	if c.CheckConfigEqual(c.FileStore.Storage, fsCanvasConfig.Production.Storage) {
		m.StopSuccess()
	} else {
		m.StopFailureNotEqual()
		notEqual = append(notEqual, checker.CheckerNotEqual{Group: "Canvas", Name: m.Name})
	}

	m.Name = "file_store:path_prefix"
	m.Start()
	if c.CheckConfigEqual(c.FileStore.PathPrefix, fsCanvasConfig.Production.PathPrefix) {
		m.StopSuccess()
	} else {
		m.StopFailureNotEqual()
		notEqual = append(notEqual, checker.CheckerNotEqual{Group: "Canvas", Name: m.Name})
	}

	return notEqual, nil
}
