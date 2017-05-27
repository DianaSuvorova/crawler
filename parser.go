package main

import (
  "runtime"
  "log"
)


func startParsing (runId uint) {
  sources := []shopSource{}
  //  db.Where("Id < ?", 222).Find(&sources)
  //db.Find(&sources)
  db.Find(&sources)
  //db.Where("deleted IS NULL and joined IS NULL").Find(&sources)
  println("sources", len(sources))
  sourcesChan := make(chan shopSource, len(sources))
  for  _, source := range sources {
    sourcesChan <- source
    wg.Add(1)
  }
  processPages(sourcesChan, runId);
}

func processPages(sourcesChan chan shopSource, runId uint) {
  for i := 0; i < 99; i++ {
    go func() {
      for source := range sourcesChan {
        page := newShopPage(source, runId)

         var mem runtime.MemStats
         runtime.ReadMemStats(&mem)
         log.Println("sys mem MB: ",mem.Sys / (1024 * 1024))
         log.Println(source.Url)
         log.Println("numRoutines", runtime.NumGoroutine())
         log.Println("remaining", len(sourcesChan))
         page.process()
         wg.Done()
      }
    }()
  }
  wg.Wait()
}
