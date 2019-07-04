package mapHandler

import (
    "log"
    "net/http"
    "gophercise/urlshort/utils"
)

type MapHandler struct {
    PathsToUrls     map[string]string
    FallbackHandler http.Handler
}

func (m MapHandler) MakeHandler() (mux *http.ServeMux) {
    mux = http.NewServeMux()
    for k, v := range m.PathsToUrls {
        log.Printf("This path %s will be mapped to this path %s\n", k, v)
        mux.HandleFunc(k, utils.NewRedirectHandler(v))
    }
    mux.Handle("/", m.FallbackHandler)
    return mux
}
