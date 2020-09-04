package canvas

import (
	"io/ioutil"

	"github.com/abmid/canvas-config-checker/internal/checker"
	"github.com/abmid/canvas-config-checker/internal/message"
	yaml "gopkg.in/yaml.v2"
)

type CheckerCanvasSec struct {
	EncryptionKey string
}

type CanvasSec struct {
	EncryptionKey string `yaml:"encryption_key"`
}

type CanvasSecConfig struct {
	Development CanvasSec
	Production  CanvasSec
}

// GetCanvasConfig is function to get configuration security from Canvas
func (c *CheckerCanvas) GetCanvasSecConfig() (*CanvasSecConfig, error) {

	m := message.New("Canvas")
	m.Name = "security file"
	m.File = "security.yml"
	m.StartGroup()
	contentDB, err := ioutil.ReadFile(c.CanvasPathConfig + "/security.yml")
	if err != nil {
		m.StopFailureNotExists()
		return nil, err
	}

	secCanvasConfig := CanvasSecConfig{}

	err = yaml.Unmarshal(contentDB, &secCanvasConfig)

	if err != nil {
		m.StopFailureNotExists()
		return nil, err
	}

	m.StopSuccess()
	return &secCanvasConfig, nil
}

// RunSec function for check config between Security(settings.yml) and canvas sec
func (c *CheckerCanvas) RunSec() (notEqual []checker.CheckerNotEqual, err error) {

	secCanvasConfig, err := c.GetCanvasSecConfig()
	if err != nil {
		return nil, err
	}

	m := message.New("Canvas")

	m.Name = "security:encryption_key"
	m.Start()
	if c.CheckConfigEqual(c.Security.EncryptionKey, secCanvasConfig.Production.EncryptionKey) {
		m.StopSuccess()
	} else {
		m.StopFailureNotEqual()
		notEqual = append(notEqual, checker.CheckerNotEqual{Group: "Canvas", Name: m.Name})
	}

	return notEqual, nil
}
