package main

import (
    "flag"
    "fmt"
    "math"
    "net/http"
    "net/url"
    "strings"
    "sync"
    "sync/atomic"

    "gophercise/sitemapGen/link"
    "gophercise/sitemapGen/sitemap"
    "gophercise/sitemapGen/urlparse"
    "gophercise/sitemapGen/utils"
)

var maxDepth *int
var wg sync.WaitGroup
var cnt int32 = 0
func main() {
    root := flag.String("r", "", "The root of your website")
    maxDepth = flag.Int("d", math.MaxInt32, "Max depth of crawl tree. Default to crawl all")
    flag.Parse()
    rootUrl := urlparse.RegisterRootUrl(*root)
    wg.Add(1)
    crawl(rootUrl, 0)
    wg.Wait()
    sitemap.WriteFile()
    fmt.Printf("crawled %d links", cnt)
}

func crawl(currentUrl *url.URL, currentDepth int) {
    defer wg.Done()
    st := sitemap.GetSiteMap()
    if st.IsVisited(currentUrl.Path) {
        return
    }
    fmt.Printf("Fetching %s\n", currentUrl.String())
    resp, err := http.Get(currentUrl.String())
    atomic.AddInt32(&cnt, 1)
    if err != nil {
        utils.CheckError(err, false)
        return
    }
    if !strings.Contains(resp.Header.Get("content-type"), "text/html") {
        return
    }
    st.MarkVisited(currentUrl.Path)
    if currentDepth == *maxDepth {
        return
    }
    linksLists, err := link.Parse(resp.Body)
    if err != nil {
        utils.CheckError(err, false)
    }
    for _, ln := range linksLists {
        adjUrl, err := urlparse.Parse(ln.Href)
        if err != nil {
            continue
        }
        st.MarkEdge(currentUrl.Path, adjUrl.Path)
        wg.Add(1)
        go crawl(adjUrl, currentDepth + 1)
    }
}
