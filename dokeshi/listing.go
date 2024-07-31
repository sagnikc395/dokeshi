package dokeshi

import (
	"fmt"
	"strings"
	"text/template"
)

//listing data will hold the data for the listing page

type ListingData struct {
	Title      string
	Date       string
	Short      string
	Link       string
	TimeToRead string
	Tags       []*Tag
}

// archivelinkdata will hold the data fro the archive link template
type ArchiveLinkData struct {
	NumPosts int
}

// configuration for the listing page
type ListingConfig struct {
	Posts                  []*Post
	SumAllPosts            int
	Template               *template.Template
	Destination, PageTitle string
	IsIndex                bool
	Writer                 *IndexWriter
}

type ListingGenerator struct {
	Config *ListingConfig
}

// add method to calculate time to read
func calculateTimeToRead(input string) string {
	timetoRead := 60.0 / 250.0

	words := timetoRead * float64(len(strings.Split(input, " ")))

	images := 5.0 * strings.Count(input, "<img")
	result := (words +float64(images)) / 60.0
	if result < 1.0 {
		result = 1.0
	}
	return fmt.Sprintf("%.0fm",result)
}
