package dokeshi

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func Run() {
	cfg, err := configRead()
	if err != nil {
		log.Fatal("configuration file format incorrect , please configure right config", err)
	}
	ds := datareader.New()
	dirs, err := ds.Fetch(cfg)

	if err != nil {
		log.Fatal(err)
	}

	g := generator.New(&generator.SiteConfig{
		Sources:     dirs,
		Destination: cfg.Generator.Dest,
		Config:      cfg,
	})

	err := g.Generate()

	if err != nil {
		log.Fatal(err)
	}
}

// to start the local HTTP server for dev/testing purposes
func Serve() {
	cwd, cwderr := os.Getwd()
	if cwderr != nil {
		log.Fatal(cwderr)
	}

	staticPath := filepath.Join(cwd, "www")
	log.Print("Serving from directory %s", staticPath)
	log.Println("HTTP server listening on 127.0.0.1:8000")
	fileSystemDir := http.FileServer(http.Dir(staticPath))
	http.Handle("/", fileSystemDir)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func readConfig() (*Config, error) {
	data, err := os.ReadFile("dokeshi-config.yml")
	if err != nil {
		return nil, fmt.Errorf("Not able to read from config file: %v", err)
	}

	cfg := Config{}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("could not parse config: %v", err)
	}

	if cfg.Generator.Repository == "" {
		return nil, fmt.Errorf("Please provide a repository URL, e.g.: https://github.com/sagnikc395/dokeshi")
	}

	if cfg.Generator.Temp == "" {
		cfg.Generator.Temp = "tmp"
	}

	if cfg.Generator.Dest == "" {
		cfg.Generator.Dest = "www"
	}
	if cfg.Blog.URL == "" {
		return nil, fmt.Errorf("Please provide a Blog URL, e.g.: https://www.zupzup.org")
	}
	if cfg.Blog.Lang == "" {
		cfg.Blog.Lang = "en-us"
	}
	if cfg.Blog.Descp == "" {
		return nil, fmt.Errorf("Please provide a Blog Description, e.g.: A blog about Go, JavaScript, Open Source and Programming in General")
	}
	if cfg.Blog.Datefmt == "" {
		cfg.Blog.Datefmt = "02.01.2006"
	}
	if cfg.Blog.Title == "" {
		return nil, fmt.Errorf("Please provide a Blog Title, e.g.: zupzup")
	}
	if cfg.Blog.Author == "" {
		return nil, fmt.Errorf("Please provide a Blog author, e.g.: Mario Zupan")
	}
	if cfg.Blog.Frontpage == 0 {
		cfg.Blog.Frontpage = 10
	}

	return &cfg, nil
}
