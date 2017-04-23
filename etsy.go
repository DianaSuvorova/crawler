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

	queue := make(chan processor)
	filteredQueue := make(chan processor)

	go filterQueue(queue, filteredQueue)
	processFilteredQueue(filteredQueue, queue)

}

func processFilteredQueue(filteredQueue chan processor, queue chan processor) {
	closer := newQueueCloser(filteredQueue)

	entryUrl := "https://www.etsy.com/dynamic-sitemaps.xml?sitemap=taxonomyindex"
	entryPage := newSiteMapMetaPage(entryUrl);

	closer.increment()
	queue <- entryPage
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

func filterQueue(queue chan processor, filteredQueue chan processor) {
	var seen = make(map[string]bool)
	for page := range queue {
		url := page.url()
		if (!seen[url]) {
			seen[url] = true
			go func(page processor) {
				filteredQueue <- page
			}(page)
		}
	}
}
