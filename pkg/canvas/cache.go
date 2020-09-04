package canvas

import (
	"io/ioutil"

	"github.com/abmid/canvas-config-checker/internal/checker"
	"github.com/abmid/canvas-config-checker/internal/message"
	"github.com/fatih/color"
	yaml "gopkg.in/yaml.v2"
)

type CheckerCanvasCache struct {
	Status     bool // use cache store or not
	CacheStore string
	Servers    []string
	Redis      CanvasRedis
}

type CanvasCache struct {
	CacheStore string `yaml:"cache_store"`
}

type CanvasCacheConfig struct {
	Development CanvasCache
	Production  CanvasCache
}

// GetCanvasCache function to get configuration from canvas project
func (c *CheckerCanvas) GetCanvasCacheConfig() (*CanvasCacheConfig, error) {
	m := message.New("Canvas")
	m.Name = "redis file"
	m.File = "cache.yml"
	m.StartGroup()
	// Read configuration from canvas
	contentDB, err := ioutil.ReadFile(c.CanvasPathConfig + "/cache_store.yml")
	if err != nil {
		m.StopFailureNotExists()
		return nil, err
	}

	cacheCanvasConf := CanvasCacheConfig{}

	err = yaml.Unmarshal(contentDB, &cacheCanvasConf)

	if err != nil {
		m.StopFailureNotExists()
		return nil, err
	}

	m.StopSuccess()
	return &cacheCanvasConf, nil
}

// RunCache function for check configuration is equal about cache_store
func (c *CheckerCanvas) RunCache() (notEqual []checker.CheckerNotEqual, err error) {

	// if config not set cache
	if !c.Cache.Status {
		c := color.New(color.FgGreen).Add(color.Bold)
		c.Println("\u21AA \u2757 You set Cache store not active [INFO]")
		return nil, nil
	}

	cacheCanvasConf, err := c.GetCanvasCacheConfig()
	if err != nil {
		return nil, err
	}

	m := message.New("Canvas")

	m.Name = "cache_store:cache_store"
	m.Start()
	// Check config cache_store from settings.yml and canvas config cache store
	if c.CheckConfigEqual(c.Cache.CacheStore, cacheCanvasConf.Production.CacheStore) {
		m.StopSuccess()
	} else {
		m.StopFailureNotEqual()
		notEqual = append(notEqual, checker.CheckerNotEqual{Group: "Canvas", Name: m.Name})
	}

	if c.Cache.CacheStore == "redis_store" {
		isEqual, err := c.RunRedis()
		if err != nil {
			return nil, err
		}
		notEqual = append(notEqual, isEqual...)
	}

	return notEqual, nil

}
