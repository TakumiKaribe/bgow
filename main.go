package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/jroimartin/gocui"
)

func main() {
	_, err := exec.Command("brew", "help").Output()

	if err != nil {
		color.New(color.FgWhite).Add(color.Bold).Add(color.BgRed).Print("\n ERROR ")
		color.New(color.FgRed).Add(color.Bold).Println(" Homebrew is not installed. Please execute the following command to install.\n")
		color.New(color.FgYellow).Add(color.Bold).Println("/usr/bin/ruby -e \"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)\"")
		os.Exit(1)
	}

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
	maxX, maxY := g.Size()

	if v, err := g.SetView("list", 0, 0, maxX/2-1, maxY/2-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		out, _ := exec.Command("brew", "list").Output()
		fmt.Fprintln(v, string(out))
	}

	if v, err := g.SetView("cask list", 0, maxY/2, maxX/2-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		out, _ := exec.Command("brew", "cask", "list").Output()
		fmt.Fprintln(v, string(out))
	}

	if v, err := g.SetView("searchField", maxX/2, 0, maxX-1, 2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "")
	}

	if v, err := g.SetView("searchResult", maxX/2, 3, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		// out, _ := exec.Command("brew", "search", "go").Output()
		// fmt.Fprintln(v, string(out))
		fmt.Fprintln(v, "")
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
