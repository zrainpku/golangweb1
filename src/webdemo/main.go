package main

import (
    "net/http"
    "log"
    //"tools/route"
)

func main() {
    log.Println("main")
    http.Handle("/css/", http.FileServer(http.Dir("template")))
    http.Handle("/js/", http.FileServer(http.Dir("template")))
    http.Handle("/music/", http.FileServer(http.Dir("template")))
    http.Handle("/bootstrap/", http.FileServer(http.Dir("template")))
    http.Handle("/img/", http.FileServer(http.Dir("template")))
    
    http.HandleFunc("/admin/", adminHandler)
    http.HandleFunc("/job/", jobHandler)
    http.HandleFunc("/login/",loginHandler)
    http.HandleFunc("/regedit/",regeditHandler)
    http.HandleFunc("/ajax/",ajaxHandler)
    http.HandleFunc("/ajax2/",ajax2Handler)
    http.HandleFunc("/",NotFoundHandler)
    http.ListenAndServe(":9090", nil)
}
