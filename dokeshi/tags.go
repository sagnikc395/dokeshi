package dokeshi

import (
	"fmt"
	"strings"
	"text/template"
)

// tag holds the data for a Tag
type Tag struct {
	Name  string
	Link  string
	Count int
}

//bycountdescn sorts the tags

type ByCountDesc []*Tag

// config
type TagsConfig struct {
	TagPostsMap map[string][]*Post
	Template    *template.Template
	Destination string
	Writer      *IndexWriter
}

// tags-generator object
type TagsGenerator struct {
	Config *TagsConfig
}

// utils
func getTagLink(tag string) string {
	return fmt.Sprintf("/tags/%s/", strings.ToLower(tag))
}

func (t ByCountDesc) Len() int {
	return len(t)
}

func (t ByCountDesc) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t ByCountDesc) Less(i, j int) bool {
	return t[i].Count > t[j].Count
}
