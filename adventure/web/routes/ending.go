package routes

import (
    "net/http"

    "gophercise/adventure/utils"
)

func EndingHandler (w http.ResponseWriter, _ *http.Request) {
    tmpl := MakeTemplate("ending.html")
    err := tmpl.Execute(w, nil)
    utils.CheckError(err)
}
