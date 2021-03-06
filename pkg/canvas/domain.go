package canvas

import (
	"io/ioutil"
	"strconv"

	"github.com/abmid/canvas-config-checker/internal/checker"
	"github.com/abmid/canvas-config-checker/internal/message"
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

// GetCanvasConfigDomain function for get configuration about domain from canvas
func (c *CheckerCanvas) GetCanvasConfigDomain() (*CanvasDomain, error) {
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

// RunDomain function for check configuration settings.yml about domain
func (c *CheckerCanvas) RunDomain() (notEqual []checker.CheckerNotEqual, err error) {
	// Get Canvas Config about Domain
	canvasDomain, err := c.GetCanvasConfigDomain()
	if err != nil {
		return nil, err
	}

	// Check Domain URL
	m := message.New("Canvas")
	m.Name = "domain:url"
	m.Start()
	if c.CheckConfigEqual(c.Domain.Url, canvasDomain.Production.Domain) {
		m.StopSuccess()
	} else {
		m.StopFailureNotEqual()
		notEqual = append(notEqual, checker.CheckerNotEqual{Group: "Canvas", Name: m.Name})
	}

	// Check Domain SSL
	m.Name = "domain:ssl"
	m.Start()
	if c.CheckConfigEqual(strconv.FormatBool(c.Domain.SSL), strconv.FormatBool(canvasDomain.Production.SSL)) {
		m.StopSuccess()
	} else {
		m.StopFailureNotEqual()
		notEqual = append(notEqual, checker.CheckerNotEqual{Group: "Canvas", Name: m.Name})
	}

	// Check Domain Service UMM
	m.Name = "domain:service_umm"
	m.Start()
	if c.CheckConfigEqual(c.Domain.ServiceUmm, canvasDomain.Production.ServiceUmm) {
		m.StopSuccess()
	} else {
		m.StopFailureNotEqual()
		notEqual = append(notEqual, checker.CheckerNotEqual{Group: "Canvas", Name: m.Name})
	}

	// Check Domain Service UMM Secret Key
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
