package profile

import (
    "database/sql"
    "os"
    "fmt"
    "net/http"
    "net/smtp"

    "golang.org/x/crypto/bcrypt"
    
    "go_authentication/authtoken"
)


func hashPassword(password string) (string, error) {

    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {

    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}


func IdUser(w http.ResponseWriter, id int) (i AllUser, err error) {
    
    row := db.QueryRow("SELECT user_id, username, email FROM users WHERE user_id=$1", id)

    err = row.Scan(
        &i.User_id,
        &i.Username,
        &i.Email,
    )

    if err == sql.ErrNoRows {
        fmt.Fprintf(w, "sql.ErrNoRows err..! : %+v\n", err)
        return
    } else if err != nil {
        fmt.Fprintf(w, "sql err..! : %+v\n", err)
        return
    }
    return i,err
}


func profilUser(w http.ResponseWriter, r *http.Request, claims *authtoken.Claims) (i *AllUser, err error) {

    i = &AllUser{}
    i = new(AllUser)
    user_id := claims.User_id
    
    row := db.QueryRow("SELECT user_id, username, email, password, created_at, updated_at FROM users WHERE user_id=$1", user_id)

    err = row.Scan(
        &i.User_id,
        &i.Username,
        &i.Email,
        &i.Password,
        &i.Created_at,
        &i.Updated_at,
    )

    if err == sql.ErrNoRows {
        fmt.Fprintf(w, "sql.ErrNoRows err..! : %+v\n", err)
        return
    } else if err != nil {
        fmt.Fprintf(w, "sql err..! : %+v\n", err)
        http.Error(w, http.StatusText(500), 500)
        return
    }

    return i,err
}


func emailSend(r *http.Request, token, email string) (err error) {

  from := os.Getenv("SMTP_EMAIL")
  password := os.Getenv("SMTP_PASS")

  to := []string{ email }

  smtpHost := "smtp.gmail.com"
  smtpPort := "587"

  message := []byte("email message\r\n" + r.Host + "/verification?veri=" + token)

  auth := smtp.PlainAuth("", from, password, smtpHost)

  err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
  if err != nil {
    fmt.Println("err SendMail..!", err)
    return
  }

  fmt.Println("Email Sent Successfully!")

  return nil
}