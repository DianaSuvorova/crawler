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
  resp, err := http.Get(p.url)
  if (err != nil) {
    return false;
  }
  buf := new(bytes.Buffer)
  buf.ReadFrom(resp.Body)
  p.body = buf.String()
  defer resp.Body.Close();
  return true;
}
