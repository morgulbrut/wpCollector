package main

import (
	"fmt"
	"os"

	"github.com/morgulbrut/color256"

	"github.com/akamensky/argparse"
	"github.com/morgulbrut/colorlog"
	"github.com/morgulbrut/wpCollector/provider/wallhaven"
	"github.com/morgulbrut/wpCollector/provider/wallpaperplay"
	"github.com/morgulbrut/wpCollector/provider/wallpapersite"
)

func main() {
	logo()
	pr := argparse.NewParser("wallpaperscraper", "downloads a bunch of wallpapers from different providers")
	query := pr.String("q", "query", &argparse.Options{Required: true, Help: "Keyword (required)"})
	size := pr.String("s", "size", &argparse.Options{Required: true, Help: "Size, format 1920x1080 (required)"})
	wpp := pr.Flag("1", "wallpaperplay", &argparse.Options{Required: false, Help: "Download from wallpaperplay"})
	wps := pr.Flag("2", "wallpapersite", &argparse.Options{Required: false, Help: "Download from wallpapersite"})
	wh := pr.Flag("3", "wallhaven", &argparse.Options{Required: false, Help: "Download from wallhaven"})
	err := pr.Parse(os.Args)
	if err != nil {
		fmt.Print(pr.Usage(err))
		os.Exit(1)
	}

	colorlog.SetLogLevel(colorlog.DEBUG)

	if *wpp {
		colorlog.Info("Scraping wallpaperplay...")
		wallpaperplay.Get(*query, *size)
	}
	if *wps {
		colorlog.Info("Scraping wallpapersite...")
		wallpapersite.Get(*query, *size, 10)
	}
	if *wh {
		colorlog.Info("Scraping wallhaven...")
		wallhaven.Get(*query, *size, 10)
	}
}

func logo() {
	logo := `
               ._________        .__  .__                __               
 __  _  _______ \_   ___ \  ____ |  | |  |   ____  _____/  |_ ___________ 
 \ \/ \/ \____ \/    \  \/ /  _ \|  | |  | _/ __ _/ ___\   __/  _ \_  __ \
  \     /|  |_> \     \___(  <_> |  |_|  |_\  ___\  \___|  |(  <_> |  | \/
   \/\_/ |   __/ \______  /\____/|____|____/\___  \___  |__| \____/|__|   
         |__|           \/                      \/    \/                 `

	color256.PrintRed(logo)
}
