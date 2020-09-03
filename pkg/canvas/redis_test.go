package canvas

import (
	"github.com/spf13/viper"
)

func (suite *CanvasTestSuite) TestGetCanvasRedisConfig() {
	viper := viper.New()
	viper.SetConfigType("yaml")
	viper.Set("canvas.path", suite.CanvasPath)

	cCanvas := New(viper)
	canvasRedis, _ := cCanvas.GetCanvasRedisConfig()
	suite.Assert().Equal(canvasRedis.Production.Database, 2)
}

func (suite *CanvasTestSuite) TestRunRedis() {
	viper := viper.New()
	viper.SetConfigType("yaml")
	viper.Set("canvas.path", suite.CanvasPath)
	// this config must equal with /test/canvas/config/database.yml
	viper.Set("cache_store.redis.servers", []string{"redis://redis01"})
	viper.Set("cache_store.redis.database", 2)

	cCanvas := New(viper)
	notEqual, err := cCanvas.RunRedis()
	suite.Assert().NoError(err)
	suite.Assert().Equal(len(notEqual), 0)
}
