package subscription

import (
    // "fmt"
    "net/http"
    "html/template"

    // "go_authentication/options"
    "go_authentication/authtoken"
)


func AllSsc(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {

        rows,err := qsAllSsc(w)
        if err != nil {
            return
        }
        names,err := allSsc(w,rows)
        if err != nil {
            return
        }
        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/ssc/all.html", "./tpl/base.html" ))
        tpl.ExecuteTemplate(w, "base", names)
    }
}


func ToUsAllSsc(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {

        cls,err := authtoken.OnToken(w,r)
        if cls == nil {
            return
        }
        if err != nil {
            return
        }

        to_user := cls.User_id

        rows,err := qsUserAllSsc(w,to_user)
        if err != nil {
            return
        }
        names,err := userSsc(w,rows)
        if err != nil {
            return
        }
        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/ssc/user.html", "./tpl/base.html" ))
        tpl.ExecuteTemplate(w, "base", names)
    }
}

func ToGroupAllSsc(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {

        cls,err := authtoken.OnToken(w,r)
        if cls == nil {
            return
        }
        if err != nil {
            return
        }

        owner := cls.User_id

        rows,err := qsGroupAllSsc(w,owner)
        if err != nil {
            return
        }
        names,err := roomSsc(w,rows)
        if err != nil {
            return
        }
        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/ssc/group.html", "./tpl/base.html" ))
        tpl.ExecuteTemplate(w, "base", names)
    }
}