package main

import (
    "net/http"
    "html/template"
    "log"
)

type regeditController struct {
}

func (this *regeditController)IndexAction(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("template/html/regedit/index.html")
    if (err != nil) {
        log.Println(err)
    }
    t.Execute(w, nil)
}
