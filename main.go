package main

// TODO: Probably should also maintain some sort of "Username -> ID" mapping

// TODO: Different row colors for different things
//	Cyan:	new since last open/refresh
//	Red:	unwatched
//	Green:	watched
//	Grey:	hidden

// TODO: Channel management

// TODO: Playlist management

import (
	"flag"
	osUser "os/user"

	"dalton.dog/YouTerm/modules/API"
	"dalton.dog/YouTerm/modules/Storage"
	"dalton.dog/YouTerm/modules/TUI"
	"dalton.dog/YouTerm/resources"

	"github.com/charmbracelet/log"
)

// Utility function for error checking
func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// Check `debug` flag and appropriately set Debug level printing
func checkDebug() {
	debugFlagPtr := flag.Bool("debug", false, "Enable debug logging")
	flag.Parse()
	if *debugFlagPtr {
		log.SetLevel(log.DebugLevel)
	}
	log.Debug("Debug logging enabled")
}

func main() {
	checkDebug()

	Storage.Startup()
	defer Storage.Shutdown()

	log.Debug("Storage startup complete")

	API.InitializeManager()

	log.Debug("API startup complete")

	curUser, err := osUser.Current()
	checkErr(err)
	user := resources.LoadOrCreateUser(curUser.Username)

	log.Debug("Successfully loaded user", "user", user.ID)

	defer Storage.SaveResource(user) // This ensures any changes to the user get closed when program closes

	// program := TUI.NewPromptProgram(user)
	program := TUI.MakeNewProgram(user)

	log.Debug("Program successfully started")

	if _, err := program.Run(); err != nil {
		log.Fatal(err)
	}
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
