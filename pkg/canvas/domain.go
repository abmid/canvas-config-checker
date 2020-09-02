package canvas

import (
	"io/ioutil"
	"strconv"

	"github.com/abmid/canvas-env-checker/internal/checker"
	"github.com/abmid/canvas-env-checker/internal/message"
	yaml "gopkg.in/yaml.v2"
)

type CheckerCanvasDomain struct {
	Url              string
	SSL              bool
	ServiceUmm       string
	ServiceUmmSecret string
}

type Domain struct {
	Domain           string
	SSL              bool
	ServiceUmm       string
	ServiceUmmSecret string
}

type CanvasDomain struct {
	Production Domain
}

func (c *CheckerCanvas) GetCanvasConfigurationDomain() (*CanvasDomain, error) {
	m := message.New("Canvas")
	m.Name = "domain file"
	m.File = "domain.yml"
	m.StartGroup()
	contentDB, err := ioutil.ReadFile(c.CanvasPathConfig + "/domain.yml")
	if err != nil {
		m.StopFailureNotExists()
		return nil, err
	}

	canvasDomain := CanvasDomain{}

	err = yaml.Unmarshal(contentDB, &canvasDomain)

	if err != nil {
		m.StopFailureNotExists()
		return nil, err
	}
	m.StopSuccess()

	return &canvasDomain, nil
}

func (c *CheckerCanvas) RunDomain() (notEqual []checker.CheckerNotEqual, err error) {

	canvasDomain, err := c.GetCanvasConfigurationDomain()
	if err != nil {
		return nil, err
	}

	m := message.New("Canvas")
	m.Name = "domain:url"
	m.Start()
	if c.CheckConfigEqual(c.Domain.Url, canvasDomain.Production.Domain) {
		m.StopSuccess()
	} else {
		m.StopFailureNotEqual()
		notEqual = append(notEqual, checker.CheckerNotEqual{Group: "Canvas", Name: m.Name})
	}

	m.Name = "domain:ssl"
	m.Start()
	if c.CheckConfigEqual(strconv.FormatBool(c.Domain.SSL), strconv.FormatBool(canvasDomain.Production.SSL)) {
		m.StopSuccess()
	} else {
		m.StopFailureNotEqual()
		notEqual = append(notEqual, checker.CheckerNotEqual{Group: "Canvas", Name: m.Name})
	}

	m.Name = "domain:service_umm"
	m.Start()
	if c.CheckConfigEqual(c.Domain.ServiceUmm, canvasDomain.Production.ServiceUmm) {
		m.StopSuccess()
	} else {
		m.StopFailureNotEqual()
		notEqual = append(notEqual, checker.CheckerNotEqual{Group: "Canvas", Name: m.Name})
	}

	m.Name = "domain:service_umm_secret"
	m.Start()
	if c.CheckConfigEqual(c.Domain.ServiceUmmSecret, canvasDomain.Production.ServiceUmmSecret) {
		m.StopSuccess()
	} else {
		m.StopFailureNotEqual()
		notEqual = append(notEqual, checker.CheckerNotEqual{Group: "Canvas", Name: m.Name})
	}

	return notEqual, nil
}
