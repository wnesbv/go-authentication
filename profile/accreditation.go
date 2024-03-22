package profile

import (
    "time"
    "fmt"
    "net/http"
    "html/template"
    
    "go_authentication/authtoken"
)


func Signup(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/signup.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", nil)
    }


    if r.Method == "POST" {

        user := User{
            Username: r.FormValue("username"),
            Email:    r.FormValue("email"),
            Password: r.FormValue("password"),
        }

        hash, _ := hashPassword(user.Password)
        fmt.Println("Password:", user.Password)
        fmt.Println("Hash:    ", hash)

        sqlStatement := `INSERT INTO users (username,email,password,created_at) VALUES ($1,$2,$3,$4)`
        
        _, err := db.Exec(sqlStatement, user.Username, user.Email, hash, time.Now())
        
        if err != nil {
            fmt.Fprintf(w, "err db.Exec()..! : %+v\n", err)
            return
        }

        http.Redirect(w, r, "/alluser", http.StatusFound)
    }
}


func UpName(w http.ResponseWriter, r *http.Request) {

    cls, tkerr := authtoken.SqlToken(w,r)

    if cls == nil {
        return
    }
    if tkerr != nil {
        return
    }

    i, ierr := profilUser(w,r,cls)
    if ierr != nil {
        return
    }

    if r.Method == "GET" {

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/profile/update_name.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", i)
    }


    if r.Method == "POST" {

        user := User{
            Username: r.FormValue("username"),
        }

        sqlStatement := `UPDATE users SET username=$2, updated_at=$3 WHERE user_id=$1;`
        
        _, err := db.Exec(sqlStatement, cls.User_id, user.Username, time.Now())

        if err != nil {
            fmt.Fprintf(w, "err db.Exec()..! : %+v\n", err)
            return
        }
        http.Redirect(w, r, "/alluser", http.StatusFound)
    }
}


func UpPass(w http.ResponseWriter, r *http.Request) {

    cls, tkerr := authtoken.SqlToken(w,r)

    if cls == nil {
        return
    }
    if tkerr != nil {
        return
    }

    i, ierr := profilUser(w,r,cls)
    if ierr != nil {
        return
    }

    if r.Method == "GET" {

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/profile/update_password.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", i)
    }


    if r.Method == "POST" {

        user := User{
            Password: r.FormValue("password"),
        }
        hash, _ := hashPassword(user.Password)

        sqlStatement := `UPDATE users SET password=$2, updated_at=$3 WHERE user_id=$1;`
        
        _, err := db.Exec(sqlStatement, cls.User_id, hash, time.Now())
        
        if err != nil {
            fmt.Fprintf(w, "err db.Exec()..! : %+v\n", err)
            return
        }
        http.Redirect(w, r, "/alluser", http.StatusFound)
    }
}


func EmailSend(w http.ResponseWriter, r *http.Request) {

    cls, tkerr := authtoken.SqlToken(w,r)

    if cls == nil {
        return
    }
    if tkerr != nil {
        return
    }

    i, ierr := profilUser(w,r,cls)
    if ierr != nil {
        return
    }

    if r.Method == "GET" {

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/profile/send_email.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", i)
    }


    if r.Method == "POST" {

        email := r.FormValue("email")

        token, err := authtoken.BuildSend(w,email)
        if err != nil {
            return
        }
        emailerr := emailSend(r,token,email)
        if emailerr != nil {
            return
        }

        http.Redirect(w, r, "/alluser", http.StatusFound)
    }
}


func VerifyEmail(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {

        cls, tkerr := authtoken.SqlToken(w,r)

        if cls == nil {
            return
        }
        if tkerr != nil {
            return
        }

        veri := r.URL.Query().Get("veri")
        
        veri_email, verr := authtoken.VerifySendToken(w,veri)
        if veri_email == nil {
            return
        }
        if verr != nil {
            return
        }

        sqlst := `UPDATE users SET email=$2, updated_at=$3 WHERE user_id=$1;`
        
        _, err := db.Exec(sqlst, cls.User_id, veri_email.Email, time.Now())

        if err != nil {
            fmt.Fprintf(w, "err db.Exec()..! : %+v\n", err)
            return
        }
        
        http.Redirect(w, r, "/alluser", http.StatusFound)
    }
}


func DelUs(w http.ResponseWriter, r *http.Request) {

    cls, tkerr := authtoken.SqlToken(w,r)

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

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/profile/delete.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", data)
    }


    if r.Method == "POST" {

        sqlStatement := `DELETE FROM users WHERE user_id=$1;`
        
        _, err := db.Exec(sqlStatement, cls.User_id)
        
        if err != nil {
            fmt.Fprintf(w, "err db.Exec()..! : %+v\n", err)
            return
        }
        http.Redirect(w, r, "/alluser", http.StatusFound)
    }
}