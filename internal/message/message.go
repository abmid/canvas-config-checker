package message

import (
	"fmt"
	"html"
	"strconv"
	"time"

	"github.com/abmid/canvas-config-checker/internal/checker"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

// Message struct for configuration message
type Message struct {
	Group string
	Name  string
	File  string
	S     *spinner.Spinner
}

var (
	TIME_LOADING = 250 * time.Millisecond
)

// New initialize message
func New(group string) *Message {
	return &Message{
		Group: group,
		S:     spinner.New(spinner.CharSets[14], 100*time.Millisecond),
	}
}

// Banner this function for introduction
func Banner() {
	var stegosaurus = `         \                      .       .
          \                    / ` + "`" + `.   .' "
           \           .---.  <    > <    >  .---.
            \          |    \  \ - ~ ~ - /  /    |
          _____           ..-~             ~-..-~
         |     |   \~~~\\.'                    ` + "`" + `./~~~/
        ---------   \__/                         \__/
       .'  O    \     /               /       \  "
      (_____,    ` + "`" + `._.'               |         }  \/~~~/
       ` + "`" + `----.          /       }     |        /    \__/
             ` + "`" + `-.      |       /      |       /      ` + "`" + `. ,~~|
                 ~-.__|      /_ - ~ ^|      /- _      ` + "`" + `..-'
                      |     /        |     /     ~-.     ` + "`" + `-. _  _  _
                      |_____|        |_____|         ~ - . _ _ _ _ _>

	Please wait... `
	fmt.Print(stegosaurus)
}

// Ready display message for ready deploy
func Ready(env string) {
	str := html.UnescapeString("&#" + strconv.Itoa(128640) + ";")
	c := color.New(color.FgGreen).Add(color.Bold)
	s := spinner.New(spinner.CharSets[36], 100*time.Millisecond)
	s.Color("green")
	s.Start()
	time.Sleep(TIME_LOADING)
	s.Stop()
	c.Println("\n[STATUS] Ready to " + env + " " + str + str + str)
}

// SummaryNotEqual display message summary for configuration not equal
func SummaryNotEqual(notEquals ...[]checker.CheckerNotEqual) {

	finalNotEquals := []checker.CheckerNotEqual{}
	for _, notEqual := range notEquals {
		finalNotEquals = append(finalNotEquals, notEqual...)
	}

	if len(finalNotEquals) == 0 {
		return
	}

	c := color.New(color.FgRed).Add(color.Bold)
	c.Println("\n=== SUMMARY NOT EQUAL, CHECK THIS ===")
	for key, notEqual := range finalNotEquals {
		strIndex := strconv.Itoa(key + 1)
		c.Println(strIndex + ". " + notEqual.Name)
	}
}

// SummaryGroupError display message about summary error group
func SummaryGroupError(errorGroups ...[]checker.GroupError) {

	finalErrors := []checker.GroupError{}
	for _, errGroup := range errorGroups {
		finalErrors = append(finalErrors, errGroup...)
	}

	if len(finalErrors) == 0 {
		return
	}

	c := color.New(color.FgRed).Add(color.Bold)
	c.Println("\n=== SUMMARY GROUP ERROR, CHECK THIS ===")
	for key, errGroup := range finalErrors {
		strIndex := strconv.Itoa(key + 1)
		c.Println(strIndex + ". GROUP = " + errGroup.Group)
		c.Println("- Message = " + errGroup.Message)
	}
}

// Start function to start spinner / loading animation
func (m *Message) Start() {
	m.S.Color("green") // Set the spinner color to red
	// str := html.UnescapeString("&#" + strconv.Itoa(128640) + ";")
	m.S.Suffix = " Check " + m.Group + " " + m.Name + "..." // Build our new spinner
	m.S.Start()
}

// StartGroup function to start spinner with text about the group
func (m *Message) StartGroup() {
	c := color.New(color.FgBlue).Add(color.Bold)
	c.Println("========= " + m.File + " =========")
	m.S.Color("blue") // Set the spinner color to red
	// str := html.UnescapeString("&#" + strconv.Itoa(128640) + ";")
	m.S.Suffix = " Check " + m.Group + " " + m.Name + "..." // Build our new spinner
	m.S.Start()
}

// StopSuccess function to stop spinner if status success
func (m *Message) StopSuccess() {
	time.Sleep(TIME_LOADING)
	m.S.Stop()
	c := color.New(color.FgGreen).Add(color.Bold)
	c.Println("\u21AA \u2705 " + m.Group + " " + m.Name + " success")
}

// StopFailure function to stop spinner if status failed
func (m *Message) StopFailure(msg string) {
	time.Sleep(TIME_LOADING)
	m.S.Stop()
	c := color.New(color.FgRed).Add(color.Bold)
	c.Println("\u21AA \u274C " + m.Group + " " + m.Name + " failed")
	c.Println("[ERROR->" + m.Name + "] " + msg)
}

// StopFailureNotEqual function to stop spinner if check config not equal
func (m *Message) StopFailureNotEqual() {
	time.Sleep(TIME_LOADING)
	m.S.Stop()
	c := color.New(color.FgRed).Add(color.Bold)
	c.Println("\u21AA \u274C " + m.Group + " " + m.Name + " failed")
	c.Println("[ERROR->" + m.Name + "] Not Equal")
}

// StopFailureNotExists function to stop spinner if file config canvas not exists
func (m *Message) StopFailureNotExists() {
	time.Sleep(TIME_LOADING)
	m.S.Stop()
	c := color.New(color.FgRed).Add(color.Bold)
	c.Println("\u21AA \u274C " + m.Group + " " + m.Name + " failed")
	c.Println("[ERROR->" + m.Name + "] file " + m.File + " not exists in your canvas")
}
