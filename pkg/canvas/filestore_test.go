package canvas

import (
	"github.com/spf13/viper"
)

func (suite *CanvasTestSuite) TestGetCanvasFSConfig() {
	viper := viper.New()
	viper.SetConfigType("yaml")
	viper.Set("canvas.path", suite.CanvasPath)

	cCanvas := New(viper)
	canvasFS, _ := cCanvas.GetCanvasFSConfig()
	suite.Assert().Equal(canvasFS.Production.Storage, "local")
}

func (suite *CanvasTestSuite) TestRunFS() {

	suite.Run("equal", func() {

		viper := viper.New()
		viper.SetConfigType("yaml")
		viper.Set("canvas.path", suite.CanvasPath)
		viper.Set("filestore.storage", "local")
		viper.Set("filestore.path_prefix", "tmp/files")

		cCanvas := New(viper)
		notEquals, err := cCanvas.RunFS()
		suite.NoError(err)
		suite.Nil(notEquals)
	})
}
