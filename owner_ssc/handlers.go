package owner_ssc

import (
    // "fmt"
    "net/http"
    "html/template"

    "go_authentication/options"
    "go_authentication/authtoken"
)


func OwrAllSsc(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {

        cls,err := authtoken.OnToken(w,r)
        if cls == nil {
            return
        }
        if err != nil {
            return
        }

        rows,err := qsOwSsc(w,cls.User_id)
        if err != nil {
            return
        }
        names,err := owSsc(w,rows)
        if err != nil {
            return
        }

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/owner_ssc/all.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", names)
    }
}


func DtlOwrSsc(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {

        id,err := options.IdUrl(w,r)
        if err != nil {
            return
        }

        cls,err := authtoken.ListToken(w,r)
        if err != nil {
            return
        }
        
        owner := cls.User_id
        i, err := ownerIdSsc(w,id,owner)
        if err != nil {
            return
        }

        type ListData struct {
            Auth string
            I Subscription
        }
        data := ListData {
            Auth: cls.Email,
            I: i,
        }

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/owner_ssc/detail.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", data)
    }
}