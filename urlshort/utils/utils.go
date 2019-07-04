package utils

import (
    "fmt"
    "log"
    "net/http"
)

func DefaultMux() *http.ServeMux {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
        fmt.Fprintln(w, "hello world")
    })
    return mux
}

func NewRedirectHandler(dest string) (res http.HandlerFunc) {
    res = func(w http.ResponseWriter, r *http.Request) {
        http.Redirect(w, r, dest, http.StatusMovedPermanently)
    }
    return
}

func CheckErr (err error) {
    if err != nil {
        log.Fatal(err)
    }
}
