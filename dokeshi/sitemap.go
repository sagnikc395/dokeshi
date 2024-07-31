package dokeshi

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/beevik/etree"
)

// sitemapconfig holds the config for the sitemap
type SiteMapConfig struct {
	Posts       []*Post
	TagPostsMap map[string][]*Post
	Destination string
	BlogURL     string
	Statics     []string
}

// sitemap generator object
type SitemapGenerator struct {
	Config *SiteMapConfig
}

// generate methods create the sitemap of the blog
func (g *SitemapGenerator) Generate() error {
	fmt.Println("\tGenerating the sitemap 🏃‍♂️🏃‍♂️🏃‍♂️")
	posts := g.Config.Posts
	destination := g.Config.Destination
	tagPostsMap := g.Config.TagPostsMap
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	urlSet := doc.CreateElement("urlset")
	urlSet.CreateAttr("xmlns", "http://www.sitemaps.org/schemas/sitemap/0.9")
	urlSet.CreateAttr("xmlns:image", "http://www.google.com/schemas/sitemap-image/1.1")

	url := urlSet.CreateElement("url")
	loc := url.CreateElement("loc")
	loc.SetText(g.Config.BlogURL)
	urlTags := (g.Config.BlogURL + "/tags")

	for _, staticURL := range g.Config.Statics {
		addURL(urlSet, staticURL, g.Config.BlogURL, nil)
	}
	addURL(urlSet, "archive", g.Config.BlogURL, nil)
	addURL(urlSet, "tags", g.Config.BlogURL, nil)

	for tag := range tagPostsMap {
		addURL(urlSet, tag, urlTags, nil)
	}

	for _, post := range posts {
		addURL(urlSet, post.Name, g.Config.BlogURL, post.Images)
	}

	filePath := filepath.Join(destination, "sitemap.xml")
	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file %s: %v", filePath, err)
	}
	f.Close()
	if err := doc.WriteToFile(filePath); err != nil {
		return fmt.Errorf("error writing to file %s: %v", filePath, err)
	}
	fmt.Println("\tFinished generating Sitemap...")
	return nil
}

func addURL(element *etree.Element, location, blogURL string, images []string) {
	url := element.CreateElement("url")
	loc := url.CreateElement("loc")
	loc.SetText(fmt.Sprintf("%s/%s/", blogURL, location))

	if len(images) > 0 {
		for _, image := range images {
			img := url.CreateElement("image:image")
			imgLoc := img.CreateElement("image:loc")
			imgLoc.SetText(fmt.Sprintf("%s/%s/images/%s", blogURL, location, image))
		}
	}
}
