package canvas

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"

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
	viper := viper.New()
	viper.SetConfigType("yaml")
	viper.Set("canvas.path", suite.CanvasPath)

	canvas := New(viper)
	test := canvas.RunCanvas()
	fmt.Print(test)
}

func (suite *CanvasTestSuite) TestCheckConfigEqual() {
	viper := viper.New()
	viper.SetConfigType("yaml")
	viper.Set("canvas.path", suite.CanvasPath)

	canvas := New(viper)
	status := canvas.CheckConfigEqual("true", "true")
	suite.Assert().Equal(status, true)
}
