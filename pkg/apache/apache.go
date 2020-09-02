package apache

import (
	"io/ioutil"
	"path"

	"github.com/abmid/canvas-env-checker/internal/checker"
	"github.com/abmid/canvas-env-checker/internal/message"
	"github.com/spf13/viper"
)

type CheckerApache struct {
	ApacheVHostPath string
	OS              string
	VHostName       string
}

func New(viper *viper.Viper) (*CheckerApache, error) {
	var apachePath string

	switch viper.GetString("apache.os") {
	case "test":
		apachePath = path.Join("../../", "test/apache/site-enabled")
	case "ubuntu":
		apachePath = "/etc/apache2/site-enabled"
	}

	return &CheckerApache{
		ApacheVHostPath: apachePath,
		VHostName:       viper.GetString("apache.vhost_name"),
		OS:              viper.GetString("apache.os"),
	}, nil

}

func (c *CheckerApache) RunVHost() (notEquals []checker.CheckerNotEqual, err error) {

	m := message.New("Apache")
	m.Name = "vhost_name"
	m.Start()
	files, err := ioutil.ReadDir(c.ApacheVHostPath)
	if err != nil {
		m.StopFailure(err.Error())
		return nil, err
	}

	isExist := false

	for _, f := range files {
		if f.Name() == c.VHostName {
			isExist = true
		}
	}

	if !isExist {
		notEquals = append(notEquals, checker.CheckerNotEqual{
			Group: "Apache",
			Name:  "vhost_name",
		})
		m.StopFailureNotEqual()
		return notEquals, nil
	}

	m.StopSuccess()

	return notEquals, nil
}

func Run(viper *viper.Viper) (notEqual []checker.CheckerNotEqual, groupError []checker.GroupError, err error) {

	apache, err := New(viper)
	if err != nil {
		return notEqual, groupError, err
	}

	isEqual, err := apache.RunVHost()
	if err != nil {
		groupError = append(groupError, checker.GroupError{Group: "file store", Message: err.Error()})
	}
	notEqual = append(notEqual, isEqual...)

	return notEqual, groupError, nil
}
