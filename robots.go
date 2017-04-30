package main

import (
    "regexp"
    "fmt"
    "reflect"
)

type robotsPage struct {
  *page
  *robots
}

type robots struct {
  shopListingSiteMaps []string
}

func newRobotsPage(url string) *robotsPage {
  rp := new(robotsPage)
  rp.page = newPage(url)
  rp.robots = new(robots)
  rp.getShopListingSiteMaps()
  return rp
}

func (rp *robotsPage) getShopListingSiteMaps() {
  re, err := regexp.Compile("Sitemap: (https:\\/\\/www.etsy.com\\/dynamic-sitemaps.xml\\?sitemap=shoplisting_index2.*)")
  if (err != nil) {
    fmt.Println(err)
  }
  res := re.FindAllStringSubmatch(rp.body, -1)
  fmt.Println(reflect.TypeOf(res[0]))
  for _, r := range res {
    rp.robots.shopListingSiteMaps = append(rp.robots.shopListingSiteMaps, r[1])
  }
}

func (rp *robotsPage) process() (siteMapMetaPages []processor) {
  for _, link := range rp.robots.shopListingSiteMaps {
    fmt.Println(link)
    siteMapMetaPages = append(siteMapMetaPages, newSiteMapMetaPage(link))
  }
  return
}

func (rp *robotsPage) url() string{
  return rp.page.url
}
