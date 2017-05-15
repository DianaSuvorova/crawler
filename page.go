package main

import (
  "net/http"
	"bytes"
)

type page struct {
  url string
  body string
  fetched chan bool
}

func newPage(url string) *page {
  p := new(page)
  p.url = url

  return p
}

func (p *page) fetch() (bool) {
  done := make(chan bool);
  go func () {
    resp, err := http.Get(p.url)
    if err != nil {
      panic(err)
      done <- false
    } else {
      buf := new(bytes.Buffer)
      buf.ReadFrom(resp.Body)
      p.body = buf.String()
      defer resp.Body.Close()
      done <- true
    }
  }();

  for {
		select {
    case success := <- done:
        return success;
    }
  }
}
