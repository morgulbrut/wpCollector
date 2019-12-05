package wallhaven

import (
	"fmt"
	"strings"

	"github.com/morgulbrut/colorlog"
	"github.com/morgulbrut/helferlein"
	"github.com/morgulbrut/soup"
)

//
//
//

//Get all the wallpapers from wallhaven
func Get(query string, size string, numPages int) {

	for i := 1; i <= numPages; i++ {
		url := fmt.Sprintf("https://wallhaven.cc/search?q=%s&purity=110&resolutions=%s", query, size)
		resp, _ := soup.Get(url)
		doc := soup.HTMLParse(resp)
		pics := doc.FindAll("a", "class", "preview")
		for _, p := range pics {
			imgPath := p.Attrs()["href"]
			imgResp, _ := soup.Get(imgPath)
			imgDoc := soup.HTMLParse(imgResp)
			if imgDoc.Find("img", "id", "wallpaper").Error == nil {
				dlPath := imgDoc.Find("img", "id", "wallpaper").Attrs()["src"]
				pathsl := strings.Split(dlPath, "/")
				filename := pathsl[len(pathsl)-1]
				colorlog.Debug("Downloading %s...", filename)
				err := helferlein.DownloadFile(dlPath, filename)
				if err != nil {
					colorlog.Error(err.Error())
				}
			} else {
				colorlog.Fatal("Error: reading %s", imgPath)
			}
		}
	}
}
