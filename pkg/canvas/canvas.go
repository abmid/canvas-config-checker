package canvas

import (
	"log"
	"os"
	"path"

	"github.com/abmid/canvas-config-checker/internal/checker"
	"github.com/abmid/canvas-config-checker/internal/message"
	"github.com/spf13/viper"
)

type CheckerCanvas struct {
	CanvasPath       string
	CanvasPathConfig string
	Database         CheckerCanvasDB     // see in database.go
	Domain           CheckerCanvasDomain // see in domain.go
	FileStore        CheckerCanvasFS
	Security         CheckerCanvasSec
	Cache            CheckerCanvasCache
}

// New initial method CheckerCanvas
func New(viper *viper.Viper) *CheckerCanvas {
	canvasPathConfig := path.Join(viper.GetString("canvas.path"), "config")
	// Get config database from settings.yml
	database := CheckerCanvasDB{
		DBName:   viper.GetString("database.dbname"),
		Host:     viper.GetString("database.host"),
		Username: viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		Port:     viper.GetString("database.port"),
	}

	// Get config domain from settings.yml
	domain := CheckerCanvasDomain{
		Url:              viper.GetString("domain.url"),
		SSL:              viper.GetBool("domain.ssl"),
		ServiceUmm:       viper.GetString("domain.integration"),
		ServiceUmmSecret: viper.GetString("domain.integration_secret"),
	}

	// Get config file_store from settings.yml
	fs := CheckerCanvasFS{
		Storage:    viper.GetString("filestore.storage"),
		PathPrefix: viper.GetString("filestore.path_prefix"),
	}

	// Get config security from settings.yml
	sec := CheckerCanvasSec{
		EncryptionKey: viper.GetString("security.encryption_key"),
	}

	// get config cache from settings.yml
	cache := CheckerCanvasCache{
		Status:     viper.GetBool("cache_store.status"),
		CacheStore: viper.GetString("cache_store.cache_store"),
		Redis: CanvasRedis{
			Servers:  viper.GetStringSlice("cache_store.redis.server"),
			Database: viper.GetInt("cache_store.redis.database")},
	}

	return &CheckerCanvas{
		CanvasPath:       viper.GetString("canvas.path"),
		CanvasPathConfig: canvasPathConfig,
		Database:         database,
		Domain:           domain,
		FileStore:        fs,
		Security:         sec,
		Cache:            cache,
	}
}

// checkDir this function for check directory is exist from configurations settings.yml
func (c *CheckerCanvas) checkDir() (bool, error) {
	if _, err := os.Stat(c.CanvasPath); os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}

// RunCanvas function for Check path configuration canvas is exists
func (c *CheckerCanvas) RunCanvas() (notEqual []checker.CheckerNotEqual, err error) {
	m := message.New("Canvas")
	m.Name = "Path"
	m.Start()
	isExist, err := c.checkDir()
	if err != nil {
		m.StopFailure(err.Error())
		notEqual = append(notEqual, checker.CheckerNotEqual{Group: "canvas", Name: "path:directory"})
		return notEqual, err
	}
	if isExist {
		m.StopSuccess()
	}
	return notEqual, nil
}

// CheckConfigEqual function for check two value is equal or not
func (c *CheckerCanvas) CheckConfigEqual(execpted, actual string) bool {

	if execpted != actual {
		return false
	}

	return true
}

// Run function for check all configuration about canvas
func Run(viper *viper.Viper) (notEqual []checker.CheckerNotEqual, groupError []checker.GroupError, err error) {

	canvas := New(viper)

	// Check Canvas
	isEqual, err := canvas.RunCanvas()
	if err != nil {
		log.Fatalf(err.Error())
		return notEqual, groupError, err
	}
	notEqual = append(notEqual, isEqual...)

	// Check DB
	isEqual, err = canvas.RunDB()
	if err != nil {
		groupError = append(groupError, checker.GroupError{Group: "canvas", Message: err.Error()})
	}
	notEqual = append(notEqual, isEqual...)

	// Check Domain
	isEqual, err = canvas.RunDomain()
	if err != nil {
		groupError = append(groupError, checker.GroupError{Group: "canvas", Message: err.Error()})
	}
	notEqual = append(notEqual, isEqual...)

	// Check FileStore
	isEqual, err = canvas.RunFS()
	if err != nil {
		groupError = append(groupError, checker.GroupError{Group: "canvas", Message: err.Error()})
	}
	notEqual = append(notEqual, isEqual...)

	// Check Security
	isEqual, err = canvas.RunSec()
	if err != nil {
		groupError = append(groupError, checker.GroupError{Group: "canvas", Message: err.Error()})
	}
	notEqual = append(notEqual, isEqual...)

	// Check Cache
	isEqual, err = canvas.RunCache()
	if err != nil {
		groupError = append(groupError, checker.GroupError{Group: "canvas", Message: err.Error()})
	}
	notEqual = append(notEqual, isEqual...)

	return notEqual, groupError, nil
}
