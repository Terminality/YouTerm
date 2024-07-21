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
	"os"
	osUser "os/user"
	"runtime/pprof"

	"github.com/google/uuid"

	"dalton.dog/YouTerm/modules/API"
	"dalton.dog/YouTerm/modules/Storage"
	"dalton.dog/YouTerm/modules/TUI"
	"dalton.dog/YouTerm/resources"
	"dalton.dog/YouTerm/utils"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Flag Parsing and Doing Stuff
	clearAll := flag.Bool("clear-all", false, "Clear channel and video buckets, and user lists")
	clearChannels := flag.Bool("clear-channel", false, "Clear channel bucket")
	clearVideos := flag.Bool("clear-videos", false, "Clear video bucket")

	shouldProfile := flag.Bool("prof", false, "Profile program run")
	// shouldDebug := flag.Bool("debug", false, "Enable debugging output and logging")

	flag.Parse()

	if *clearAll || *clearChannels {
		Storage.ClearBucket(Storage.CHANNELS)
	}
	if *clearAll || *clearVideos {
		Storage.ClearBucket(Storage.VIDEOS)
	}

	if *shouldProfile {
		file, err := os.Create("YouTerm.prof")
		utils.CheckErrFatal(err, "Couldn't open profiling file")
		pprof.StartCPUProfile(file)
		defer pprof.StopCPUProfile()
	}

	f, err := tea.LogToFile("debug.log", "debug")
	utils.CheckErrFatal(err, "Couldn't open debug logging file")
	defer f.Close()
	// End admin parsing stuff

	execUUID := uuid.NewString()

	log.Printf("\n~~~~~ New YouTerm Execution %v ~~~~~", execUUID)
	defer log.Printf("~~~~~ End YouTerm Execution %v ~~~~~\n", execUUID)

	// Startup and ensure shutdown of database
	Storage.Startup()
	defer Storage.Shutdown()

	// Initialize the API manager, create context/server connection
	err = API.InitializeManager()
	utils.CheckErrFatal(err, "Unable to initialize API manager")

	// Load User stuff from the OS and the database
	var user *resources.User
	curUser, err := osUser.Current()
	utils.CheckErrFatal(err, "Unable to load current user")
	user = resources.LoadOrCreateUser(curUser.Username)

	// This ensures any changes to the user get closed when program closes
	defer Storage.SaveResource(user)

	err = TUI.RunProgram(user)
	utils.CheckErrFatal(err, "Error encountered while running program")
}
