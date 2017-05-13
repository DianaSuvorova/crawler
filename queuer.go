package main

import (
  "crawler/queueCloser"
)

func startQueuer(entryPage processor) {
  queue := make(chan processor, 10000)
	filteredQueue := make(chan processor, 10000)

	go filterQueue(queue, filteredQueue)
	processFilteredQueue(filteredQueue, queue, entryPage)

}

func processFilteredQueue(filteredQueue chan processor, queue chan processor, entryPage processor) {
  closer := queueCloser.NewQueueCloser()

	closer.Increment()
	queue <- entryPage
  for {
    select {
      case page := <- filteredQueue:
        pages := page.process()
        for _, addPage := range pages {
          go func (addPage processor) {
            println(addPage.url())
            closer.Increment()
            queue <- addPage
            println("todo: ", closer.Todo())
          }(addPage);
        }
        closer.Decrement()
      case <- closer.Quit:
        close(filteredQueue)
        return;
    }
  }
}

func filterQueue(queue chan processor, filteredQueue chan processor) {
	var seen = make(map[string]bool)
  for {
    select {
      case page := <- queue:
        url := page.url()
        if (!seen[url]) {
          seen[url] = true
          go func(page processor) {
            filteredQueue <- page
          }(page)
        }
    }
  }
}
