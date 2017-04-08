package main

import (
  "net/http"
  "log"
	"bytes"
)

type Page struct {
    Url string
    Body string
}

func NewPage(url string) *Page {
  p := new(Page)
  p.Url = url
  resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
  p.Body = buf.String()
  defer resp.Body.Close()

  return p
}
