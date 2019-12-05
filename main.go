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
	err := pr.Parse(os.Args)
	if err != nil {
		fmt.Print(pr.Usage(err))
		os.Exit(1)
	}
	colorlog.Info("Scraping wallpaperplay...")
	wallpaperplay.Get(*query, *size)
	colorlog.Info("Scraping wallpapersite...")
	wallpapersite.Get(*query, *size, 10)
	colorlog.Info("Scraping wallhaven...")
	wallhaven.Get(*query, *size, 10)
}
