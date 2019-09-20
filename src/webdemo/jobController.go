package main

import (
    "log"
    "net/http"
    "html/template"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)


type JobsMysql struct {
     Job_name,Job_code,Job_nums,Major,Education  string
     Others string
}

type jobController struct {
}


func (this *jobController)IndexAction(w http.ResponseWriter, r *http.Request, user string) {
    t, err := template.ParseFiles("template/html/job/index.html")
    if (err != nil) {
        log.Println(err)
    }
    db,err := sql.Open("mysql","root:zhurui123@tcp(127.0.0.1:3306)/test")
    if err != nil {
     panic(err.Error())
     }

    defer db.Close()
    rows, err := db.Query("select * from jobtable ")
    if err != nil {
        log.Println(err)
        OutputJson(w, 0, "数据库操作失败了啊！", nil)
        return
    }
    
    job := JobsMysql{}
    jobs := []JobsMysql{}

    for rows.Next() {
        var job_name, job_code, job_nums, major,education string 
        var others string
        var id int
        err = rows.Scan(&id, &job_name, &job_code, &job_nums, &major, &education, &others)
        checkErr(err)
        job.Job_name = job_name
        job.Job_code = job_code
        job.Job_nums = job_nums
        job.Major = major
        job.Education = education
        job.Others = others
        jobs = append(jobs, job)
    }
    t.Execute(w, &jobs)
}
