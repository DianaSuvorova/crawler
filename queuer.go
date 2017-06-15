package main

import (
  "runtime"
  "log"
)

func startQueuer(entryPage processor) {
  queue := make(chan processor)
	filteredQueue := make(chan processor, 1000000000)


  go filterQueue(queue, filteredQueue)
	processFilteredQueue(filteredQueue, queue, entryPage)

}

func processFilteredQueue(filteredQueue chan processor, queue chan processor, entryPage processor) {
  wg.Add(1)
  go func() {
    queue <- entryPage
  }()

  for i := 0; i < 99; i++ {
    go func() {
      for page := range filteredQueue {
        var mem runtime.MemStats
         runtime.ReadMemStats(&mem)
         log.Println("sys mem MB: ",mem.Sys / (1024 * 1024))
         log.Println(page.url())
         log.Println("numRoutines", runtime.NumGoroutine())
         log.Println("length of the filteredQueue", len(filteredQueue))
          pages := page.process()
        for _, addPage := range pages {

          queue <- addPage
          wg.Add(1)
        }
        wg.Done()
        }
    }()
  }
  wg.Wait()

}

func filterQueue(queue chan processor, filteredQueue chan processor) {
	var seen = make(map[string]bool)
  for {
    select {
      case page := <- queue:
        url := page.url()
        if (!seen[url]) {
          seen[url] = true
          filteredQueue <- page
        }
    }
  }
}
