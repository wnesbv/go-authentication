package sqlcsv

import (
    "database/sql"
    "time"
    "fmt"
    "os"

    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
)


var db *sql.DB

func conncsv() {
    if err := godotenv.Load(); err != nil {
        switch {
            case true:
            fmt.Println("err: no .env file found sqlcsv", err)
            break
        }
    }

    connstr := os.Getenv("DATABASE_URL")

    var err error
    db, err = sql.Open("postgres", connstr)
    if err != nil {
        fmt.Println("err: sql.Open() sqlcsv", err)
    }

    err = db.Ping()
    if err != nil {
        fmt.Println("err: db.Ping() sqlcsv", err)
    }
    if err == nil {
        fmt.Println("init OK..! is connected sqlcsv")
    }

}


func init() {

    start := time.Now()

    go conncsv()
    time.Sleep(1 * time.Second)

    elapsed := time.Since(start)
    fmt.Printf(" sql sqlcsv time.. :  %s \n", elapsed)
}