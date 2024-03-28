package article

import (
    "time"
    "fmt"
    "net/http"
    "html/template"

    "go_authentication/options"
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

        start := time.Now()

        rows,err := qArt(w)
        if err != nil {
            return
        }
        list,err := allArt(w, rows)
        if err != nil {
            return
        }

        elapsed := time.Since(start)
        fmt.Printf(" list article time.. :  %s \n", elapsed)

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/art/all.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", list)
    }
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
        rows,err := qsUserArt(w, owner)
        if err != nil {
            return
        }
        list,err := userArt(w, rows)
        if err != nil {
            return
        }

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
        
        i,err := idArt(w, id)
        if err != nil {
            return
        }

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/art/detail.html", "./tpl/base.html" ))
        
        tpl.ExecuteTemplate(w, "base", i)
    }
}