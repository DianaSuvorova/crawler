package main

import (
	"encoding/xml"
)

type siteMapMetaPage struct {
 *page
 *siteMapMeta
}

type siteMapMeta struct {
 Links []xmlLink `xml:"sitemap"`
}

func newSiteMapMetaPage(url string) *siteMapMetaPage {
  smmp := new(siteMapMetaPage)
	smmp.page = newPage(url)
  return smmp
}

func (smmp *siteMapMetaPage) process(availSpaceInQueue int) (siteMapPages []processor) {
	success := smmp.page.fetch();
	if success {
		if (len(smmp.siteMapMeta.Links) < availSpaceInQueue) {
			xml.Unmarshal([]byte(smmp.page.body), &smmp.siteMapMeta)
			for _, link := range smmp.siteMapMeta.Links {
				siteMap := newSiteMapPage(link.String())
				siteMapPages = append(siteMapPages, siteMap)
			}
		} else {
			siteMapPages = append(siteMapPages, smmp)
		}
	}
	return
}

func (smmp *siteMapMetaPage)url() string {
	return smmp.page.url
}
