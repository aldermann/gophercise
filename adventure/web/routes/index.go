package routes

import (
    "net/http"

    "gophercise/adventure/utils"
)

func IndexHandler (w http.ResponseWriter, _ *http.Request) {
    tmpl := MakeTemplate("index.html")
    err := tmpl.Execute(w, nil)
    utils.CheckError(err)
}
