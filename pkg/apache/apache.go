package apache

import (
	"io/ioutil"

	"github.com/abmid/canvas-config-checker/internal/checker"
	"github.com/abmid/canvas-config-checker/internal/message"
	"github.com/spf13/viper"
)

type CheckerApache struct {
	ApacheVHostPath string
	OS              string
	VHostName       string
}

// New init apache checker
func New(viper *viper.Viper) (*CheckerApache, error) {
	var apachePath string

	switch viper.GetString("apache.os") {
	case "test":
		apachePath = viper.GetString("apache.path")
	case "ubuntu":
		apachePath = "/etc/apache2/site-enabled"
	}

	return &CheckerApache{
		ApacheVHostPath: apachePath,
		VHostName:       viper.GetString("apache.vhost_name"),
		OS:              viper.GetString("apache.os"),
	}, nil

}

// RunVHost function for check name configuration virtual host apache
func (c *CheckerApache) RunVHost() (notEquals []checker.CheckerNotEqual, err error) {
	// ini message
	m := message.New("Apache")
	m.Name = "vhost_name"
	m.Group = "Apache"
	m.File = "Apache : Virtual Host"
	// Start message
	m.StartGroup()

	// Read directory ensite apache
	files, err := ioutil.ReadDir(c.ApacheVHostPath)
	if err != nil {
		m.StopFailure(err.Error())
		return nil, err
	}

	isExist := false

	// Loop file in directory ensite apache and check file is exists or not
	for _, f := range files {
		if f.Name() == c.VHostName {
			isExist = true
		}
	}

	// If name vhost not found in ensite apache stop message not equal
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

//Run function to check all configuration about apache
func Run(viper *viper.Viper) (notEqual []checker.CheckerNotEqual, groupError []checker.GroupError, err error) {

	apache, err := New(viper)
	if err != nil {
		return notEqual, groupError, err
	}

	// Run virtual host
	isEqual, err := apache.RunVHost()
	if err != nil {
		groupError = append(groupError, checker.GroupError{Group: "apache", Message: err.Error()})
	}
	notEqual = append(notEqual, isEqual...)

	return notEqual, groupError, nil
}
