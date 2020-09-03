package canvas

import (
	"io/ioutil"
	"strconv"

	"github.com/abmid/canvas-config-checker/internal/checker"
	"github.com/abmid/canvas-config-checker/internal/message"
	yaml "gopkg.in/yaml.v2"
)

type CanvasRedis struct {
	Servers  []string
	Database int
}

type CanvasRedisConfig struct {
	Development CanvasRedis
	Production  CanvasRedis
}

func (c *CheckerCanvas) GetCanvasRedisConfig() (*CanvasRedisConfig, error) {
	m := message.New("Canvas")
	m.Name = "redis file"
	m.File = "redis.yml"
	m.StartGroup()
	contentDB, err := ioutil.ReadFile(c.CanvasPathConfig + "/redis.yml")
	if err != nil {
		m.StopFailureNotExists()
		return nil, err
	}

	redisCanvasConf := CanvasRedisConfig{}

	err = yaml.Unmarshal(contentDB, &redisCanvasConf)

	if err != nil {
		m.StopFailureNotExists()
		return nil, err
	}

	m.StopSuccess()
	return &redisCanvasConf, nil
}

func (c *CheckerCanvas) RunRedis() (notEqual []checker.CheckerNotEqual, err error) {

	canvasRedisConf, err := c.GetCanvasRedisConfig()
	if err != nil {
		return nil, err
	}

	m := message.New("Canvas")

	for index, server := range c.Cache.Redis.Servers {
		number := strconv.Itoa(index + 1)
		m.Name = "redis:servers #" + number
		m.Start()
		if c.CheckConfigEqual(server, canvasRedisConf.Production.Servers[index]) {
			m.StopSuccess()
		} else {
			m.StopFailureNotEqual()
			notEqual = append(notEqual, checker.CheckerNotEqual{Group: "Canvas", Name: m.Name})
		}
	}

	m.Name = "redis:database"
	m.Start()
	if c.CheckConfigEqual(strconv.Itoa(c.Cache.Redis.Database), strconv.Itoa(canvasRedisConf.Production.Database)) {
		m.StopSuccess()
	} else {
		m.StopFailureNotEqual()
		notEqual = append(notEqual, checker.CheckerNotEqual{Group: "Canvas", Name: m.Name})
	}

	return notEqual, nil
}
