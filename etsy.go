package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"flag"
	"crawler/queueCloser"
)

var db *gorm.DB


type processor interface {
	url() (string)
	process() ([]processor)
}


func main() {
	var err error
	flag.Parse()
	// connection := flag.Args()[0]
	connection := "boris:B@ckspace123@tcp(54.215.211.253:3306)/etsy"
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
	closer := queueCloser.NewQueueCloser()

	entryUrl := "https://www.etsy.com/dynamic-sitemaps.xml?sitemap=taxonomyindex"
	entryPage := newSiteMapMetaPage(entryUrl);

	closer.Increment()
	queue <- entryPage
	go func () {
		quit := <- closer.Quit
		if (quit) {
			close(filteredQueue)
		}
	}()


	for page := range filteredQueue {
		pages := page.process()
		for _, addPage := range pages {
			closer.Increment()
			go func(addPage processor) {
				queue <- addPage
		 }(addPage)
		}
		closer.Decrement()
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
