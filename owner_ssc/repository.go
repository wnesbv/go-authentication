package owner_ssc

import (
    "database/sql"
    "net/http"
    "fmt"
)


func qsOwSsc(w http.ResponseWriter, owner int) (rows *sql.Rows, err error) {

    rows,err = db.Query("SELECT * FROM subscription WHERE owner=$1", owner)

    if err != nil {
        switch {
            case true:
            fmt.Fprintf(w, "Error: Query()..! : %+v\n", err)
            break
        }
        return
    }
    return rows,err
}


