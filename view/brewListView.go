package view

import (
	"fmt"
	"os/exec"

	"github.com/jroimartin/gocui"
)

type FormulaList struct {
	name      string
	title     string
	cmdOutput []byte
}

func NewFormulaList() *FormulaList {
	fl := &FormulaList{
		name:  "formula list view",
		title: " Formula List ",
	}
	return fl
}

func (fl *FormulaList) Edit(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
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

func (fl *FormulaList) SetView(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(fl.name, 0, 0, maxX/2-1, maxY/2-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		v.Frame = true
		v.Title = fl.title

		v.SelBgColor = gocui.ColorWhite
		v.FgColor = gocui.AttrBold | gocui.ColorWhite

		// TODO: goroutine
		fl.cmdOutput, _ = exec.Command("brew", "list").Output()
		fmt.Fprintln(v, string(fl.cmdOutput))
	}

	return nil
}
