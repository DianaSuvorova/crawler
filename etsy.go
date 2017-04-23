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

	go filterQueue(queue, filteredQueue, todoQueue);
	go closeQueue(filteredQueue, todoQueue);
	go func() {
		todoQueue <- 1
		queue <- entryPage
	}()

	for page := range filteredQueue {
		enqueue(page, queue, todoQueue)
		todoQueue <- -1
	}
}


func closeQueue(filteredQueue chan processor, todoQueue chan int) {
	todo:=0
	for i := range todoQueue {
		todo += i
		if (todo == 0) {
			close(filteredQueue)
		}
	}
}

func enqueue(page processor, queue chan processor, todoQueue chan int) {
	pages := page.process()
	for _, addPage := range pages {
		todoQueue <- 1
		go func(addPage processor) {
			queue <- addPage
		}(addPage)
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
