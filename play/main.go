package main

import (
	"time"

	"github.com/gotk3/gotk3/gtk"
)

type app struct {
	window    *gtk.Window
	actionBar *gtk.ActionBar
	headerBar *gtk.HeaderBar
	infoBar   *gtk.InfoBar
	popover   *gtk.Popover
}

func main() {
	gtk.Init(nil)
	builder, err := gtk.BuilderNew()
	if err != nil {
		panic(err)
	}
	builder.AddFromFile("ui.glade")
	var app app

	obj, err := builder.GetObject("window1")
	if err != nil {
		panic(err)
	}
	if w, ok := obj.(*gtk.Window); ok {
		app.window = w
	}
	app.window.Connect("destroy", gtk.MainQuit)

	obj, err = builder.GetObject("headerbar1")
	if err != nil {
		panic(err)
	}
	if h, ok := obj.(*gtk.HeaderBar); ok {
		app.headerBar = h
	}

	obj, err = builder.GetObject("infobar1")
	if err != nil {
		panic(err)
	}
	if i, ok := obj.(*gtk.InfoBar); ok {
		app.infoBar = i
	}
	app.infoBar.Connect("response", app.infoBar.Hide)

	obj, err = builder.GetObject("actionbar1")
	if err != nil {
		panic(err)
	}
	if i, ok := obj.(*gtk.ActionBar); ok {
		app.actionBar = i
	}

	//obj, err = builder.GetObject("popover1")
	//if err != nil {
	//	panic(err)
	//}
	//if i, ok := obj.(*gtk.Popover); ok {
	//	app.popover = i
	//}
	app.popover, err = gtk.PopoverNew(app.headerBar)
	if err != nil {
		panic(err)
	}

	go func() {
		for i := 0; ; i++ {
			<-time.After(2 * time.Second)
			toggle(app.infoBar)
			toggle(app.actionBar)
		}
	}()
	app.window.Show()
	gtk.Main()
}

type hideShower interface {
	Hide()
	Show()
	IsVisible() bool
}

func toggle(h hideShower) {
	if h.IsVisible() {
		h.Hide()
	} else {
		h.Show()
	}
}
