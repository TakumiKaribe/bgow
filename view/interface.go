package view

import "github.com/jroimartin/gocui"

type View interface {
	Edit(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier)
	SetView(g *gocui.Gui) error
}
