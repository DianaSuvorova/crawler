package main

import (
	"fmt"
)

type processor interface {
	url() (string)
	process() ([]processor)
}

func main() {
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
