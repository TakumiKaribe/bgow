package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
)

func main() {
	out, err := exec.Command("brew", "").Output()

	if err != nil {
		color.New(color.FgWhite).Add(color.Bold).Add(color.BgRed).Print("\n ERROR ")
		fontRed := color.New(color.FgRed).Add(color.Bold)
		fontRed.Println(" Homebrew is not installed. Please execute the following command to install.\n")
		color.New(color.FgYellow).Add(color.Bold).Println("/usr/bin/ruby -e \"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)\"")
		os.Exit(1)
	}

	fmt.Print(string(out))
}
