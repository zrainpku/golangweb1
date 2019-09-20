package main

import (
    "net/http"
    //"github.com/ziutek/mymysql/mysql"
    //_ "github.com/ziutek/mymysql/thrsafe"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "encoding/json"
    "log"
    "fmt"
)

type Result struct{
    Ret int
    Reason string
    Data interface{}
}

type ajaxController struct {
}


func (this *ajaxController)RegeditAction(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("content-type", "application/json")
    err := r.ParseForm()
    if err != nil {
        OutputJson(w, 0, "参数错误", nil)
        return
    }
    
    idcard := r.FormValue("idcard")
    mobilephone := r.FormValue("mobilephone")
    username := r.FormValue("username")
    school := r.FormValue("school")
    subschool := r.FormValue("subschool")
    major := r.FormValue("major")
    jobcode1 := r.FormValue("jobcode1")
    jobname1 := r.FormValue("jobname1")
    jobcode2 := r.FormValue("jobcode2")
    jobname2 := r.FormValue("jobname2")
    

    OutputJson(w, 0, idcard, nil)
    db,err := sql.Open("mysql","root:zhurui123@tcp(127.0.0.1:3306)/test")
    if err != nil {
     panic(err.Error())
}
    defer db.Close()
    
   stmt,err :=db.Prepare(`INSERT into userinfo (id_card,mobile_phone,user_name,school,sub_school,major,job_code1,job_name1,job_code2,job_name2) VALUES (?,?,?,?,?,?,?,?,?,?)`) 
    if err != nil {
     panic(err.Error())
}
   res,err :=stmt.Exec(idcard,mobilephone,username,school,subschool,major,jobcode1,jobname1,jobcode2,jobname2) 
    if err != nil {
     panic(err.Error())
}
   id,err := res.LastInsertId()
    if err != nil {
     panic(err.Error())
}
   fmt.Println(id)
   stmt.Close()

    OutputJson(w, 1, "操作成功", nil)
    return
}

func (this *ajaxController)LoginAction(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("content-type", "application/json")
    err := r.ParseForm()
    if err != nil {
        OutputJson(w, 0, "参数错误", nil)
        return
    }
    
    admin_name := r.FormValue("admin_name")
    admin_password := r.FormValue("admin_password")
    
    if admin_name == "" || admin_password == ""{
        OutputJson(w, 0, "参数错误", nil)
        return
    }
    
    //db := mysql.New("tcp", "", "192.168.100.166", "root", "test", "webdemo")
    db,err := sql.Open("mysql","root:zhurui123@tcp(127.0.0.1:3306)/test")
    if err != nil {
     panic(err.Error())
}
    /*if err := db.Connect(); err != nil {
        log.Println(err)
        OutputJson(w, 0, "数据库操作失败", nil)
        return
    }*/
    defer db.Close()
    //rows, res, err := db.Query("select * from webdemo_admin where admin_name = '%s'", admin_name)
    rows, err := db.Query("select * from webdemo_admin where admin_name = ? ", admin_name)
    if err != nil {
        log.Println(err)
        OutputJson(w, 0, "数据库操作失败了啊！", nil)
        return
    }
   var admin_password_db,admin_name_db string
   var admin_id_db int
   for rows.Next() {
    err := rows.Scan(&admin_id_db,&admin_name_db,&admin_password_db)
    if err != nil{
        log.Println(err)
        return
    }
    //fmt.Println("search mysql completed!")
} 
    if admin_password_db != admin_password {
        OutputJson(w, 0, "密码输入错误", nil)
        return
   } 
   // fmt.Println("password correct!")
    //name := res.Map("admin_password")
    //admin_password_db := rows[0].Str(name)
    
    /*if admin_password_db != admin_password {
        OutputJson(w, 0, "密码输入错误", nil)
        return
    }*/
    
    // 存入cookie,使用cookie存储
    //cookie := http.Cookie{Name: "admin_name", Value: rows[0].Str(res.Map("admin_name")), Path: "/"}
    //http.SetCookie(w, &cookie)
    cookie := http.Cookie{Name: "admin_name", Value: admin_password_db, Path: "/"}
    http.SetCookie(w, &cookie)
    
    OutputJson(w, 1, "操作成功", nil)
    return
}

func OutputJson(w http.ResponseWriter, ret int, reason string, i interface{}) {
    out := &Result{ret, reason, i}
    b, err := json.Marshal(out)
    if err != nil {
        return
    }
    w.Write(b)
}
