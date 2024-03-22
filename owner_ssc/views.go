package owner_ssc

import (
    "database/sql"
    "fmt"
    "net/http"
)


func owSsc(w http.ResponseWriter, rows *sql.Rows) (names []*Subscription, err error) {

    defer rows.Close()
    for rows.Next() {
        i := new(Subscription)
        err = rows.Scan(
            &i.Id,
            &i.Title,
            &i.Description,
            &i.Owner,
            &i.To_user,
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
    return names,err
}


func ownerIdSsc(w http.ResponseWriter, id int, owner int) (i Subscription, err error) {
    
    row := db.QueryRow("SELECT * FROM subscription WHERE id=$1 AND owner=$2", id,owner)

    err = row.Scan(
        &i.Id,
        &i.Title,
        &i.Description,
        &i.Owner,
        &i.To_user,
        &i.To_group,
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

    return i,err
}


