package view

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

type SearchField struct {
	name      string
	title     string
	cmdOutput []byte
}

func NewSearchField() *SearchField {
	sf := &SearchField{
		name:  "search field view",
		title: " Search Formula or Cask ",
	}
	return sf
}

func (s *SearchField) Edit(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
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

func (s *SearchField) SetView(g *gocui.Gui) error {
	maxX, _ := g.Size()
	if v, err := g.SetView(s.name, maxX/2, 0, maxX-1, 2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		v.Frame = true
		v.Title = s.title

		v.SelBgColor = gocui.ColorWhite
		v.FgColor = gocui.AttrBold | gocui.ColorWhite
		fmt.Fprintln(v, "")
	}

	return nil
}
