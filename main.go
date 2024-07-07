package main

import (
	"dalton.dog/YouTerm/API"
	//"dalton.dog/YouTerm/TUI"
	// "dalton.dog/YouTerm/StorageIO"
	// "errors"
	// "fmt"
	// "os"
	//tea "github.com/charmbracelet/bubbletea"
	//"github.com/charmbracelet/log"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// TODO: Figure out where to handle channel management. Presently leaning toward its own module, or a ChannelManager thing in API
func main() {
	// apiKEY, _ := API.GetKeyFromEnv()
	// check(err)

	// fmt.Printf("Got the API Key: %s", apiKEY)

	//visuals.TestProgram()
	API.MainAPI()
}

func getAPIKey() (string, error) {
	return "", nil
}

// Snippets / Notes / References / Whatever

// ~ Environment Variables ~
// os.Getenv("VAR_NAME")
// os.Setenv("VAR_NAME", value)

// ~ CLI Arguments ~
// os.Args (includes the program call)
// os.Args[1:] (only includes the arguments)

// ~ Flags ~
// Set up all of the flags ahead of parsing. Use the `flag` package
// wordPtr := flag.String("opt", "default_val", "help text")
// flag.Parse()
// *wordPtr will contain `val` if program is called with `-opt=val`, else `default_val`

// Psuedocode
// Check env for API key. If not there, prompt user for it
// Display main table
//	Ensure there's an informative zero state
// Config keymaps at the bottom
//	Add to Watch Later
//	Channel Management
//	Filter
//	Open in Browser
//	Refresh
//	Settings
//	Toggle Watched
