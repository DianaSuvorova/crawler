package main

import (
	"fmt"
)

type XmlLink struct {
  Url string `xml:"loc"`
}

func (l XmlLink) String() string {
  return fmt.Sprintf("%s", l.Url)
}
