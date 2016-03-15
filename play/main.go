package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/gotk3/gotk3/gdk"
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
	playerImages     map[int][]*gtk.Image
	playerAvatar     map[int]*gtk.Image
	images           map[string]*gdk.Pixbuf
}

func init() {
	gtk.Init(&os.Args)
	rand.Seed(time.Now().UnixNano())
}

func main() {
	const builderFile = "../ui/ui.glade"

	//Experimentation
	const tilePath = "../img/tiles"
	app, err := GetObjects(builderFile)
	if err != nil {
		panic(err)
	}

	app.images, err = GetImages(tilePath)
	if err != nil {
		panic(err)
	}

	for _, pixbuf := range app.images {
		app.playerImages[rand.Intn(3)][rand.Intn(12)].SetFromPixbuf(pixbuf)
	}
	//

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

	app.playerImages = make(map[int][]*gtk.Image)
	for player := 0; player < 4; player++ {
		app.playerImages[player] = make([]*gtk.Image, 13)
		for tile := 1; tile < 14; tile++ {
			obj, err = builder.GetObject("tile" + fmt.Sprint(tile+player*13))
			if err != nil {
				return app, err
			}
			if i, ok := obj.(*gtk.Image); ok {
				app.playerImages[player][tile-1] = i
			}
		}
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

func GetImages(path string) (map[string]*gdk.Pixbuf, error) {
	dir, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	names, err := dir.Readdirnames(0)
	if err != nil {
		return nil, err
	}
	images := make(map[string]*gdk.Pixbuf)
	for _, name := range names {
		images[name], err = gdk.PixbufNewFromFileAtScale(path+string(os.PathSeparator)+name, 32, 32, true)
		if err != nil {
			return images, err
		}
	}
	return images, nil
}
