package profile

import (
    "fmt"
    "net/http"
    "html/template"

    "go_authentication/options"
)


func Home(w http.ResponseWriter, r *http.Request) {
    tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/index.html", "./tpl/base.html" ))
    tpl.ExecuteTemplate(w, "base", nil)
}


func Alluser(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/profile/all.html", "./tpl/base.html" ))

        rows, err := db.Query("SELECT user_id,username,email FROM users")

        if err != nil {
            fmt.Fprintf(w, "Error: Query()..! : %+v\n", err)
            return
        }
        defer rows.Close()

        cls, err := options.WhoisWho(w,r)
        if err != nil {
            return
        }
        type ListData struct {
            Auth string
            I []*UserList
        }
        if cls != nil {
        var names []*UserList
        for rows.Next() {
            i := new(UserList)
            err := rows.Scan(&i.User_id, &i.Username, &i.Email)

            if err != nil {
                fmt.Fprintf(w, "Error: Scan()..! : %+v\n", err)
                return
            }
            names = append(names, i)
        }

        data := ListData {
            Auth: cls.Email,
            I: names,
        }
        tpl.ExecuteTemplate(w, "base", data)

        } else {

        var names []*UserList
        for rows.Next() {
            i := new(UserList)
            err := rows.Scan(&i.User_id, &i.Username, &i.Email)

            if err != nil {
                fmt.Fprintf(w, "Error: Scan()..! : %+v\n", err)
                return
            }
            names = append(names, i)
        }
        data := ListData {
            I: names,
        }
        tpl.ExecuteTemplate(w, "base", data)
        }

        if err = rows.Close(); err != nil {
            fmt.Fprintf(w, "Error: sql..! : %+v\n", err)
        }
    }
}


/*func Alluser(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {

        rows, err := db.Query("SELECT username,email FROM users")

        if err != nil {
            fmt.Fprintf(w, "Error: db.Query()..! : %+v\n", err)
            return
        }
        defer rows.Close()

        var albums []*UserList
        
        for rows.Next() {
            data := new(UserList)
            err := rows.Scan(&data.Username, &data.Email)

            if err != nil {
                fmt.Fprintf(w, "Error: Scan()..! : %+v\n", err)
                return
            }
            albums = append(albums, data)
        }

        if err = rows.Close(); err != nil {
            fmt.Fprintf(w, "Error: sql..! : %+v\n", err)
        }

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/profile/all.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", albums)
    }
}*/