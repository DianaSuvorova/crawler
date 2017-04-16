package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"flag"
)

var db *gorm.DB


type processor interface {
	url() (string)
	process() ([]processor)
}

func main() {
	var err error
	flag.Parse()
	connection := flag.Args()[0]
	db, err = gorm.Open("mysql", connection)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	entryUrl := "https://www.etsy.com/dynamic-sitemaps.xml?sitemap=browseindex"

	entryPage := newSiteMapMetaPage(entryUrl);

	queue := make(chan processor)
	filteredQueue := make(chan processor)

	go filterQueue(queue , filteredQueue);
	go func() {
		queue <- entryPage
	}()

	for page := range filteredQueue {
		enqueue(page, queue)
	}

}

func enqueue(page processor, queue chan processor) {
	pages := page.process();
		for _, addPage := range pages {
			go func(addPage processor) {
				queue <- addPage
			}(addPage)
	}
}

func filterQueue(in chan processor, out chan processor) {
	var seen = make(map[string]bool)
	for page := range in {
		url := page.url()
		fmt.Println(url)
		if (!seen[url]) {
			seen[url] = true
			out <- page
		}
	}
}
