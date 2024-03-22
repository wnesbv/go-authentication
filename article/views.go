package article

import (
    "database/sql"
    "fmt"
    "os"
    "errors"
    "net/http"
    // "path/filepath"
    
    "github.com/golang-jwt/jwt/v5"

    "go_authentication/authtoken"
)


func allArt(w http.ResponseWriter, rows *sql.Rows) (names []*Article, err error) {

    defer rows.Close()
    for rows.Next() {
        i := new(Article)
        err = rows.Scan(
            &i.Id,
            &i.Title,
            &i.Description,
            &i.Img,
            &i.Owner,
            &i.Completed,
            &i.Created_at,
            &i.Updated_at,
        )
        if err != nil {
            fmt.Fprintf(w, "Error Scan()..! : %+v\n", err)
            return
        }
        names = append(names, i)
    }
    return names, err
}


func userArt(w http.ResponseWriter, rows *sql.Rows) (names []*Article, err error) {

    defer rows.Close()
    for rows.Next() {
        i := new(Article)
        err = rows.Scan(
            &i.Id,
            &i.Title,
            &i.Description,
            &i.Img,
            &i.Owner,
            &i.Completed,
            &i.Created_at,
            &i.Updated_at,
        )
        if err != nil {
            fmt.Fprintf(w, "Error Scan()..! : %+v\n", err)
            return
        }
        names = append(names, i)
    }

    // if qerr = rows.Close(); qerr != nil {
    //     fmt.Fprintf(w, "Error: sql..! : %+v\n", qerr)
    // }
    
    // if closeErr := rows.Close(); closeErr != nil {
    //     http.Error(w, closeErr.Error(), http.StatusInternalServerError)
    //     return
    // }
    // if err != nil {
    //     http.Error(w, err.Error(), http.StatusInternalServerError)
    //     return
    // }
    // if err = rows.Err(); err != nil {
    //     http.Error(w, err.Error(), http.StatusInternalServerError)
    //     return
    // }

    return names, err
}


func authorArt(w http.ResponseWriter, r *http.Request, claims *authtoken.Claims, id int) (i *Article, err error) {

    i = &Article{}
    i = new(Article)
    owner := claims.User_id
    
    row := db.QueryRow("SELECT * FROM article WHERE id=$1 AND owner=$2", id,owner)

    err = row.Scan(
        &i.Id,
        &i.Title,
        &i.Description,
        &i.Img,
        &i.Owner,
        &i.Completed,
        &i.Created_at,
        &i.Updated_at,
    )

    if err == sql.ErrNoRows {
        fmt.Fprintf(w, "err sql.ErrNoRows..! : %+v\n", err)
        return
    } else if err != nil {
        fmt.Fprintf(w, "err sql..! : %+v\n", err)
        return
    }

    return i, err
}


func idArt(w http.ResponseWriter, id int) (i Article, err error) {
    
    row := db.QueryRow("SELECT * FROM article WHERE id=$1", id)

    err = row.Scan(
        &i.Id,
        &i.Title,
        &i.Description,
        &i.Img,
        &i.Owner,
        &i.Completed,
        &i.Created_at,
        &i.Updated_at,
    )

    if err == sql.ErrNoRows {
        fmt.Fprintf(w, "err sql.ErrNoRows..! : %+v\n", err)
        return
    } else if err != nil {
        fmt.Fprintf(w, "err sql..! : %+v\n", err)
        return
    }

    return i, err
}


func OnAuth(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {
        
        c, err := r.Cookie("Visitor")
        if err != nil {
            switch {
            case errors.Is(err, http.ErrNoCookie):
                http.Redirect(w, r, "/login", http.StatusUnauthorized)
            }
            return
        }

        tkstr := c.Value
        claims := &authtoken.Claims{}

        token, err := jwt.ParseWithClaims(tkstr, claims, func(token *jwt.Token) (interface{}, error) {
            return []byte(os.Getenv("JWT_SECRET")), nil
        })


        if err != nil {
            if err == jwt.ErrSignatureInvalid {
                w.WriteHeader(http.StatusUnauthorized)
                fmt.Fprintf(w, "err jwt.ErrSignatureInvalid..! : %+v\n", err)
                return
            }
            w.WriteHeader(http.StatusBadRequest)
            fmt.Fprintf(w, "err jwt.ParseWithClaims()..! : %+v\n", err)
            return
        }
        if !token.Valid {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }

    }
    
}