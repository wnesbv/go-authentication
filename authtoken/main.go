package authtoken

import (
    "time"
    "fmt"
    "os"
    "errors"
    "net/http"

    "github.com/golang-jwt/jwt/v5"
)


func RedirectToken(w http.ResponseWriter, r *http.Request) (claims *Claims, tkerr error) {

    c, err := r.Cookie("Visitor")
    if err != nil {
        switch {
        case errors.Is(err, http.ErrNoCookie):
            http.Redirect(w, r, "/login", http.StatusFound)
        }
        return
    }

    claims = &Claims{}

    token, err := jwt.ParseWithClaims(c.Value, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("JWT_SECRET")), nil
    })

    if err != nil {
        http.Redirect(w, r, "/login", http.StatusFound)
        return
    }

    if !token.Valid {
        http.Redirect(w, r, "/login", http.StatusFound)
        return
    }

    return claims, err
}


func OnToken(w http.ResponseWriter, r *http.Request) (claims *Claims, tkerr error) {

    c, err := r.Cookie("Visitor")
    if err != nil {
        switch {
        case errors.Is(err, http.ErrNoCookie):
            http.Redirect(w, r, "/login", http.StatusUnauthorized)
        }
        return
    }

    tkstr := c.Value
    cls := &Claims{}

    token, err := jwt.ParseWithClaims(tkstr, cls, func(token *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("JWT_SECRET")), nil
    })

    if token == nil {
        fmt.Fprintf(w, "token nil..! : %+v\n", err)
        return
    }
    if err != nil {
        fmt.Fprintf(w, "OnToken err: jwt.ParseWithClaims()..! : %+v\n", err)
        return
    }

    return cls, err
}


func SqlToken(w http.ResponseWriter, r *http.Request) (claims *Claims, err error) {

    c, err := r.Cookie("Visitor")
    if err != nil {
        switch {
        case errors.Is(err, http.ErrNoCookie):
            fmt.Println("err r.Cookie().. ", err)
            http.Redirect(w, r, "/login", http.StatusUnauthorized)
        }
        return
    }

    claims = &Claims{}

    token, err := jwt.ParseWithClaims(c.Value, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("JWT_SECRET")), nil
    })


    if err != nil {
        if err == jwt.ErrSignatureInvalid {
            fmt.Fprintf(w, "err jwt.ErrSignatureInvalid..! : %+v\n", err)
            return
        }
        fmt.Fprintf(w, "SqlToken err: jwt.ParseWithClaims()..! : %+v\n", err)
        return
    }
    if !token.Valid {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    return claims, err
}


func ListToken(w http.ResponseWriter, r *http.Request) (*Claims, error) {

    c,err := r.Cookie("Visitor")

    cls := &Claims{}

    switch {
    case errors.Is(err, http.ErrNoCookie):
        break
    default:
        _,err = jwt.ParseWithClaims(c.Value, cls, func(token *jwt.Token) (interface{}, error) {
            return []byte(os.Getenv("JWT_SECRET")), nil
        })
    }
    return cls,err
}


func WhoisWho(w http.ResponseWriter, r *http.Request) *Claims {

    c,err := r.Cookie("Visitor")

    cls := &Claims{}

    switch {
    case errors.Is(err, http.ErrNoCookie):
        break
    default:
        _,err = jwt.ParseWithClaims(c.Value, cls, func(token *jwt.Token) (interface{}, error) {
            return []byte(os.Getenv("JWT_SECRET")), nil
        })
    }
    return cls
}



func BuildSend(w http.ResponseWriter, email string) (tkstr string, tkerr error) {

    token := jwt.New(jwt.SigningMethodHS256)
    cls := token.Claims.(jwt.MapClaims)

    cls["email"] = email
    cls["exp"] = time.Now().Add(time.Minute * 60).Unix()

    tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "err SignedString()..! : %+v\n", err)
        return
    }
    return tokenString, err
}


func VerifySendToken(w http.ResponseWriter, veri string) (claims *Claims, tkerr error) {

    cls := &Claims{}

    token, err := jwt.ParseWithClaims(veri, cls, func(token *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("JWT_SECRET")), nil
    })

    if token == nil {
        fmt.Fprintf(w, "token nil..! : %+v\n", err)
        return
    }
    if err != nil {
        fmt.Fprintf(w, "Verify err: jwt.ParseWithClaims()..! : %+v\n", err)
        return
    }

    return cls, err
}