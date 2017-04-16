package main

import (
	"fmt"
)

type xmlLink struct {
  Url string `xml:"loc"`
}

func (l xmlLink) String() string {
  return fmt.Sprintf("%s", l.Url)
}
