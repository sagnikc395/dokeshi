package dokeshi

import (
	"log"
)

func RunApp() {
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
