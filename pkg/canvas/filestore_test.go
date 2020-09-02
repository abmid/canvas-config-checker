package canvas

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

func TestSuiteFileStore(t *testing.T) {
	suite.Run(t, new(CanvasTestSuite))
}

func (suite *CanvasTestSuite) TestGetCanvasFSConfig() {
	viper := viper.New()
	viper.SetConfigType("yaml")
	viper.Set("canvas.path", suite.CanvasPath)

	cCanvas := New(viper)
	canvasFS, _ := cCanvas.GetCanvasFSConfig()
	suite.Assert().Equal(canvasFS.Production.Storage, "local")
}

func (suite *CanvasTestSuite) TestRunFS() {
	viper := viper.New()
	viper.SetConfigType("yaml")
	viper.Set("canvas.path", suite.CanvasPath)
	viper.Set("filestore.storage", "local")
	viper.Set("filestore.path_prefix", "tmp/files")

	cCanvas := New(viper)
	notEquals, err := cCanvas.RunFS()
	suite.NoError(err)
	suite.Nil(notEquals)
}
