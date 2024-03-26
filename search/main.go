package search

import (
    "database/sql"
    "time"
    "fmt"
    "os"

    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
)


 var db *sql.DB

func connsearch() {
    if err := godotenv.Load(); err != nil {
        switch {
            case true:
            fmt.Println("err: no .env file found search", err)
            break
        }
    }

    connstr := os.Getenv("DATABASE_URL")

    var err error
    db, err = sql.Open("postgres", connstr)
    if err != nil {
        fmt.Println("err: sql.Open() search", err)
    }

    err = db.Ping()
    if err != nil {
        fmt.Println("err: db.Ping() search", err)
    }
    if err == nil {
        fmt.Println("init OK..! is connected search")
    }

}


func init() {

    start := time.Now()

    go connsearch()
    time.Sleep(1 * time.Second)

    elapsed := time.Since(start)
    fmt.Printf(" sql search time.. :  %s \n", elapsed)
}