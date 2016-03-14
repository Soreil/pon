package main

import (
	"os"

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
	buttons          map[string]*gtk.Button
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

	//Handle quit signal
	app.window.Connect("destroy", gtk.MainQuit)
	//Handle collapsing indicator
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
		for _, v := range app.buttons {
			v.SetSensitive(!v.GetSensitive())
		}
		toggle(app.actionBar)
	})

	//go func() {
	//	for {
	//		<-time.After(2 * time.Second)
	//		toggle(app.infoBar)
	//		toggle(app.actionBar)
	//	}
	//}()
	app.window.Show()
	gtk.Main()
}

func GetObjects(builderFile string) (app, error) {
	builder, err := gtk.BuilderNew()
	if err != nil {
		return app{}, err
	}
	builder.AddFromFile(builderFile)
	var app app
	app.buttons = make(map[string]*gtk.Button)

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
	obj, err = builder.GetObject("moves")
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

	for _, name := range []string{"kan", "pon", "riichi", "tsumo", "ron", "chi"} {
		obj, err = builder.GetObject(name)
		if err != nil {
			return app, err
		}
		if i, ok := obj.(*gtk.Button); ok {
			app.buttons[name] = i
		}
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
