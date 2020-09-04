package canvas

import (
	"io/ioutil"

	"github.com/abmid/canvas-config-checker/internal/checker"
	"github.com/abmid/canvas-config-checker/internal/message"
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
	m := message.New("Canvas")
	m.Name = "database file"
	m.File = "database.yml"
	m.StartGroup()
	contentDB, err := ioutil.ReadFile(c.CanvasPathConfig + "/database.yml")
	if err != nil {
		m.StopFailureNotExists()
		return nil, err
	}

	dbCanvasConfig := CanvasDBConfig{}

	err = yaml.Unmarshal(contentDB, &dbCanvasConfig)

	if err != nil {
		m.StopFailureNotExists()
		return nil, err
	}

	m.StopSuccess()
	return &dbCanvasConfig, nil
}

// RunDB function for check configuration settings.yml about database and canvas
func (c *CheckerCanvas) RunDB() (notEqual []checker.CheckerNotEqual, err error) {

	canvasConfig, err := c.GetCanvasDBConfig()
	if err != nil {
		return nil, err
	}

	m := message.New("Canvas")

	// Check Database Name
	m.Name = "database:name"
	m.Start()
	if c.CheckConfigEqual(c.Database.DBName, canvasConfig.Production.Database) {
		m.StopSuccess()
	} else {
		m.StopFailureNotEqual()
		notEqual = append(notEqual, checker.CheckerNotEqual{Group: "Canvas", Name: "database:name"})
	}

	// Check Database Host
	m.Name = "database:host"
	m.Start()
	if c.CheckConfigEqual(c.Database.Host, canvasConfig.Production.Host) {
		m.StopSuccess()
	} else {
		m.StopFailureNotEqual()
		notEqual = append(notEqual, checker.CheckerNotEqual{Group: "Canvas", Name: "database:host"})
	}

	// Check Database Username
	m.Name = "database:username"
	m.Start()
	if c.CheckConfigEqual(c.Database.Username, canvasConfig.Production.Username) {
		m.StopSuccess()
	} else {
		m.StopFailureNotEqual()
		notEqual = append(notEqual, checker.CheckerNotEqual{Group: "Canvas", Name: "database:username"})
	}

	// Check Database Password
	m.Name = "database:password"
	m.Start()
	if c.CheckConfigEqual(c.Database.Password, canvasConfig.Production.Password) {
		m.StopSuccess()
	} else {
		m.StopFailureNotEqual()
		notEqual = append(notEqual, checker.CheckerNotEqual{Group: "Canvas", Name: "database:password"})
	}

	return notEqual, nil
}
