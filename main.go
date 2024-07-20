package main

//
// ⠀⣠⣶⣶⢿⡿⣿⢿⡿⣿⢿⡿⣿⢿⡿⣿⢿⡿⣿⢿⡿⣶⣶⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
// ⢰⣿⡽⣾⣻⣽⢯⣿⡽⣯⡿⣽⣯⢿⣽⢯⣿⡽⣯⡿⣽⡷⣯⢿⡇⠀⣿⣿⣿⣧⠀⠀⢀⣾⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢰⣶⣶⣶⣶⣶⣶⣶⣶⣶⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
// ⣼⣯⣟⣷⢿⣽⣻⣞⣿⣳⠛⢿⣾⣻⡽⣿⣞⡿⣽⣻⢷⣟⡿⣯⣿⠀⠈⢿⣿⣿⣧⢀⣾⣿⣿⡿⠁⢀⣀⣤⣤⣤⣄⡀⠀⢠⣤⣤⣤⠀⢠⣤⣤⣤⠈⠉⠉⠉⣿⣿⡏⠉⠉⠉⠀⠀⢀⣀⣀⣀⠀⠀⠀⢀⣀⡀⢀⣀⡀⢀⣀⡀⢀⣀⣀⡀⠀⢀⣀⣀⡀⠀
// ⣿⣾⣿⣾⣿⣷⣿⣿⣾⣿⠀⠀⠘⠻⣿⣷⣿⣿⣿⣿⣿⣾⣿⣷⣿⠀⠀⠘⢿⣿⣿⣿⣿⣿⡿⠁⢰⣿⣿⣿⣿⣿⣿⣿⣆⢸⣿⣿⣿⠀⢸⣿⣿⣿⠀⠀⠀⠀⣿⣿⡇⠀⠀⠀⢀⣾⣿⣿⢿⣿⣿⣆⠀⢸⣿⣷⣿⣿⡇⢸⣿⣿⡿⢿⣿⣿⣾⡿⢿⣿⣿⡆
// ⣿⣷⣻⢷⣯⣟⣾⣽⢾⡿⠀⠀⠀⠀⠀⢉⣾⣯⢿⣳⣯⢿⣽⣻⢿⠀⠀⠀⠈⢿⣿⣿⣿⡿⠁⠀⣿⣿⣿⡏⠀⠹⣿⣿⣿⢸⣿⣿⣿⠀⢸⣿⣿⣿⠀⠀⠀⠀⣿⣿⡇⠀⠀⠀⣸⣿⣟⣀⣀⣘⣿⣿⠀⢸⣿⣟⠁⠀⠀⢸⣿⡏⠀⠀⢹⣿⡏⠀⠀⢹⣿⡧
// ⣿⣷⣻⣟⣾⣽⣳⣯⢿⣻⠀⠀⣀⣴⣾⣟⡿⣞⡿⣯⣟⣯⡿⣽⣿⠀⠀⠀⠀⢸⣿⣿⣿⠁⠀⠀⣿⣿⣿⣇⠀⣠⣿⣿⣿⢸⣿⣿⣿⠀⢸⣿⣿⣿⠀⠀⠀⠀⣿⣿⡇⠀⠀⠀⢻⣿⣟⠛⠛⠛⠛⠛⠀⢸⣿⣯⠀⠀⠀⢸⣿⡇⠀⠀⢸⣿⡇⠀⠀⢸⣿⡇
// ⢸⡷⣟⣾⣻⢾⣽⢯⣟⡿⣶⢿⣻⣟⣾⣽⣻⢯⡿⣷⢯⡿⣽⡷⡿⠀⠀⠀⠀⢸⣿⣿⣿⠀⠀⠀⠹⣿⣿⣿⣿⣿⣿⣿⠏⠘⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⣿⣿⡇⠀⠀⠀⠈⠿⣿⣶⣦⣶⣾⡦⠀⢸⣿⣷⠀⠀⠀⢸⣿⡇⠀⠀⢸⣿⣇⠀⠀⢸⣿⣇
// ⠘⣿⣻⢷⣻⣯⣟⣯⡿⣽⣻⣯⣟⣾⣽⢾⣯⣟⣿⣽⣻⣽⡷⣿⠇⠀⠀⠀⠀⠘⠛⠛⠛⠀⠀⠀⠀⠈⠙⠛⠛⠛⠋⠁⠀⠀⠈⠛⠛⠛⠛⠛⠛⠛⠀⠀⠀⠀⠉⠉⠁⠀⠀⠀⠀⠀⠈⠉⠉⠉⠁⠀⠀⠈⠉⠉⠀⠀⠀⠈⠉⠁⠀⠀⠈⠉⠉⠀⠀⠈⠉⠉
// ⠀⠈⠛⠻⠽⠾⠽⠷⠿⠿⠷⠿⠾⠿⠾⠿⠾⠽⠾⠷⠿⠷⠛⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
//
// A program by Dalton Williams

// TODO: Channel management
// TODO: Playlist management
// TODO: Different row colors for different things
//	Cyan:	new since last open/refresh
//	Red:	unwatched
//	Green:	watched
//	Grey:	hidden

// BUG: Updates take a really long time to happen after a startup sometimes
//      Set up some profiling to see what's taking so long ASAP

import (
	"flag"
	"log"
	osUser "os/user"

	"github.com/google/uuid"

	"dalton.dog/YouTerm/modules/API"
	"dalton.dog/YouTerm/modules/Storage"
	"dalton.dog/YouTerm/modules/TUI"
	"dalton.dog/YouTerm/resources"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	checkErr(err)
	defer f.Close()

	execUUID := uuid.NewString()

	log.Printf("\n~~~~~ New YouTerm Execution %v ~~~~~", execUUID)
	defer log.Printf("~~~~~ End YouTerm Execution %v ~~~~~\n", execUUID)

	// Startup and ensure shutdown of database
	Storage.Startup()
	defer Storage.Shutdown()

	// Initialize the API manager, create context/server connection
	err = API.InitializeManager()
	checkErr(err)

	curUser, err := osUser.Current()
	checkErr(err)
	user := resources.LoadOrCreateUser(curUser.Username)

	// This ensures any changes to the user get closed when program closes
	defer Storage.SaveResource(user)

	err = TUI.RunProgram(user)
	checkErr(err)
}

// Utility function for error checking
func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func checkFlags() {
	clearAll := flag.Bool("clear-all", false, "Clear channel and video buckets, and user lists")
	clearChannels := flag.Bool("clear-channel", false, "Clear channel bucket")
	clearVideos := flag.Bool("clear-videos", false, "Clear video bucket")

	flag.Parse()

	if *clearAll || *clearChannels {
		Storage.ClearBucket(Storage.CHANNELS)
	}
	if *clearAll || *clearVideos {
		Storage.ClearBucket(Storage.VIDEOS)
	}
}
