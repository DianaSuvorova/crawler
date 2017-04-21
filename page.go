package main

import (
  "net/http"
  "log"
	"bytes"
)

type page struct {
    url string
    body string
}

func newPage(url string) *page {
  p := new(page)
  p.url = url
  resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
  p.body = buf.String()
  defer resp.Body.Close()

  return p
}
