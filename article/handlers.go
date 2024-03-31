package article

import (
    "fmt"
    "net/http"
    "html/template"
    "runtime"

    "go_authentication/options"
    "go_authentication/connect"
    "go_authentication/authtoken"
)


func HomeArticle(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {
        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/art/index.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", nil)
    }
}


func Allarticle(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {

        conn := connect.ConnSql()
        rows,err := qArt(w, conn)
        if err != nil {
            return
        }
        list,err := allArt(w, rows)
        if err != nil {
            return
        }
        defer conn.Close()

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/art/all.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", list)
    }

    fmt.Println(" All article goroutine..", runtime.NumGoroutine())
}


func UsAllArt(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {
        
        cls,err := authtoken.OnToken(w,r)
        if cls == nil {
            return
        }
        if err != nil {
            return
        }

        owner := cls.User_id
        conn := connect.ConnSql()
        rows,err := qsUserArt(w, conn,owner)
        if err != nil {
            return
        }
        list,err := userArt(w, rows)
        if err != nil {
            return
        }
        defer conn.Close()

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/art/author_id_article.html", "./tpl/base.html" ))
        
        tpl.ExecuteTemplate(w, "base", list)
    }
}


func DetArt(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {

        id,err := options.IdUrl(w,r)
        if err != nil {
            return
        }
        
        conn := connect.ConnSql()
        i,err := idArt(w, conn,id)
        if err != nil {
            return
        }
        defer conn.Close()

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/art/detail.html", "./tpl/base.html" ))
        
        tpl.ExecuteTemplate(w, "base", i)
    }
}