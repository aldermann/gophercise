package routes

import (
    "html/template"
    "path"
)

func p(f string) string {
    return path.Join("web/templates/", f)

}

func MakeTemplate(file string) *template.Template {
    return template.Must(template.ParseFiles(p("layout.html"), p(file)))
}
