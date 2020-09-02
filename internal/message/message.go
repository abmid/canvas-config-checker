package message

import (
	"fmt"
	"html"
	"strconv"
	"time"

	"github.com/abmid/canvas-env-checker/internal/checker"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

type Message struct {
	Group string
	Name  string
	File  string
	S     *spinner.Spinner
}

func New(group string) *Message {
	return &Message{
		Group: group,
		S:     spinner.New(spinner.CharSets[14], 100*time.Millisecond),
	}
}

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

func Ready(env string) {
	str := html.UnescapeString("&#" + strconv.Itoa(128640) + ";")
	c := color.New(color.FgGreen).Add(color.Bold)
	s := spinner.New(spinner.CharSets[36], 100*time.Millisecond)
	s.Color("green")
	s.Start()
	time.Sleep(400 * time.Millisecond)
	s.Stop()
	c.Println("\n[STATUS] Ready to " + env + " " + str + str + str)
}

func SummaryNotEqual(notEquals []checker.CheckerNotEqual) {

	if len(notEquals) == 0 {
		return
	}

	c := color.New(color.FgRed).Add(color.Bold)
	c.Println("=== SUMMARY NOT EQUAL, CHECK THIS ===")
	for key, notEqual := range notEquals {
		strIndex := strconv.Itoa(key + 1)
		c.Println(strIndex + ". " + notEqual.Name)
	}
}

func SummaryGroupError(errorGroups []checker.GroupError) {

	if len(errorGroups) == 0 {
		return
	}

	c := color.New(color.FgRed).Add(color.Bold)
	c.Println("=== SUMMARY GROUP ERROR, CHECK THIS ===")
	for key, errGroup := range errorGroups {
		strIndex := strconv.Itoa(key + 1)
		c.Println(strIndex + ". GROUP = " + errGroup.Group)
		c.Println("- Message = " + errGroup.Message)
	}
}

func (m *Message) Start() {
	m.S.Color("green") // Set the spinner color to red
	// str := html.UnescapeString("&#" + strconv.Itoa(128640) + ";")
	m.S.Suffix = " Check " + m.Group + " " + m.Name + "..." // Build our new spinner
	m.S.Start()
}

func (m *Message) StartGroup() {
	c := color.New(color.FgBlue).Add(color.Bold)
	c.Println("========= " + m.File + " =========")
	m.S.Color("blue") // Set the spinner color to red
	// str := html.UnescapeString("&#" + strconv.Itoa(128640) + ";")
	m.S.Suffix = " Check " + m.Group + " " + m.Name + "..." // Build our new spinner
	m.S.Start()
}

func (m *Message) StopSuccess() {
	time.Sleep(400 * time.Millisecond)
	m.S.Stop()
	c := color.New(color.FgGreen).Add(color.Bold)
	c.Println("\u21AA \u2705 " + m.Group + " " + m.Name + " success")
}

func (m *Message) StopFailure(msg string) {
	time.Sleep(400 * time.Millisecond)
	m.S.Stop()
	c := color.New(color.FgRed).Add(color.Bold)
	c.Println("\u21AA \u274C " + m.Group + " " + m.Name + " failed")
	c.Println("[ERROR->" + m.Name + "] " + msg)
}

func (m *Message) StopFailureNotEqual() {
	time.Sleep(400 * time.Millisecond)
	m.S.Stop()
	c := color.New(color.FgRed).Add(color.Bold)
	c.Println("\u21AA \u274C " + m.Group + " " + m.Name + " failed")
	c.Println("[ERROR->" + m.Name + "] Not Equal")
}

func (m *Message) StopFailureNotExists() {
	time.Sleep(400 * time.Millisecond)
	m.S.Stop()
	c := color.New(color.FgRed).Add(color.Bold)
	c.Println("\u21AA \u274C " + m.Group + " " + m.Name + " failed")
	c.Println("[ERROR->" + m.Name + "] file " + m.File + " not exists in your canvas")
}
