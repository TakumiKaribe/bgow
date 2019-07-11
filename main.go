package main

import (
	"log"
	"os"
	"os/exec"

	"bgow/view"

	"github.com/fatih/color"
	"github.com/jroimartin/gocui"
)

func checkAvailable() {
	_, err := exec.Command("brew", "help").Output()

	if err != nil {
		color.New(color.FgWhite).Add(color.Bold).Add(color.BgRed).Print("\n ERROR ")
		color.New(color.FgRed).Add(color.Bold).Println(" Homebrew is not installed. Please execute the following command to install.\n")
		color.New(color.FgYellow).Add(color.Bold).Println("/usr/bin/ruby -e \"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)\"")
		os.Exit(1)
	}
}

func main() {
	checkAvailable()

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}

	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	views := []view.View{view.NewFormulaList(), view.NewCaskList(), view.NewSearchField(), view.NewSearchResult()}
	for _, v := range views {
		v.SetView(g)
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
