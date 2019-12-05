package wallpapersite

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/morgulbrut/colorlog"
	"github.com/morgulbrut/helferlein"
	"github.com/morgulbrut/soup"
)

//Get all the wallpapers from wallpapersite
func Get(query string, size string, numPages int) {
	var re = regexp.MustCompile(`(?m)[0-9]{4}x[0-9]{4}`) // cuts out resolution

	for i := 1; i <= numPages; i++ {
		url := fmt.Sprintf("https://wallpapersite.com/wallpaper/%s?page=%d", query, i)
		resp, _ := soup.Get(url)
		doc := soup.HTMLParse(resp)
		pics := doc.Find("div", "id", "pics-list").Children()

		for _, p := range pics {
			imgPath := p.Children()[0].Attrs()["href"]
			imgURL := fmt.Sprintf("https://wallpapersite.com/%s", imgPath)
			imgResp, _ := soup.Get(imgURL)
			imgDoc := soup.HTMLParse(imgResp)
			dlPath := re.ReplaceAllString(imgDoc.Find("a", "class", "original").Attrs()["href"], size)
			pathsl := strings.Split(dlPath, "/")
			filename := pathsl[len(pathsl)-1]
			colorlog.Debug("Downloading %s...", filename)
			err := helferlein.DownloadFile("https://wallpapersite.com/"+dlPath, filename)
			if err != nil {
				colorlog.Error(err.Error())
			}
		}
	}
}
