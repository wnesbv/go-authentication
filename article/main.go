package article

import (
    "database/sql"
    "time"
    "fmt"
    "os"

    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
)


var db *sql.DB

func connart() {
    if err := godotenv.Load(); err != nil {
        switch {
            case true:
            fmt.Println("err: no .env file found article", err)
            break
        }
    }

    connstr := os.Getenv("DATABASE_URL")

    var err error
    db, err = sql.Open("postgres", connstr)
    if err != nil {
        fmt.Println("err: sql.Open() article", err)
    }

    err = db.Ping()
    if err != nil {
        fmt.Println("err: db.Ping() article", err)
    }
    if err == nil {
        fmt.Println("init OK..! is connected article")
    }

}


func init() {

    start := time.Now()

    go connart()
    time.Sleep(1 * time.Second)

    elapsed := time.Since(start)
    fmt.Printf(" sql article time.. :  %s \n", elapsed)
}