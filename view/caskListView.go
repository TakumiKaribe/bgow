package view

import (
	"fmt"
	"os/exec"

	"github.com/jroimartin/gocui"
)

type CaskList struct {
	name      string
	title     string
	cmdOutput []byte
}

func NewCaskList() *CaskList {
	fl := &CaskList{
		name:  "cask list view",
		title: " Cask List ",
	}
	return fl
}

func (cl *CaskList) Edit(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	switch {
	case ch != 0 && mod == 0:
		v.EditWrite(ch)
	case key == gocui.KeySpace:
		v.EditWrite(' ')
	case key == gocui.KeyBackspace || key == gocui.KeyBackspace2:
		v.EditDelete(true)
	case key == gocui.KeyArrowLeft:
		v.MoveCursor(-1, 0, false)
		return
	case key == gocui.KeyArrowRight:
		v.MoveCursor(1, 0, false)
		return
	case key == gocui.KeyArrowUp:
		v.MoveCursor(0, -1, false)
		return
	case key == gocui.KeyArrowDown:
		v.MoveCursor(0, 1, false)
		return
	}
}

func (cl *CaskList) SetView(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(cl.name, 0, maxY/2, maxX/2-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		v.Frame = true
		v.Title = cl.title

		v.SelBgColor = gocui.ColorWhite
		v.FgColor = gocui.AttrBold | gocui.ColorWhite

		// TODO: goroutine
		cl.cmdOutput, _ = exec.Command("brew", "cask", "list").Output()
		fmt.Fprintln(v, string(cl.cmdOutput))
	}

	return nil
}
