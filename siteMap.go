package main

import (
	"encoding/xml"
)

type siteMapPage struct {
 *page
 *siteMap
}

type siteMap struct {
 Links []xmlLink `xml:"url"`
}

func newSiteMapPage(url string) *siteMapPage {
  smp := new(siteMapPage)
  smp.page = newPage(url)

  return smp
}

func (smp *siteMapPage)process() (pages []processor) {
	success := smp.page.fetch();
	if (success) {
		xml.Unmarshal([]byte(smp.page.body), &smp.siteMap)
		for _, link := range smp.Links {
			page := newListingPage(link.String())
			pages = append(pages, page)
		}

		// if (len(smp.Links) > 0) {
		// 	link := smp.Links[0].String()
		// 	page := newListingPage(link)
		// 	// page := newCategoryPage(link)
		// 	pages = append(pages, page)
		// }
	}
  return
}

func (smp *siteMapPage)url() string {
	return smp.page.url
}
