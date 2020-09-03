package canvas

import (
	"github.com/spf13/viper"
)

func (suite *CanvasTestSuite) TestGetCanvasCacheConfig() {
	viper := viper.New()
	viper.SetConfigType("yaml")
	viper.Set("canvas.path", suite.CanvasPath)

	cCanvas := New(viper)
	canvasCache, _ := cCanvas.GetCanvasCacheConfig()
	suite.Assert().Equal(canvasCache.Production.CacheStore, "redis_store")
}

func (suite *CanvasTestSuite) TestRunCache() {
	viper := viper.New()
	listServer := []string{"redis://redis01"}
	viper.SetConfigType("yaml")
	viper.Set("canvas.path", suite.CanvasPath)
	viper.Set("cache_store.status", true)
	viper.Set("cache_store.cache_store", "redis_store")
	viper.Set("cache_store.redis.servers", listServer)
	viper.Set("cache_store.redis.database", 2)

	cCanvas := New(viper)
	res, err := cCanvas.RunCache()
	suite.NoError(err)
	suite.Nil(res)

}
