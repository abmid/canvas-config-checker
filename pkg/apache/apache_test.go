package apache

import (
	"path"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type ApacheTestSuite struct {
	suite.Suite
	OS              string
	VHostName       string
	ApacheVHostPath string
}

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func TestRunApacheSuite(t *testing.T) {
	suite.Run(t, new(ApacheTestSuite))
}

// SetupTest create virtual Canvas File
func (suite *ApacheTestSuite) SetupTest() {
	// Get Root Caller
	root := RootDir()
	suite.OS = "test"
	suite.VHostName = "canvas-prod.conf"
	suite.ApacheVHostPath = path.Join(root, "../test/apache/site-enabled")
}

func (suite *ApacheTestSuite) TestNew() {
	// Join path to test canvas dir
	suite.Run("os-test", func() {
		viper := viper.New()
		viper.SetConfigType("yml")
		viper.Set("apache.os", suite.OS)
		viper.Set("apache.vhost_name", suite.VHostName)
		viper.Set("apache.path", suite.ApacheVHostPath)

		apache, err := New(viper)
		suite.NoError(err)
		suite.Equal(apache.OS, "test")
	})

	suite.Run("os-ubuntu", func() {
		viper := viper.New()
		viper.SetConfigType("yml")
		viper.Set("apache.os", "ubuntu")
		viper.Set("apache.vhost_name", suite.VHostName)
		viper.Set("apache.path", suite.ApacheVHostPath)

		apache, err := New(viper)
		suite.NoError(err)
		suite.Equal(apache.OS, "ubuntu")
	})
}

func (suite *ApacheTestSuite) TestRun() {
	viper := viper.New()
	viper.SetConfigType("yml")
	viper.Set("apache.os", suite.OS)
	viper.Set("apache.vhost_name", suite.VHostName)
	viper.Set("apache.path", suite.ApacheVHostPath)

	notEqual, groupErr, err := Run(viper)
	suite.NoError(err)
	suite.Nil(notEqual)
	suite.Nil(groupErr)
}

func (suite *ApacheTestSuite) TestRunVHost() {

	viper := viper.New()
	viper.SetConfigType("yml")
	viper.Set("apache.os", suite.OS)
	viper.Set("apache.vhost_name", suite.VHostName)
	viper.Set("apache.path", suite.ApacheVHostPath)

	apache, _ := New(viper)
	res, err := apache.RunVHost()
	suite.NoError(err)
	suite.Nil(res)
}
