package main

import (
    "net/http"
    "strings"
    "reflect"
    "log"
    "html/template"
)


func jobHandler(w http.ResponseWriter, r *http.Request) {
    // 获取cookie
    cookie, err := r.Cookie("admin_name")
    if err != nil || cookie.Value == ""{
        http.Redirect(w, r, "/login/index", http.StatusFound)
    }
    
    pathInfo := strings.Trim(r.URL.Path, "/")
    parts := strings.Split(pathInfo, "/")
    var action = ""
    if len(parts) > 1 {
        action = strings.Title(parts[1]) + "Action"
    }
    
    job := &jobController{}
    controller := reflect.ValueOf(job)
    method := controller.MethodByName(action)
    if !method.IsValid() {
        method = controller.MethodByName(strings.Title("index") + "Action")
    }
    requestValue := reflect.ValueOf(r)
    responseValue := reflect.ValueOf(w)
    userValue := reflect.ValueOf(cookie.Value)
    method.Call([]reflect.Value{responseValue, requestValue, userValue})
}


func adminHandler(w http.ResponseWriter, r *http.Request) {
    // 获取cookie
    cookie, err := r.Cookie("admin_name")
    if err != nil || cookie.Value == ""{
        http.Redirect(w, r, "/login/index", http.StatusFound)
    }
    
    pathInfo := strings.Trim(r.URL.Path, "/")
    parts := strings.Split(pathInfo, "/")
    var action = ""
    if len(parts) > 1 {
        action = strings.Title(parts[1]) + "Action"
    }
    
    admin := &adminController{}
    controller := reflect.ValueOf(admin)
    method := controller.MethodByName(action)
    if !method.IsValid() {
        method = controller.MethodByName(strings.Title("index") + "Action")
    }
    requestValue := reflect.ValueOf(r)
    responseValue := reflect.ValueOf(w)
    userValue := reflect.ValueOf(cookie.Value)
    method.Call([]reflect.Value{responseValue, requestValue, userValue})
}

func ajaxHandler(w http.ResponseWriter, r *http.Request) {
    pathInfo := strings.Trim(r.URL.Path, "/")
    parts := strings.Split(pathInfo, "/")
    var action = ""
    if len(parts) > 1 {
        action = strings.Title(parts[1]) + "Action"
    }

    ajax := &ajaxController{}
    controller := reflect.ValueOf(ajax)
    method := controller.MethodByName(action)
    if !method.IsValid() {
        method = controller.MethodByName(strings.Title("index") + "Action")
    }
    requestValue := reflect.ValueOf(r)
    responseValue := reflect.ValueOf(w)
    method.Call([]reflect.Value{responseValue, requestValue})
}

func ajax2Handler(w http.ResponseWriter, r *http.Request) {
    pathInfo := strings.Trim(r.URL.Path, "/")
    parts := strings.Split(pathInfo, "/")
    var action = ""
    if len(parts) > 1 {
        action = strings.Title(parts[1]) + "Action"
    }

    ajax2 := &ajaxController{}
    controller := reflect.ValueOf(ajax2)
    method := controller.MethodByName(action)
    if !method.IsValid() {
        method = controller.MethodByName(strings.Title("index") + "Action")
    }
    requestValue := reflect.ValueOf(r)
    responseValue := reflect.ValueOf(w)
    method.Call([]reflect.Value{responseValue, requestValue})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("loginHandler")
    pathInfo := strings.Trim(r.URL.Path, "/")
    parts := strings.Split(pathInfo, "/")
    var action = ""
    if len(parts) > 1 {
        action = strings.Title(parts[1]) + "Action"
    }

    login := &loginController{}
    controller := reflect.ValueOf(login)
    method := controller.MethodByName(action)
    if !method.IsValid() {
        method = controller.MethodByName(strings.Title("index") + "Action")
    }
    requestValue := reflect.ValueOf(r)
    responseValue := reflect.ValueOf(w)
    method.Call([]reflect.Value{responseValue, requestValue})
}

func regeditHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("regeditHandler")
    pathInfo := strings.Trim(r.URL.Path, "/")
    parts := strings.Split(pathInfo, "/")
    var action = ""
    if len(parts) > 1 {
        action = strings.Title(parts[1]) + "Action"
    }

    regedit := &regeditController{}
    controller := reflect.ValueOf(regedit)
    method := controller.MethodByName(action)
    if !method.IsValid() {
        method = controller.MethodByName(strings.Title("index") + "Action")
    }
    requestValue := reflect.ValueOf(r)
    responseValue := reflect.ValueOf(w)
    method.Call([]reflect.Value{responseValue, requestValue})
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/" {
        http.Redirect(w, r, "/login/index", http.StatusFound)
    }
    
    t, err := template.ParseFiles("template/html/404.html")
    if (err != nil) {
        log.Println(err)
    }
    t.Execute(w, nil)
}
