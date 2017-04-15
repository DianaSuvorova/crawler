package main

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

func enqueue(page Processor, queue chan Processor) {
	pages := page.Process();
		for _, addPage := range pages {
			go func(addPage Processor) {
				queue <- addPage
			}(addPage)
	}
}

func filterQueue(in chan Processor, out chan Processor) {
	var seen = make(map[string]bool)
	for page := range in {
		url := page.Url()
		if (!seen[url]) {
			seen[url] = true
			out <- page
		}
	}
}
