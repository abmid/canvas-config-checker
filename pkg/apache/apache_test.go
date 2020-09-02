package apache

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type ApacheTestSuite struct {
	suite.Suite
	OS string
}

// SetupTest create virtual Canvas File
func (suite *ApacheTestSuite) SetupTest() {
	suite.OS = "test"
}

func (suite *ApacheTestSuite) TestNew(t *testing.T) {
	// Join path to test canvas dir
	viper := viper.New()
	viper.SetConfigType("yml")
	viper.Set("apache.os", suite.OS)
	viper.Set("apache.vhost_name", "canvas-prod.conf")

	New(viper)
}

func TestRunApacheSuite(t *testing.T) {
	suite.Run(t, new(ApacheTestSuite))
}

func (suite *ApacheTestSuite) TestRunVHost() {

	viper := viper.New()
	viper.SetConfigType("yml")
	viper.Set("apache.os", suite.OS)
	viper.Set("apache.vhost_name", "canvas-prod.conf")

	apache, _ := New(viper)
	res, err := apache.RunVHost()
	suite.NoError(err)
	suite.Nil(res)
}
