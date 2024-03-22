package subscription

import (
    "time"
    "fmt"
    "net/http"
    "html/template"

    "go_authentication/options"
    "go_authentication/authtoken"
)


func ToUpUsSsc(w http.ResponseWriter, r *http.Request) {

    id,err := options.IdUrl(w,r)
    if err != nil {
        return
    }

    cls,err := authtoken.SqlToken(w,r)
    if cls == nil {
        return
    }
    if err != nil {
        return
    }

    flag,err := options.ParseBool(r.FormValue("completed"))
    if err != nil {
        fmt.Fprintf(w, "err ParseBool()..  : %+v\n", err)
        return
    }

    if r.Method == "GET" {

        i,err := userIdSsc(w,id)
        if err != nil {
            return
        }

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/ssc/update_ssc.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", i)
    }


    if r.Method == "POST" {

        sqlStatement := `UPDATE subscription SET completed=$2, updated_at=$3 WHERE id=$1;`
        
        _, err := db.Exec(sqlStatement, id,flag,time.Now())
        
        if err != nil {
            fmt.Fprintf(w, "err db.Exec()..! : %+v\n", err)
            return
        }
        http.Redirect(w, r, "/all-touser-ssc", http.StatusFound)
    }
}


func ToUpGroupSsc(w http.ResponseWriter, r *http.Request) {

    id,err := options.IdUrl(w,r)
    if err != nil {
        return
    }

    cls,err := authtoken.SqlToken(w,r)
    if cls == nil {
        return
    }
    if err != nil {
        return
    }

    flag,err := options.ParseBool(r.FormValue("completed"))

    if err != nil {
        fmt.Fprintf(w, "err ParseBool()..  : %+v\n", err)
        return
    }


    if r.Method == "GET" {

        owner := cls.User_id

        i,err := roomIdSsc(w,id,owner)
        if err != nil {
            return
        }

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/ssc/update_ssc.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", i)
    }


    if r.Method == "POST" {

        sqlStatement := `UPDATE subscription SET completed=$2, updated_at=$3 WHERE id=$1;`
        
        _, err := db.Exec(sqlStatement, id,flag,time.Now())
        
        if err != nil {
            fmt.Fprintf(w, "err db.Exec()..! : %+v\n", err)
            return
        }
        http.Redirect(w, r, "/all-toroom-ssc", http.StatusFound)
    }
}