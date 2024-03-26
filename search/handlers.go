package search

import (
    "net/http"
    "html/template"
)


 func SearchHandler(w http.ResponseWriter, r *http.Request) {

    rows,err := qSearchArt(w,r)
    if err != nil {
        return
    }
    list,err := searcArt(w,rows)
    if err != nil {
        return
    }

    if r.Method == "GET" {
        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/art/search.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", list)
    }

     // fmt.Fprintln(w, "search result..")
     // for _, result := range list {
     //     fmt.Fprintln(w, "search..! : %+v\n", result)
     // }
 }
