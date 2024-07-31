package dokeshi

import (
	"fmt"
	"html/template"
	"path/filepath"
	"sort"
	"time"
)

// metadata container
type Meta struct {
	Title      string
	Short      string
	Date       string
	Tags       []string
	ParsedDate time.Time
	Canonical  string
}

//data container for the landing page -> IndexData

type IndexData struct {
	HTMLTitle       string
	PageTitle       string
	Content         template.HTML
	Year            int
	Name            string
	CanonicalLink   string
	MetaDescription string
	HighlightCSS    template.CSS
}

type Generator interface {
	Generate() error
}

type SiteGenerator struct {
	Config *SiteConfig
}

// siteconfig holds the sources and destination folder
type SiteConfig struct {
	Sources     []string
	Destination string
	Config      *Config
}

// create a new sitegenerator
func New(config *SiteConfig) *SiteGenerator {
	return &SiteGenerator{Config: config.Config}
}

// generate the start of the static blog
func (g *SiteGenerator) Generate() error {
	templatePath := filepath.Join("static", "template.html")
	fmt.Println("Generating Site 🏃‍♂️🏃‍♂️🏃‍♂️")
	sources := g.Config.Sources
	destination := g.Config.Destination

	if err := clearAndCreateDestination(destination); err != nil {
		return err
	}

	if err := clearAndCreateDestination(filepath.Join(destination, "archive")); err != nil {
		return err
	}

	t, err := getTemplate(templatePath)
	if err != nil {
		return err
	}

	var posts []*Post
	for _, path := range sources {
		post, err := newPost(path, g.Config.Config.Blog.Datefmt)
		if err != nil {
			return err
		}
		posts = append(posts, post)
	}
	sort.Sort(ByDateDesc(posts))
	if err := runTasks(posts, t, destination, g.Config.Config); err != nil {
		return err
	}
	fmt.Println("⚡️ Finished generating the site.")
	return nil
}
