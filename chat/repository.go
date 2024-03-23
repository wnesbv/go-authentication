package chat

import (
    "database/sql"
    "net/http"
    "fmt"

    "github.com/lib/pq"
)


func qSscChUs(w http.ResponseWriter, id int) (user []int,completed bool,err error) {

    row := db.QueryRow("SELECT id,owner,to_user,completed FROM subscription WHERE id=$1 AND completed=$2", id,true)

    var i Subscription
    err = row.Scan(&i.Id,&i.Owner,&i.To_user,&i.Completed)

    if err == sql.ErrNoRows {
        fmt.Fprintf(w, "user ErrNoRows..! : %+v\n", err)
    } else if err != nil {
        fmt.Fprintf(w, "user err sql..! : %+v\n", err)
    }

    user = append(user,i.Owner,i.To_user)
    return user,i.Completed,err
}


func qSscGrChUs(w http.ResponseWriter, id int) (user []int,err error) {

    rows,err := db.Query("SELECT owner,to_group,completed FROM subscription WHERE to_group=$1 AND completed=$2", id,true)

    if err != nil {
        switch {
            case true:
            fmt.Fprintf(w, "Error: qSscGrChUs Query()..! : %+v\n", err)
            break
        }
        return
    }

    defer rows.Close()
    for rows.Next() {
        i := new(Subscription)
        err = rows.Scan(
            &i.Owner,
            &i.To_group,
            &i.Completed,
        )
        if err != nil {
            fmt.Fprintf(w, "Error qSscGrChUs Scan()..! : %+v\n", err)
            return
        }
        user = append(user, i.Owner)
    }
    fmt.Println("user..", user)
    return user, err
}


func qUsCh(w http.ResponseWriter, owner int, to_user int) (rows *sql.Rows, err error) {

    rows1,err := db.Query("SELECT id FROM msguser WHERE owner=$1 AND to_user=$2 AND completed=$3", owner,to_user,true)
    rows2,err := db.Query("SELECT id FROM msguser WHERE owner=$1 AND to_user=$2 AND completed=$3", to_user,owner,true)

    if err != nil {
        switch {
            case true:
            fmt.Fprintf(w, "Error: Query()..! : %+v\n", err)
            break
        }
        return
    }

    defer rows1.Close()
    var obj1 []int
    for rows1.Next() {
        i := new(MsgUser)
        err = rows1.Scan(
            &i.Id,
        )
        if err != nil {
            fmt.Fprintf(w, "Error qUsCh Scan()..! : %+v\n", err)
            return
        }
        obj1 = append(obj1, i.Id)
    }
    fmt.Println("obj1", obj1)

    defer rows2.Close()
    var obj2 []int
    for rows2.Next() {
        i := new(MsgUser)
        err = rows2.Scan(
            &i.Id,
        )
        if err != nil {
            fmt.Fprintf(w, "Error qUsCh 2 Scan()..! : %+v\n", err)
            return
        }
        obj2 = append(obj2, i.Id)
    }
    fmt.Println("obj2", obj2)

    var names []int
    names = append(obj1, obj2...)

    rows,err = db.Query("SELECT * FROM msguser WHERE id = ANY($1)", pq.Array(names))

    if err != nil {
        switch {
            case true:
            fmt.Fprintf(w, "Error: Query()..! : %+v\n", err)
            break
        }
        return
    }

    return rows, err
}


func qGroup(w http.ResponseWriter) (rows *sql.Rows, err error) {

    rows,err = db.Query("SELECT * FROM groups WHERE completed=$1", true)

    if err != nil {
        switch {
            case true:
            fmt.Fprintf(w, "Error: Query()..! : %+v\n", err)
            break
        }
        return
    }
    return rows, err
}


func qUsGroup(w http.ResponseWriter, owner int) (rows *sql.Rows, err error) {

    rows,err = db.Query("SELECT * FROM groups WHERE owner=$1", owner)

    if err != nil {
        switch {
            case true:
            fmt.Fprintf(w, "Error: Query()..! : %+v\n", err)
            break
        }
        return
    }
    return rows, err
}


func qGrChat(w http.ResponseWriter, to_group int) (rows *sql.Rows, err error) {

    rows,err = db.Query("SELECT * FROM msggroups WHERE to_group=$1 AND completed=$2", to_group,true)

    if err != nil {
        switch {
            case true:
            fmt.Fprintf(w, "Error: Query()..! : %+v\n", err)
            break
        }
        return
    }
    return rows, err
}


            // w.Header().Set("Content-Type", "text/html")
            // w.WriteHeader(http.StatusInternalServerError)
            // w.Write([]byte("500 - Something bad happened!"))
            // http.Error(w, "my own error message", 400)