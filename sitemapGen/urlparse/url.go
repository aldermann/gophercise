package urlparse

import (
    "errors"
    "net/url"
    "regexp"

    "gophercise/sitemapGen/utils"
)

var rootUrl *url.URL

func trimIndex(s string) string {
    rg := regexp.MustCompile("/index\\.[a-zA-Z]+$")
    return string(rg.ReplaceAll([]byte(s), []byte("")))
}

func RegisterRootUrl(path string) *url.URL {
    var err error
    rootUrl, err = url.Parse(path)
    if err != nil {
        utils.CheckError(err, false)
        return nil
    }
    if rootUrl.Host == "" {
        utils.CheckError(errors.New("no host specified. try specifying both the host and a scheme"), true)
    }
    rootUrl.Path = trimIndex(rootUrl.Path)
    if rootUrl.Scheme == "" {
        rootUrl.Scheme = "http"
    }
    if rootUrl.Path == "" {
        rootUrl.Path = "/"
    }
    return rootUrl
}

func Parse(rawpath string) (*url.URL, error) {
    u, err := url.Parse(rawpath)
    if err != nil {
        utils.CheckError(err, false)
        return nil, err
    }
    if u.Host != "" && u.Host != rootUrl.Host {
        return nil, errors.New("path not in the same domain")
    }
    return rootUrl.ResolveReference(u), nil
}
