package dokeshi

import (
	"fmt"
	"time"

	"github.com/beevik/etree"
)

// rss generator config holds the configuration for an RSS feed
type RSSConfig struct {
	Posts           []*Post
	Destination     string
	DateFormat      string
	Language        string
	BlogURL         string
	BlogDescription string
	BlogTitle       string
}

type RSSGenerator struct {
	Config *RSSConfig
}

const rssDateFormat string = "01 Aug 2024 15:04 +0530"

func addItem(element *etree.Element, post *Post, path, dateFormat string) error {
	meta := post.Meta
	item := element.CreateElement("item")

	item.CreateElement("title").SetText(meta.Title)
	item.CreateElement("link").SetText(path)
	item.CreateElement("guid").SetText(path)

	pubDate, err := time.Parse(dateFormat, meta.Date)
	if err != nil {
		return fmt.Errorf("Error parsing date %s:%v", meta.Date, err)
	}

	item.CreateElement("pubDate").SetText(pubDate.Format(rssDateFormat))
	item.CreateElement("description").SetText(string(post.HTML))

	return nil
}
