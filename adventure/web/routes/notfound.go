package routes

import (
    "net/http"

    "gophercise/adventure/utils"
)

func NotFoundHandler (w http.ResponseWriter, _ *http.Request) {
    tmpl := MakeTemplate("404.html")
    w.WriteHeader(http.StatusNotFound)
    err := tmpl.Execute(w, nil)
    utils.CheckError(err)
}
