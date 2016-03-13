package main

import (
	"os"
	"time"

	"github.com/gotk3/gotk3/gtk"
)

type app struct {
	window           *gtk.Window
	actionBar        *gtk.ActionBar
	headerBar        *gtk.HeaderBar
	infoBar          *gtk.InfoBar
	menu             *gtk.Menu
	settings         *gtk.Settings
	isDarkTheme      *gtk.CheckMenuItem
	isActionsEnabled *gtk.CheckMenuItem
}

func init() {
	gtk.Init(&os.Args)
}

func main() {
	const builderFile = "../ui/ui.glade"
	app, err := GetObjects(builderFile)
	if err != nil {
		panic(err)
	}

	app.window.Connect("destroy", gtk.MainQuit)
	app.infoBar.Connect("response", app.infoBar.Hide)

	//Switch application theme
	app.isDarkTheme.Connect("toggled", func() {
		isDark, err := app.settings.GetProperty("gtk-application-prefer-dark-theme")
		if b, ok := isDark.(bool); ok {
			err = app.settings.Set("gtk-application-prefer-dark-theme", !b)
		}
		if err != nil {
			panic(err)
		}
	})

	//Enable action buttons
	app.isActionsEnabled.Connect("toggled", func() {
		//widget, err := app.actionBar.GetChild()
		//if err != nil {
		//	panic(err)
		//}
		//revealerChild, err := gtk.Bin(widget).GetChild()
		//if err != nil {
		//	panic(err)
		//}

		//name, err := revealerChild.GetName()
		//if err != nil {
		//	panic(err)
		//}

		//fmt.Println(name)
	})

	go func() {
		for {
			<-time.After(2 * time.Second)
			toggle(app.infoBar)
			toggle(app.actionBar)
		}
	}()
	app.window.Show()
	defer gtk.MainQuit()
	gtk.Main()
}

func GetObjects(builderFile string) (app, error) {
	builder, err := gtk.BuilderNew()
	if err != nil {
		return app{}, err
	}
	builder.AddFromFile(builderFile)
	var app app

	settings, err := gtk.SettingsGetDefault()
	if err != nil {
		return app, err
	}
	app.settings = settings

	obj, err := builder.GetObject("window1")
	if err != nil {
		return app, err
	}
	if w, ok := obj.(*gtk.Window); ok {
		app.window = w
	}

	obj, err = builder.GetObject("headerbar1")
	if err != nil {
		return app, err
	}
	if h, ok := obj.(*gtk.HeaderBar); ok {
		app.headerBar = h
	}
	obj, err = builder.GetObject("infobar1")
	if err != nil {
		return app, err
	}
	if i, ok := obj.(*gtk.InfoBar); ok {
		app.infoBar = i
	}
	obj, err = builder.GetObject("actionbar1")
	if err != nil {
		return app, err
	}
	if i, ok := obj.(*gtk.ActionBar); ok {
		app.actionBar = i
	}

	obj, err = builder.GetObject("menu1")
	if err != nil {
		return app, err
	}
	if i, ok := obj.(*gtk.Menu); ok {
		app.menu = i
	}

	obj, err = builder.GetObject("isDarkTheme")
	if err != nil {
		return app, err
	}
	if i, ok := obj.(*gtk.CheckMenuItem); ok {
		app.isDarkTheme = i
	}
	obj, err = builder.GetObject("isActionsEnabled")
	if err != nil {
		return app, err
	}
	if i, ok := obj.(*gtk.CheckMenuItem); ok {
		app.isActionsEnabled = i
	}
	return app, nil
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
