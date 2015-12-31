package main

import (
	"time"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
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

	var infobar *gtk.InfoBar
	obj, err = builder.GetObject("infobar1")
	if err != nil {
		panic(err)
	}
	if i, ok := obj.(*gtk.InfoBar); ok {
		infobar = i
	}
	infobar.Connect("response", infobar.Hide)

	var actionbar *gtk.Actionbar
	obj, err = builder.GetObject("actionbar1")
	if err != nil {
		panic(err)
	}
	if i, ok := obj.(*gtk.Actionbar); ok {
		actionbar = i
	}
	//infobar.Connect("response", infobar.Hide)

	go func() {
		for i := 0; ; i++ {
			<-time.After(2 * time.Second)
			if i%2 == 0 {
				infobar.Hide()
			} else {
				infobar.Show()
			}
		}
	}()
	window.Show()
	gtk.Main()
}
