package main

import (
  "net/http"
  "log"
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

func (p *page) fetch() {
  done := make(chan bool);
  go func () {
    resp, err := http.Get(p.url)
    if err != nil {
      log.Fatal(err)
    }
    buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    p.body = buf.String()
    defer resp.Body.Close()
    done <- true
  }();

  for {
		select {
      case  <- done:
        return;
    }
  }
}
