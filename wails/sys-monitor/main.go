package main

import (
	"context"
	"embed"

	"github.com/labstack/gommon/log"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "sys-monitor",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		/**
		 * Just before the frontend is about to load index.html, a callback is made to the function provided in OnStartup. A standard Go context is passed to this method. This context is required when calling the runtime so a standard pattern is to save a reference to in this method. Just before the application shuts down, the OnShutdown callback is called in the same way, again with the context. There is also an OnDomReady callback for when the frontend has completed loading all assets in index.html and is equivalent of the body onload event in JavaScript. It is also possible to hook into the window close (or application quit) event by setting the option OnBeforeClose.
		 */
		OnStartup:     app.startup,
		OnDomReady:    onDomReady,
		OnBeforeClose: onBeforeClose,
		OnShutdown:    onShutdown,

		/**
				 * The Bind option is one of the most important options in a Wails application. It specifies which struct methods to expose to the frontend. Think of structs like "controllers" in a traditional web application. When the application starts, it examines the struct instances listed in the Bind field in the options, determines which methods are public (starts with an uppercase letter) and will generate JavaScript versions of those methods that can be called by the frontend code.
				 * Make sure you create an instance of it and pass it in Bind.
		     * Here we are binding the instance of the App struct we created earlier. This means that all public methods on the App struct will be available to call from the frontend. For example, the Greet method and the LogConsoleInfo method will be available to call from the frontend.
		*/
		Bind: []interface{}{
			app,
		},
		EnumBind: []interface{}{
			AllWeekdays,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

func onShutdown(ctx context.Context) {
	log.Infof("Shutdown event called.")
}
func onBeforeClose(ctx context.Context) (prevent bool) {
	log.Infof("Window close event called.")
	return false
}
func onDomReady(ctx context.Context) {
	log.Infof("DomReady event called.")
}
