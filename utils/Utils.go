package utils

import (
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/charmbracelet/x/term"
)

func CheckErrFatal(err error, errMsg string) {
	CheckErr(err, errMsg, true)
}

func CheckErr(err error, errMsg string, fatal bool) {
	if err != nil {
		log.Printf("%v: %v", errMsg, err)
		if fatal {
			os.Exit(1)
		}
	}
}

func GetTerminalSize() (int, int, error) {
	width, height, err := term.GetSize(os.Stdin.Fd())
	if err != nil {
		return 0, 0, err
	}
	return width, height, nil
}

func LaunchURL(url string) error {
	log.Printf("Attempting to launch URL %v\n", url)
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
		args = []string{url}
	default:
		if isWSL() {
			cmd = "cmd.exe"
			args = []string{"/c", "start", url}
		} else {
			cmd = "xdg-open"
			args = []string{url}
		}
	}

	if len(args) > 1 {
		args = append(args[:1], append([]string{""}, args[1:]...)...)
	}
	return exec.Command(cmd, args...).Start()
}

func isWSL() bool {
	data, err := exec.Command("uname", "-r").Output()
	if err != nil {
		return false
	}
	return strings.Contains(strings.ToLower(string(data)), "microsoft")
}
