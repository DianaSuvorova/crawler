package main

import (
	"fmt"
)


func enqueue(page Processor, queue chan Processor) {
	pages := page.Process();

	for _, addPage := range pages {
		fmt.Println("it is supposed to add", addPage.Url())
		go func () { queue <- addPage } ()
	}
}

type Processor interface {
	Url() (string)
	Process() ([]Processor)
}

func main() {
	entryUrl := "https://www.etsy.com/dynamic-sitemaps.xml?sitemap=browseindex"

	entryPage := NewSiteMapMetaPage(entryUrl);

	queue := make(chan Processor)
	filteredQueue := make(chan Processor)

	go filterQueue(queue , filteredQueue);
	go func() {
		queue <- entryPage
	}()

	for page := range filteredQueue {
		enqueue(page, queue)
	}
}

func filterQueue(in chan Processor, out chan Processor) {
	var seen = make(map[string]bool)
	fmt.Println("seen", seen)
	for page := range in {
		if (!seen[page.Url()]) {
			fmt.Println("adding to seen", page.Url())
			seen[page.Url()] = true
			out <- page
		} else {
			 fmt.Println("not adding to seen", page.Url())
		 }
	}
}
