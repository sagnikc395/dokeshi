package dokeshi

import (
	"html/template"
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
