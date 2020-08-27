package main

import (
	"time"

	"github.com/abmid/umm-canvas-checklist/internal/message"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

func CanvasPath() {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Color("green") // Set the spinner color to red
	// str := html.UnescapeString("&#" + strconv.Itoa(128640) + ";")
	s.Suffix = " Check Canvas Path..." // Build our new spinner
	s.Start()                          // Start the spinner
	time.Sleep(2 * time.Second)
	s.Stop()
	c := color.New(color.FgGreen).Add(color.Bold)
	c.Println("[+] Canvas Path Success \u2705")

}

func CanvasDB() {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Color("green") // Set the spinner color to red
	// str := html.UnescapeString("&#" + strconv.Itoa(128640) + ";")
	s.Suffix = " Check Canvas DB..." // Build our new spinner
	s.Start()                        // Start the spinner
	time.Sleep(2 * time.Second)
	s.Stop()
	c := color.New(color.FgHiGreen).Add(color.Bold)
	c.Println("[+] Canvas DB Success \u2705")

}

// type Message struct {
// 	Status bool
// 	Err    error
// }

// func message(group, name string, ch <-chan Message, wg *sync.WaitGroup) {
// 	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
// 	s.Color("green") // Set the spinner color to red
// 	// str := html.UnescapeString("&#" + strconv.Itoa(128640) + ";")
// 	s.Suffix = " Check " + group + " " + name + "..." // Build our new spinner
// 	s.Start()                                         // Start the spinner
// 	time.Sleep(2 * time.Second)
// 	for _ = range ch {
// 		c := color.New(color.FgHiGreen).Add(color.Bold)
// 		c.Println("[+] Canvas " + name + " success \u2705")
// 		s.Stop()
// 		// if m.Status {
// 		// 	c := color.New(color.FgHiGreen).Add(color.Bold)
// 		// 	c.Println("[+] Canvas " + name + " success \u2705")
// 		// 	s.Stop()
// 		// } else {
// 		// 	c := color.New(color.FgRed).Add(color.Bold)
// 		// 	c.Println("[x] Canvas " + name + " failed \u2705")
// 		// 	s.Stop()
// 		// }
// 		wg.Done()
// 	}
// }

func main() {

	message.Banner()

	m := message.New("Canvas", "database:host")
	m.Start()
	m.StopSuccess()

	m2 := message.New("Canvas", "database:port")
	m2.Start()
	m2.StopFailure("Configuration not valid")

	m3 := message.New("Canvas", "database:password")
	m3.Start()
	m3.StopSuccess()
}

// // func workerMessage()

// // func check(){
// 	ch := string
// 	go workerMessage(ch)
// //	res := checkEqualHost
// 		ch <- res
// //  res := checkEqualDatabase
// // 	res := checkEqualPassword
// //}

// // checkEqualHost(){
// 	message := new Test{}
// 	message.Start()
// 	if(equal == config)
// 		message.StopSuccess(status)
// 	else
// 		message.StopFailed(messageError)
// 	endif
// }
