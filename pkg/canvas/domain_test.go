package canvas

import (
	"github.com/spf13/viper"
)

func (suite *CanvasTestSuite) TestGetCanvasConfigurationDomain() {
	viper := viper.New()
	viper.SetConfigType("yaml")
	viper.Set("canvas.path", suite.CanvasPath)

	cCanvas := New(viper)
	canvasDomain, _ := cCanvas.GetCanvasDBConfig()
	suite.Assert().Equal(canvasDomain.Production.Host, "localhost")
}

func (suite *CanvasTestSuite) TestRunDomain() {
	viper := viper.New()
	viper.SetConfigType("yaml")
	viper.Set("canvas.path", suite.CanvasPath)
	viper.Set("domain.url", "canvas.example.com")
	viper.Set("domain.ssl", true)
	viper.Set("domain.service_umm", "integration.local")
	viper.Set("domain.service_umm_secret", "integration.secret")

	cCanvas := New(viper)
	notEquals, err := cCanvas.RunDomain()
	suite.NoError(err)
	suite.Nil(notEquals)
}
