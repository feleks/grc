package main

import (
	"fmt"
	"github.com/getlantern/systray"
	"log"
	"net/http"
	"os"
)

func main() {
	err := prepareAssets()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to run main: %w", err))
	}
	systray.Run(onReady, onExit)
}

func onReady() {
	iconImage, err := getAsset("assets/icon.ico")
	if err != nil {
		log.Fatal(err)
	}

	systray.SetIcon(iconImage)
	systray.SetTitle("grc")
	systray.SetTooltip("Golang [web] remote control tool")

	mQuit := systray.AddMenuItem("Quit", "Quit the app")

	run()

	for {
		select {
		case <-mQuit.ClickedCh:
			os.Exit(0)
		}
	}
}

func onExit() {
	// clean up here
}

func run() {
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// HttpServer
	go func() {
		setupRoutes()
		err := http.ListenAndServe(":8971", nil)
		if err != nil {
			log.Fatal(fmt.Errorf("failed to run http server: %w", err))
		}
	}()
}
