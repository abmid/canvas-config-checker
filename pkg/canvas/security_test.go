package canvas

import (
	"github.com/spf13/viper"
)

func (suite *CanvasTestSuite) TestGetCanvasSecurityConfig() {
	viper := viper.New()
	viper.SetConfigType("yaml")
	viper.Set("canvas.path", suite.CanvasPath)

	cCanvas := New(viper)
	canvasSec, _ := cCanvas.GetCanvasSecConfig()
	suite.Assert().Equal(canvasSec.Production.EncryptionKey, "12345asd")
}

func (suite *CanvasTestSuite) TestRunSec() {

	suite.Run("equal", func() {
		viper := viper.New()
		viper.SetConfigType("yaml")
		viper.Set("canvas.path", suite.CanvasPath)
		// this config must equal with /test/canvas/config/database.yml
		viper.Set("security.encryption_key", "12345asd")

		cCanvas := New(viper)
		notEqual, err := cCanvas.RunSec()
		suite.Assert().NoError(err)
		suite.Assert().Equal(len(notEqual), 0)
	})
}
