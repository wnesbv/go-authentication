package owner_ssc

import (
    "time"
    "fmt"
    "net/http"
    "html/template"

    "go_authentication/options"
    "go_authentication/authtoken"
)


func AddSscUs(w http.ResponseWriter, r *http.Request) {

    cls,err := authtoken.OnToken(w,r)
    if cls == nil {
        return
    }
    if err != nil {
        return
    }

    id,err := options.IdUrl(w,r)
    if err != nil {
        return
    }

    if r.Method == "GET" {

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/owner_ssc/creativity.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", nil)
    }


    if r.Method == "POST" {
        
        user := CreatSubscription{
            Title: r.FormValue("title"),
            Description: r.FormValue("description"),
        }
        sqlStatement := `INSERT INTO subscription (title, description, owner, to_user, created_at) VALUES ($1,$2,$3,$4,$5)`

        _, err := db.Exec(sqlStatement, user.Title, user.Description, cls.User_id, id, time.Now())

        if err != nil {
            fmt.Fprintf(w, "err db.Exec()..! : %+v\n", err)
            return
        }
        http.Redirect(w, r, "/all-ssc", http.StatusFound)
    }
}

func AddSscGr(w http.ResponseWriter, r *http.Request) {

    cls,err := authtoken.OnToken(w,r)
    if cls == nil {
        return
    }
    if err != nil {
        return
    }

    id,err := options.IdUrl(w,r)
    if err != nil {
        return
    }

    if r.Method == "GET" {

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/owner_ssc/creativity.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", nil)
    }


    if r.Method == "POST" {
        
        user := CreatSubscription{
            Title: r.FormValue("title"),
            Description: r.FormValue("description"),
        }
        sqlStatement := `INSERT INTO subscription (title, description, owner, to_group, created_at) VALUES ($1,$2,$3,$4,$5)`

        _, err := db.Exec(sqlStatement, user.Title, user.Description, cls.User_id, id, time.Now())

        if err != nil {
            fmt.Fprintf(w, "err db.Exec()..! : %+v\n", err)
            return
        }
        http.Redirect(w, r, "/all-ssc", http.StatusFound)
    }
}


func OwrUpSsc(w http.ResponseWriter, r *http.Request) {

    cls,err := authtoken.SqlToken(w,r)
    if cls == nil {
        return
    }
    if err != nil {
        return
    }

    id,err := options.IdUrl(w,r)
    if err != nil {
        return
    }

    owner := cls.User_id
    i,err := ownerIdSsc(w,id,owner)
    if err != nil {
        return
    }

    if r.Method == "GET" {

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/owner_ssc/update.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", i)
    }


    if r.Method == "POST" {

        art := Subscription{
            Title: r.FormValue("title"),
            Description: r.FormValue("description"),
        }

        sqlStatement := "UPDATE subscription SET title=$3, description=$4, updated_at=$5 WHERE id=$1 AND owner=$2;"

        _, err := db.Exec(sqlStatement, id, cls.User_id, art.Title, art.Description, time.Now())
        
        if err != nil {
            fmt.Fprintf(w, "err db.Exec()..! : %+v\n", err)
            return
        }
        http.Redirect(w, r, "/all-ssc", http.StatusFound)
    }
}

func OwrDelSsc(w http.ResponseWriter, r *http.Request) {

    id, iderr := options.IdUrl(w,r)
    if iderr != nil {
        return
    }

    cls, tkerr := authtoken.OnToken(w,r)
    if cls == nil {
        return
    }
    if tkerr != nil {
        return
    }

    if r.Method == "GET" {

        data := struct {
            Items string
        }{
            Items: cls.Email,
        }

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/owner_ssc/delete.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", data)
    }


    if r.Method == "POST" {

        sqlStatement := `DELETE FROM subscription WHERE id=$1 AND owner=$2;`
        
        _, err := db.Exec(sqlStatement, id, cls.User_id)
        
        if err != nil {
            fmt.Fprintf(w, "err db.Exec()..! : %+v\n", err)
            return
        }
        
        http.Redirect(w, r, "/all-ssc", http.StatusFound)
    }
}