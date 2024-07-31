package dokeshi

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"sort"
	"strings"
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
	return &SiteGenerator{Config: config}
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

func getTemplate(path string) (*template.Template, error) {
	t, err := template.ParseFiles(path)
	if err != nil {
		return nil, fmt.Errorf("❌ Error reading template %s: %v", path, err)
	}
	return t, nil
}

func clearAndCreateDestination(path string) error {
	if err := os.RemoveAll(path); err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("❌ Error removing fodler at destination %s: %v", path, err)
		}
	}
	return os.Mkdir(path, os.ModePerm)
}

//indexwriter will create the index.html files

type IndexWriter struct {
	BlogTitle       string
	BlogDescription string
	BlogAuthor      string
	BlogURL         string
}

// writeIndexHTML will write the index.html file
func (i *IndexWriter) WriteIndexHTML(path, pageTitle, metaDscp string, content template.HTML,
	t *template.Template, canonicalLink string) error {

	filepath := filepath.Join(path, "index.html")
	f, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("❌ Error creating file %s: %v", filepath, err)
	}
	defer f.Close()

	metaDesc := metaDscp
	if metaDesc == "" {
		metaDesc = i.BlogDescription
	}

	hlBuffer := bytes.Buffer{}
	hlw := bufio.NewWriter(&hlBuffer)
	formatter.WriteCSS(hlw, styles.Lovelace)
	hlw.Flush()
	w := bufio.NewWriter(f)

	if canonicalLink == "" {
		canonicalLink = buildCanonicalLink(path, i.BlogURL)
	}

	td := IndexData{
		Name:            i.BlogAuthor,
		Year:            time.Now().Year(),
		HTMLTitle:       getHTMLTitle(pageTitle, i.BlogTitle),
		PageTitle:       pageTitle,
		Content:         content,
		CanonicalLink:   canonicalLink,
		MetaDescription: metaDesc,
		HighlightCSS:    template.CSS(hlbuf.String()),
	}
	if err := t.Execute(w, td); err != nil {
		return fmt.Errorf("❌ Error executing template %s: %v", filepath, err)
	}

	if err := w.Flush(); err != nil {
		return fmt.Errorf("❌ Error writing file %s: %v", filepath, err)
	}

	return nil

}
func getHTMLTitle(pageTitle, blogTitle string) string {
	if pageTitle == "" {
		return blogTitle
	}
	return fmt.Sprintf("%s - %s", pageTitle, blogTitle)
}

func buildCanonicalLink(path, baseURL string) string {
	parts := strings.Split(path, "/")
	if len(parts) > 1 {
		return fmt.Sprintf("%s/%s/index.html", baseURL, strings.Join(parts[1:], "/"))
	}
	return "/"
}

func getNumsOfPagesOnFrontpage(posts []*Post, numposts int) int {
	if len(posts) < numposts {
		return len(posts)
	}
	return numposts
}

func runTasks(posts []*Post, t *template.Template, destination string,
	cfg *Config) error {

}
