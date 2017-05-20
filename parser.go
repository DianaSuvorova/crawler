package main

import (
  "runtime"
  "log"
)


func startParsing () {
  sources := []shopSource{}
  //db.Where("id in (?)", []int{1, 12}).Find(&sources)
  db.Find(&sources)
  println(sources, len(sources))
  sourcesChan := make(chan shopSource, len(sources))
  for  _, source := range sources {
    sourcesChan <- source
    wg.Add(1)
  }
  processPages(sourcesChan);
}

func processPages(sourcesChan chan shopSource) {
  for i := 0; i < 1; i++ {
    go func() {
      for source := range sourcesChan {
        page := newShopPage(source)

         var mem runtime.MemStats
         runtime.ReadMemStats(&mem)
         log.Println("sys mem MB: ",mem.Sys / (1024 * 1024))
         log.Println(source.Url)
         log.Println("numRoutines", runtime.NumGoroutine())
         page.process()
         wg.Done()
      }
    }()
  }
  wg.Wait()
}
