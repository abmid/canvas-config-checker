package canvas

import (
	"io/ioutil"

	"github.com/abmid/canvas-env-checker/internal/message"
	"github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

// Struct for save configurations from settings.yml
type CheckerCanvasDB struct {
	Host     string
	Port     string
	DBName   string
	Username string
	Password string
}

// Struct for save configuration for canvas
type CanvasDB struct {
	Host     string
	Database string
	Username string
	Password string
}

// Struct for save configuration for Canvas
type CanvasDBConfig struct {
	Production  CanvasDB
	Development CanvasDB
}

var (
	DBHOST     = 1
	DBDATABASE = 2
	DBUSERNAME = 3
	DBPASSWORD = 4
)

// GetCanvasConfig is function to get configuration database from Canvas
func (c *CheckerCanvas) GetCanvasDBConfig() (*CanvasDBConfig, error) {

	contentDB, err := ioutil.ReadFile(c.CanvasPathConfig + "/database.yml")
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	dbCanvasConfig := CanvasDBConfig{}

	err = yaml.Unmarshal(contentDB, &dbCanvasConfig)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return &dbCanvasConfig, nil
}

func (c *CheckerCanvas) RunDB() (notEqual []CheckerNotEqual, err error) {

	canvasConfig, err := c.GetCanvasDBConfig()
	if err != nil {
		return nil, err
	}

	m := message.New("Canvas")

	m.Name = "database:name"
	m.Start()
	if c.CheckConfigEqual(c.Database.DBName, canvasConfig.Production.Database) {
		m.StopSuccess()
	} else {
		m.StopFailureNotEqual()
		notEqual = append(notEqual, CheckerNotEqual{Group: "Canvas", Name: "database:name"})
	}

	m.Name = "database:host"
	m.Start()
	if c.CheckConfigEqual(c.Database.Host, canvasConfig.Production.Host) {
		m.StopSuccess()
	} else {
		m.StopFailureNotEqual()
		notEqual = append(notEqual, CheckerNotEqual{Group: "Canvas", Name: "database:host"})
	}

	m.Name = "database:username"
	m.Start()
	if c.CheckConfigEqual(c.Database.Username, canvasConfig.Production.Username) {
		m.StopSuccess()
	} else {
		m.StopFailureNotEqual()
		notEqual = append(notEqual, CheckerNotEqual{Group: "Canvas", Name: "database:username"})
	}

	m.Name = "database:password"
	m.Start()
	if c.CheckConfigEqual(c.Database.Password, canvasConfig.Production.Password) {
		m.StopSuccess()
	} else {
		m.StopFailureNotEqual()
		notEqual = append(notEqual, CheckerNotEqual{Group: "Canvas", Name: "database:password"})
	}

	return notEqual, nil
}

// Architecture app

// canvas := canvas.New(viper)
// canvas.RunDB()
// canvas.RunDomain()
// canvas.RunFileStore()
// canvas.RunSecurity()
