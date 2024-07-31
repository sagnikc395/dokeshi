package dokeshi

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
)

// post holds the data for a post
type Post struct {
	Name      string
	HTML      []byte
	Meta      *Meta
	ImagesDir string
	Images    []string
}

// postconfig will hold the post's configuration
type PostConfig struct {
	Post        *Post
	Destination string
	Template    *template.Template
	Writer      *IndexWriter
}

// bydatedesc -> sorting object for posts
type ByDateDesc []*Post

// post generator object
type PostGenerator struct {
	Config *PostConfig
}

// generaete a post
func (g *PostGenerator) Generate() error {
	post := g.Config.Post
	destination := g.Config.Destination
	t := g.Config.Template

	fmt.Printf("\tGenerating the Post : %s...\n", post.Meta.Title)
	staticPath := filepath.Join(destination, post.Name)
	if err := os.Mkdir(staticPath, os.ModePerm); err != nil {
		return fmt.Errorf("Error creating directory at %s: %v", staticPath, err)
	}

	if post.ImagesDir != "" {
		if err := copyImagesDir(post.ImagesDir, staticPath); err != nil {
			return nil
		}
	}

	if err := g.Config.Writer.WriteIndexHTML(staticPath, post.Meta.Title, post.Meta.Short, template.HTML(string(post.HTML)), t, post.Meta.Canonical); err != nil {
		return err
	}

	fmt.Printf("\tFinished generating Post: %s ... \n", post.Meta.Title)

	return nil
}

func newPost(path, dateFormat string) (*Post, error) {
	meta, err := getMeta(path, dateFormat)
	if err != nil {
		return nil, err
	}
	html, err := getHTML(path)
	if err != nil {
		return nil, err
	}
	imagesDir, images, err := getImages(path)
	if err != nil {
		return nil, err
	}
	name := filepath.Base(path)

	return &Post{Name: name, Meta: meta, HTML: html, ImagesDir: imagesDir, Images: images}, nil
}

func getImages(path string) (string, []string, error) {
	dirPath := filepath.Join(path, "images")
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil, nil
		}
		return "", nil, fmt.Errorf("Error while reading folder %s: %v", dirPath, err)
	}
	images := []string{}
	for _, file := range files {
		images = append(images, file.Name())
	}
	return dirPath, images, nil
}

// utility methods
func (p ByDateDesc) Len() int {
	return len(p)
}

func (p ByDateDesc) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p ByDateDesc) Less(i, j int) bool {
	return p[i].Meta.ParsedDate.After(p[j].Meta.ParsedDate)
}
