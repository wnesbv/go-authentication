package profile

import (
    "database/sql"
    // "context"
    "time"
    "fmt"
    "os"
    "net/http"
    "html/template"
    
    "github.com/joho/godotenv"
    "github.com/golang-jwt/jwt/v5"
)


func Login(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {
        
        tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/auth/login.html", "./tpl/base.html" ))

        tpl.ExecuteTemplate(w, "base", nil)
    }

    if r.Method == "POST" {

        email := r.FormValue("email")
        password := r.FormValue("password")

        var storedPassword string

        err := db.QueryRow(`SELECT password FROM users WHERE email=$1`, email).Scan(&storedPassword)

        switch {
        case err == sql.ErrNoRows:
            fmt.Fprintf(w, "Error: login failed email..! : %+v\n", email)
            fmt.Fprintf(w, "Error: login failed email err..! : %+v\n", err)
            return
        case err != nil:
            fmt.Fprintf(w, "Error: QueryRow..! : %+v\n", err)
            break
        default:
            fmt.Println("email is : ", email)
        }


        match := checkPasswordHash(password, storedPassword)
        if match == false {
            fmt.Fprintf(w, "Match matching passwords..! : %+v\n", match)
            return 
        }

        if match {

            if err := godotenv.Load(); err != nil {
                fmt.Fprintf(w, "No .env file found : %+v\n", err)
                return
            }

            // var ctx context.Context
            var user_id int

            err := db.QueryRow(`SELECT user_id FROM users WHERE email=$1`, email).Scan(&user_id)
            // err := db.QueryRowContext(ctx, "SELECT user_id FROM users WHERE email=$1", email).Scan(&user_id)

            switch {
            case err == sql.ErrNoRows:
                fmt.Fprintf(w, "no user with user_id : %+v\n", user_id)
                fmt.Fprintf(w, "no user with user_id err : %+v\n", err)
                return
            case err != nil:
                break
            default:
                fmt.Println("user_id is : ", user_id)
            }

            token := jwt.New(jwt.SigningMethodHS256)
            cls := token.Claims.(jwt.MapClaims)

            cls["authorized"] = true
            cls["user_id"] = user_id
            cls["email"] = email
            cls["exp"] = time.Now().Add(time.Minute * 60).Unix()

            tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

            if err != nil {
                w.WriteHeader(http.StatusBadRequest)
                fmt.Fprintf(w, "err token.SignedString()..! : %+v\n", err)
                return
            }

            cookie := http.Cookie{
                Name:     "Visitor",
                Value:    tokenString,
                Path:     "/",
                MaxAge:   3600,
                HttpOnly: true,
                Secure:   false,
                SameSite: http.SameSiteLaxMode,
            }
            http.SetCookie(w, &cookie)

            fmt.Fprintf(w, "OK : token..!")
            fmt.Fprintf(w, " OK : login successful..!")
            
            return
        }
        
        fmt.Fprintf(w, "Error: login failed..!")
        return
    }

}


func AuthToken(w http.ResponseWriter, r *http.Request) {

    tpl := template.Must(template.ParseFiles("./tpl/navbar.html", "./tpl/auth/auth.html", "./tpl/base.html" ))

    c, err := r.Cookie("Visitor")
    if err != nil {
        if err == http.ErrNoCookie {
            w.WriteHeader(http.StatusUnauthorized)
            fmt.Fprintf(w, "err http.ErrNoCookie..! : %+v\n", err)
            return
        }
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "err r.Cookie()..! : %+v\n", err)
        return
    }

    tkstr := c.Value
    cls := &Claims{}

    token, err := jwt.ParseWithClaims(tkstr, cls, func(token *jwt.Token) (any, error) {
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

    tpl.ExecuteTemplate(w, "base", cls)
}