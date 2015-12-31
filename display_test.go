package pon

import (
	"testing"

	"github.com/gotk3/gotk3/gtk"
)

func TestDisplay(t *testing.T) {
	gtk.Init(nil)
	builder, err := gtk.BuilderNew()
	if err != nil {
		panic(err)
	}
	builder.AddFromFile("ui.glade")

	var window *gtk.Window
	obj, err := builder.GetObject("window1")
	if err != nil {
		panic(err)
	}
	if w, ok := obj.(*gtk.Window); ok {
		window = w
	}
	window.Connect("destroy", gtk.MainQuit)
	window.Show()
	gtk.Main()
	t.Fatal("Window closed")
}
