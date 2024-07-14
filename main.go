package main

// TODO: Figure out caching and storage
// TODO: Different row colors for different things
//	Cyan:	new since last open/refresh
//	Red:	unwatched
//	Green:	watched
//	Grey:	hidden
// TODO: Channel management
// TODO: Playlist management
// TODO: Some sort of actual visualization

import (
	osUser "os/user"

	"dalton.dog/YouTerm/models"
	"dalton.dog/YouTerm/modules/API"
	"dalton.dog/YouTerm/modules/Storage"
	"dalton.dog/YouTerm/modules/TUI"

	"github.com/charmbracelet/log"
)

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// NOTE:
// Probably should also maintain some sort of "Username -> ID" mapping. The user definitely won't just have the Channel ID readily available

// TODO:
// Search by username, generate Channel struct
// Save Channel struct to database, using its ID as the key
// Also update the username -> ID map
// Add a few channels by username, list them out, and then be able to load them between executions

func main() {
	Storage.Startup()
	defer Storage.Shutdown()

	API.InitializeManager()

	curUser, err := osUser.Current()
	checkErr(err)

	user := models.LoadOrCreateUser(curUser.Username)
	defer Storage.SaveResource(user) // This ensures any changes to the user get closed when program closes

	program := TUI.NewPromptProgram(user)

	if _, err := program.Run(); err != nil {
		log.Fatal(err)
	}
	// fmt.Println("Launching API test")
	// ctx := context.Background()
	//
	// service := API.GetService(ctx)
	//
	// fmt.Println("Service successfully created")
	//
	// videoIDs := API.GetUploadsForChannel(service, "", "Northernlion", "")
	//
	// for _, video := range videoIDs {
	// 	API.PrintInfoForVideo(service, video)
	// }

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
