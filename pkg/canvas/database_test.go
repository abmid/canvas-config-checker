package canvas

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

func TestCanvasTestSuiteDB(t *testing.T) {
	suite.Run(t, new(CanvasTestSuite))
}

func (suite *CanvasTestSuite) TestGetCanvasDBConfig() {
	viper := viper.New()
	viper.SetConfigType("yaml")
	viper.Set("canvas.path", suite.CanvasPath)

	cCanvas := New(viper)
	canvasDB, _ := cCanvas.GetCanvasDBConfig()
	suite.Assert().Equal(canvasDB.Production.Host, "localhost")
}

func (suite *CanvasTestSuite) TestRunDB() {
	viper := viper.New()
	viper.SetConfigType("yaml")
	viper.Set("canvas.path", suite.CanvasPath)
	// this config must equal with /test/canvas/config/database.yml
	suite.Run("equal", func() {
		viper.Set("database.host", "localhost")
		viper.Set("database.dbname", "canvas_production")
		viper.Set("database.username", "canvas")
		viper.Set("database.password", "your_password")

		cCanvas := New(viper)
		notEqual, err := cCanvas.RunDB()
		suite.Assert().NoError(err)
		suite.Assert().Equal(len(notEqual), 0)
	})
	suite.Run("not-equal", func() {
		viper.Set("database.host", "localhosts")

		cCanvas := New(viper)
		notEqual, err := cCanvas.RunDB()
		suite.Assert().NoError(err)
		suite.Assert().Equal(len(notEqual), 1)
	})
}
