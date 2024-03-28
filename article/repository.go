package article

import (
    "database/sql"
    "net/http"
    "fmt"
    "go_authentication/connect"
)


func qArt(w http.ResponseWriter) (rows *sql.Rows, err error) {

    conn := connect.Conn()
    rows,err = conn.Query("SELECT id, title, description, img, owner, completed, created_at, updated_at FROM article WHERE Completed=$1", true)

    if err != nil {
        switch {
            case true:
            fmt.Fprintf(w, "Error: Query..! : %+v\n", err)
            break
        }
        return
    }
    defer conn.Close()
    return rows,err
}


func qsUserArt(w http.ResponseWriter, owner int) (rows *sql.Rows, err error) {

    conn := connect.Conn()
    rows,err = conn.Query("SELECT * FROM article WHERE owner=$1", owner)

    if err != nil {
        switch {
            case true:
            fmt.Fprintf(w, "Error: Query()..! : %+v\n", err)
            break
        }
        return
    }
    defer conn.Close()
    return rows,err
}
