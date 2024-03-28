package connect

import (
    "database/sql"
    "time"
    "fmt"
    "os"

    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
)


func Conn() (*sql.DB) {

    start := time.Now()

    if err := godotenv.Load(); err != nil {
        switch {
            case true:
            fmt.Println("err: no .env file found..", err)
            break
        }
    }

    connstr := os.Getenv("DATABASE_URL")
    var err error
    var conn *sql.DB

    conn,err = sql.Open("postgres", connstr)
    if err != nil {
        fmt.Println("coon err: sql Open..", err)
    }

    err = conn.Ping()
    if err != nil {
        fmt.Println("coon err: Ping..", err)
    }
    if err == nil {
        fmt.Println("sql OK..! : sql is connected")
    }

    elapsed := time.Since(start)
    fmt.Printf(" sql connect time.. :  %s \n", elapsed)

    return conn
}

