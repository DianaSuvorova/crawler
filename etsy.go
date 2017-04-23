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

	entryUrl := "https://www.etsy.com/dynamic-sitemaps.xml?sitemap=taxonomyindex"

	entryPage := newSiteMapMetaPage(entryUrl);

	queue := make(chan processor)
	filteredQueue := make(chan processor)
	todoQueue := make(chan int)

	go filterQueue(queue, filteredQueue, todoQueue)
	closer := newQueueCloser(filteredQueue)
	go func() {
		closer.increment()
		queue <- entryPage
	}()

	for page := range filteredQueue {
		pages := page.process()
		for _, addPage := range pages {
			closer.increment()
			go func(addPage processor) {
				queue <- addPage
			}(addPage)
		}
		closer.decrement()
	}
}

func filterQueue(in chan processor, out chan processor, todoQueue chan int) {
	var seen = make(map[string]bool)
	for page := range in {
		url := page.url()
		if (!seen[url]) {
			seen[url] = true
			out <- page
		}
	}
}
