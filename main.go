package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
	"github.com/morgulbrut/colorlog"
	"github.com/morgulbrut/wpCollector/provider/wallhaven"
	"github.com/morgulbrut/wpCollector/provider/wallpaperplay"
	"github.com/morgulbrut/wpCollector/provider/wallpapersite"
)

func main() {
	pr := argparse.NewParser("wallpaperscraper", "downloads a bunch of wallpapers from different providers")
	query := pr.String("q", "query", &argparse.Options{Required: true, Help: "Keyword"})
	size := pr.String("s", "size", &argparse.Options{Required: true, Help: "Size, format 1920x1080"})
	wpp := pr.Flag("1", "wallpaperplay", &argparse.Options{Required: false, Help: "Download from wallpaperplay"})
	wps := pr.Flag("2", "wallpapersite", &argparse.Options{Required: false, Help: "Download from wallpapersite"})
	wh := pr.Flag("3", "wallhaven", &argparse.Options{Required: false, Help: "Download from wallhaven"})
	err := pr.Parse(os.Args)
	if err != nil {
		fmt.Print(pr.Usage(err))
		os.Exit(1)
	}
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
