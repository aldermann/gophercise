package web

import (
    "fmt"
    "log"
    "net/http"

    "gophercise/adventure/story"
    "gophercise/adventure/web/routes"
)

func Run() {
    http.Handle("/static/", http.FileServer(http.Dir("web/templates/")))

    ArcList := story.ParseStory()

    http.HandleFunc("/ending", routes.EndingHandler)

    for k, arc := range *ArcList {
        fmt.Println(k, arc.Title)
        http.HandleFunc(fmt.Sprintf("/%s", k), routes.MakeArcHandler(arc))
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.RequestURI == "/" {
            routes.IndexHandler(w, r)
        } else {
            routes.NotFoundHandler(w, r)
        }
    })

    log.Println("Started server")
    log.Fatal(http.ListenAndServe(":3000", nil))
}
