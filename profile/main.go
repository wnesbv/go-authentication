package profile

import (
    "database/sql"
    "time"
    "fmt"
    "os"
    "runtime"

    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
)


var db *sql.DB

func connprof() {
    if err := godotenv.Load(); err != nil {
        switch {
            case true:
            fmt.Println("err: no .env file found profile", err)
            break
        }
    }

    connstr := os.Getenv("DATABASE_URL")

    var err error
    db,err = sql.Open("postgres", connstr)
    if err != nil {
        fmt.Println("err: sql.Open() profile", err)
    }

    err = db.Ping()
    if err != nil {
        fmt.Println("err: db.Ping() profile", err)
    }
    if err == nil {
        fmt.Println("init OK..! is connected profile")
    }

}


func init() {

    start := time.Now()

    go connprof()
    time.Sleep(1 * time.Second)

    elapsed := time.Since(start)
    fmt.Printf(" sql profile time.. :  %s \n", elapsed)

    fmt.Println(" profile goroutine..", runtime.NumGoroutine())
}