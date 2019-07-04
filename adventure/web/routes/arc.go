package routes

import (
    "net/http"

    "gophercise/adventure/story"
    "gophercise/adventure/utils"
)

func MakeArcHandler(arc story.Arc) http.HandlerFunc {
    return func(writer http.ResponseWriter, request *http.Request) {
        tmpl := MakeTemplate("arc.html")
        err := tmpl.Execute(writer, arc)
        utils.CheckError(err)
    }
}
