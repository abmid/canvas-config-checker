package canvas

import (
	"os"
	"path"

	"github.com/abmid/canvas-env-checker/internal/message"
	"github.com/spf13/viper"
)

type CheckerCanvas struct {
	CanvasPath       string
	CanvasPathConfig string
	Database         CheckerCanvasDB     // see in database.go
	Domain           CheckerCanvasDomain // see in domain.go
}

type CheckerNotEqual struct {
	Group string
	Name  string
}

// New initial method CheckerCanvas
func New(viper *viper.Viper) *CheckerCanvas {
	canvasPathConfig := path.Join(viper.GetString("canvas.path"), "config")

	database := CheckerCanvasDB{
		DBName:   viper.GetString("database.dbname"),
		Host:     viper.GetString("database.host"),
		Username: viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		Port:     viper.GetString("database.port"),
	}

	domain := CheckerCanvasDomain{
		Url:              viper.GetString("domain.url"),
		SSL:              viper.GetBool("domain.ssl"),
		ServiceUmm:       viper.GetString("domain.integration"),
		ServiceUmmSecret: viper.GetString("domain.integration_secret"),
	}

	return &CheckerCanvas{
		CanvasPath:       viper.GetString("canvas.path"),
		CanvasPathConfig: canvasPathConfig,
		Database:         database,
		Domain:           domain,
	}
}

// checkDir this function for check directory is exist from configurations settings.yml
func (c *CheckerCanvas) checkDir() (bool, error) {
	if _, err := os.Stat(c.CanvasPath); os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}

// RunCanvas is function for Check Configuration Canvas
func (c *CheckerCanvas) RunCanvas() (notEqual []CheckerNotEqual, err error) {
	m := message.New("Canvas")
	m.Name = "Path"
	m.Start()
	isExist, err := c.checkDir()
	if err != nil {
		m.StopFailure(err.Error())
		notEqual = append(notEqual, CheckerNotEqual{Group: "canvas", Name: "path:directory"})
	}
	if isExist {
		m.StopSuccess()
	}
	return notEqual, nil
}

func (c *CheckerCanvas) CheckConfigEqual(execpted, actual string) bool {

	if execpted != actual {
		return false
	}

	return true
}
