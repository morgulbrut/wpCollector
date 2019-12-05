package wallpaperplay

import (
	"fmt"
	"strings"

	"github.com/morgulbrut/colorlog"

	"github.com/morgulbrut/helferlein"
	"github.com/morgulbrut/soup"
)

//Get all the wallpapers from wallpaperplay
func Get(query string, size string) {
	url := fmt.Sprintf("https://wallpaperplay.com/board/%s-wallpapers", query)
	resp, _ := soup.Get(url)
	doc := soup.HTMLParse(resp)
	imgs := doc.FindAll("div", "class", "flexbox_item")
	for _, d := range imgs {
		if d.Attrs()["data-or"] == size {
			path := d.Attrs()["data-fullimg"]
			pathsl := strings.Split(path, "/")
			filename := pathsl[len(pathsl)-1]
			colorlog.Debug("Downloading %s...", filename)
			err := helferlein.DownloadFile("https://wallpaperplay.com"+path, filename)
			if err != nil {
				colorlog.Error(err.Error())
			}
		}
	}
}
