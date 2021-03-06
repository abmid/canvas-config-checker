package canvas

import (
	"path"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type CanvasTestSuite struct {
	suite.Suite
	CanvasPath string
}

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

// SetupTest create virtual Canvas File
func (suite *CanvasTestSuite) SetupTest() {
	// Get Root Caller
	root := RootDir()
	// Join path to test canvas dir
	canvasTestDir := path.Join(root, "../test/canvas")
	// Set suite.canvasPath to test dir
	suite.CanvasPath = canvasTestDir
}

func (suite *CanvasTestSuite) TestCheckDir() {

	viper := viper.New()
	viper.SetConfigType("yaml")
	viper.Set("canvas.path", suite.CanvasPath)

	canvas := New(viper)
	exist, _ := canvas.checkDir()
	suite.Assert().Equal(true, exist)
}

func (suite *CanvasTestSuite) TestCheckCanvas() {
	suite.Run("equal", func() {
		viper := viper.New()
		viper.SetConfigType("yaml")
		viper.Set("canvas.path", suite.CanvasPath)

		canvas := New(viper)
		notEquals, err := canvas.RunCanvas()
		suite.NoError(err)
		suite.Nil(notEquals)
	})
}

func (suite *CanvasTestSuite) TestCheckConfigEqual() {
	suite.Run("equal", func() {
		viper := viper.New()
		viper.SetConfigType("yaml")
		viper.Set("canvas.path", suite.CanvasPath)

		canvas := New(viper)
		status := canvas.CheckConfigEqual("true", "true")
		suite.Assert().Equal(status, true)
	})
}

func TestCanvasTestSuiteCache(t *testing.T) {
	suite.Run(t, new(CanvasTestSuite))
}
