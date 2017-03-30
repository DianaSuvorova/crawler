package main

import (
	"fmt"
  "net/http"
  "golang.org/x/net/html"
  "regexp"
)

type Items struct {
  url string
  items string
}

func Crawl(url string, ch chan Items, chFinished chan bool) {
	resp, _ := http.Get(url)

  b := resp.Body
  defer func() {
    b.Close()
    chFinished <- true
  }()

  z := html.NewTokenizer(b)

  for {
    token := z.Next()

    switch {
      case token == html.ErrorToken:
        // End of the document, we're done
        return
      case token == html.TextToken:
        text := string(z.Text())
        match, _ := regexp.MatchString("\\([,0-9]* items\\)", text)
        if (match) {
          ch <- Items{url, text}
        }

    }
  }

}

func main() {
  seedUrls := []string{"https://www.etsy.com/c/clothing/womens-clothing", "https://www.etsy.com/c/clothing/mens-clothing", "https://www.etsy.com/c/clothing/boys-clothing"}

  ch := make(chan Items);
  chFinished := make(chan bool)

  for _, url := range seedUrls {
    go Crawl(url, ch, chFinished)
  }

  for c := 0; c < len(seedUrls); {
    select {
    case items := <- ch:
        fmt.Println(items)
      case <- chFinished:
        c++
    }
  }

}
