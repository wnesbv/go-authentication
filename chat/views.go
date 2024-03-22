package chat

import (
    "database/sql"
    "fmt"
    // "os"
    "net/http"
    // "path/filepath"

    "github.com/lib/pq"
    // "go_authentication/authtoken"
)


func idUs(w http.ResponseWriter, users []int, completed bool) (names []*ChUser, err error) {

    if completed == true {
        rows,err := db.Query("SELECT user_id,username,email FROM users WHERE user_id = ANY($1)", pq.Array(users))

        if err != nil {
            switch {
                case true:
                fmt.Fprintf(w, "Error: Query()..! : %+v\n", err)
                break
            }
        }

        defer rows.Close()
        for rows.Next() {
            i := new(ChUser)
            err = rows.Scan(
                &i.User_id,
                &i.Username,
                &i.Email,
            )
            if err != nil {
                fmt.Fprintf(w, "Error idUs Scan()..! : %+v\n", err)
            }
            names = append(names, i)
        }
        return names,err
    }
    return
}


func usChat(w http.ResponseWriter, rows *sql.Rows) (names []*MsgUser, err error) {

    defer rows.Close()
    for rows.Next() {
        i := new(MsgUser)
        err = rows.Scan(
            &i.Id,
            &i.Coming,
            &i.Img,
            &i.Owner,
            &i.To_user,
            &i.Completed,
            &i.Created_at,
            &i.Updated_at,
        )
        if err != nil {
            fmt.Fprintf(w, "Error usChat Scan()..! : %+v\n", err)
            return
        }
        names = append(names, i)
    }
    return names, err
}


func allGroup(w http.ResponseWriter, rows *sql.Rows) (names []*Group, err error) {

    defer rows.Close()
    for rows.Next() {
        i := new(Group)
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


func userGroup(w http.ResponseWriter, rows *sql.Rows) (names []*Group, err error) {

    defer rows.Close()
    for rows.Next() {
        i := new(Group)
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


func idGroup(w http.ResponseWriter, id int) (i Group, err error) {
    
    row := db.QueryRow("SELECT * FROM groups WHERE id=$1", id)

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


func rmChat(w http.ResponseWriter, rows *sql.Rows, owner int, to_group int) (names []*MsgGroup, err error) {

    defer rows.Close()
    for rows.Next() {
        i := new(MsgGroup)
        err = rows.Scan(
            &i.Id,
            &i.Coming,
            &i.Img,
            &i.Owner,
            &i.To_group,
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