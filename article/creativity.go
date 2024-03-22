package article

import (
    "time"
    "fmt"
    "os"
    "io"
    "net/http"
    "html/template"

    "go_authentication/options"
    "go_authentication/authtoken"
)


func Creativity(w http.ResponseWriter, r *http.Request) {

    cls,err := authtoken.OnToken(w,r)
    if cls == nil {
        return
    }
    if err != nil {
        return
    }


    if r.Method == "GET" {

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/art/creativity.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", nil)
    }


    if r.Method == "POST" {
        
        user := CreatArticle{
            Title: r.FormValue("title"),
            Description: r.FormValue("description"),
        }
        sqlStatement := `INSERT INTO article (title, description, owner, created_at) VALUES ($1,$2,$3,$4)`

        _, err := db.Exec(sqlStatement, user.Title, user.Description, cls.User_id, time.Now())

        if err != nil {
            fmt.Fprintf(w, "err db.Exec()..! : %+v\n", err)
            return
        }
        http.Redirect(w, r, "/author-id-article", http.StatusFound)
    }
}


func UpArt(w http.ResponseWriter, r *http.Request) {

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

    i, err := authorArt(w,r,cls,id)
    if err != nil {
        return
    }

    flag,err := options.ParseBool(r.FormValue("completed"))
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "err ParseBool()..  : %+v\n", err)
        return
    }

    if r.Method == "GET" {

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/art/update.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", i)
    }


    if r.Method == "POST" {

        art := UpdateArticle{
            Title: r.FormValue("title"),
            Description: r.FormValue("description"),
        }

        sqlStatement := `UPDATE article SET title=$3, description=$4, completed=$5, updated_at=$6 WHERE id=$1 AND owner=$2;`
        
        _, err := db.Exec(sqlStatement, id, cls.User_id, art.Title, art.Description, flag, time.Now())
        
        if err != nil {
            fmt.Fprintf(w, "err db.Exec()..! : %+v\n", err)
            return
        }
        http.Redirect(w, r, "/author-id-article", http.StatusFound)
    }
}


func DeleteArt(w http.ResponseWriter, r *http.Request) {

    id,err := options.IdUrl(w,r)
    if err != nil {
        return
    }

    cls,err := authtoken.OnToken(w,r)
    if cls == nil {
        return
    }
    if err != nil {
        return
    }

    if r.Method == "GET" {

        data := struct {
            Items string
        }{
            Items: cls.Email,
        }

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/art/delete.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", data)
    }


    if r.Method == "POST" {

        sid := r.URL.Query().Get("id")
        fpath := "./static/img/art/" + cls.Email + "/" + sid

        p := options.DelFolder(fpath)
        if p != nil {
            fmt.Fprintf(w, "del path img..! : %+v\n", p)
            return
        }

        sqlStatement := `DELETE FROM article WHERE id=$1 AND owner=$2;`
        
        _, err := db.Exec(sqlStatement, id, cls.User_id)
        
        if err != nil {
            fmt.Fprintf(w, "err db.Exec()..! : %+v\n", err)
            return
        }
        
        http.Redirect(w, r, "/author-id-article", http.StatusFound)
    }
}


func ImgArt(w http.ResponseWriter, r *http.Request) {

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

    i,err := authorArt(w,r,cls,id)
    if err != nil {
        return
    }


    if r.Method == "GET" {

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/art/img.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", i)
    }


    if r.Method == "POST" {

        if i.Img != nil {
            err := os.Remove("." + *i.Img) 
            if err != nil { 
                fmt.Println("err os.Remove().. ", err)
            }
        }

        file, handler, err := r.FormFile("file")
        if err != nil{
            fmt.Println("Error Data retrieving", err)
            return
        }
        defer file.Close()

        flname := handler.Filename
        sid := r.URL.Query().Get("id")

        fpath := "./static/img/art/" + cls.Email + "/" + sid + "/"
        fname := "./static/img/art/" + cls.Email + "/" + sid + "/"  + flname
        fle := "/static/img/art/" + cls.Email + "/" + sid + "/" + flname

        fmt.Printf("fle %+v\n: ", fle)
        fmt.Printf("Uploaded File : %+v\n", flname)
        fmt.Printf("File Size : %+v\n" , handler.Size)
        fmt.Printf("MIME Header : %+v\n" , handler.Header)


        mkdirerr := os.MkdirAll(fpath, 0750)
        if mkdirerr != nil {
            fmt.Println("Error os.MkdirAll():", mkdirerr)
        }
        img, err := os.Create(fname)
        if err != nil {
            fmt.Println("Error os.Create():", err)
        }
        defer img.Close()

        if _, err := io.Copy(img, file); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        sqlStatement := `UPDATE article SET img=$3, updated_at=$4 WHERE id=$1 AND owner=$2;`
        
        _, err = db.Exec(sqlStatement,id,cls.User_id,fle,time.Now())
        
        if err != nil {
            fmt.Fprintf(w, "err db.Exec()..! : %+v\n", err)
            return
        }

        http.Redirect(w, r, fname, http.StatusFound)
    }
}


func DelImgArt(w http.ResponseWriter, r *http.Request) {

    id,err := options.IdUrl(w,r)
    if err != nil {
        return
    }

    cls,err := authtoken.OnToken(w,r)
    if cls == nil {
        return
    }
    if err != nil {
        return
    }

    i,err := authorArt(w,r,cls,id)
    if err != nil {
        return
    }


    if r.Method == "GET" {

        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/art/imgdel.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", i)
    }


    if r.Method == "POST" {

        e := os.Remove("." + *i.Img) 
        if e != nil { 
            fmt.Println("e.. ", e)
        } 

        sqlStatement := `UPDATE article SET img=$3, updated_at=$4 WHERE id=$1 AND owner=$2;`
        
        _, err = db.Exec(sqlStatement,id,cls.User_id,nil,time.Now())
        
        if err != nil {
            fmt.Fprintf(w, "err db.Exec()..! : %+v\n", err)
            return
        }

        http.Redirect(w, r, "/author-id-article", http.StatusFound)
    }
}