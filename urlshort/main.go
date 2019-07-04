package main

import (
    "log"
    "net/http"

    "gophercise/urlshort/utils"
    "gophercise/urlshort/yamlHandler"
)

func main() {
    filePath := utils.ParseFlag()
    var handler RedirectHandler = yamlHandler.YAMLHandler{FilePath: filePath, FallbackHandler: utils.DefaultMux()}
    log.Println("Started server")
    log.Fatal(http.ListenAndServe(":3000", handler.MakeHandler()))
}

type RedirectHandler interface {
    MakeHandler() (mux *http.ServeMux)
}
