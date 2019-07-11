package view

import (
	"fmt"
	"os/exec"

	"github.com/jroimartin/gocui"
)

type SearchResult struct {
	name      string
	title     string
	cmdOutput []byte
}

func NewSearchResult() *SearchResult {
	sr := &SearchResult{
		name:  "search result view",
		title: " Search Result ",
	}
	return sr
}

func (sr *SearchResult) Edit(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
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

func (sr *SearchResult) SetView(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(sr.name, maxX/2, 3, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		v.Frame = true
		v.Title = sr.title

		v.SelBgColor = gocui.ColorWhite
		v.FgColor = gocui.AttrBold | gocui.ColorWhite

		// TODO: goroutine
		sr.cmdOutput, _ = exec.Command("brew", "search", "go").Output()
		fmt.Fprintln(v, string(sr.cmdOutput))
	}

	return nil
}
