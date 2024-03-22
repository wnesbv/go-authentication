package subscription

import (
	"database/sql"
	"time"
	"fmt"
    "os"
	
    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
)


var db *sql.DB

func connsubscription() {
    if err := godotenv.Load(); err != nil {
        switch {
            case true:
            fmt.Println("err: no .env file found connsubscription", err)
            break
        }
    }

    connstr := os.Getenv("DATABASE_URL")

    var err error
    db, err = sql.Open("postgres", connstr)
    if err != nil {
        fmt.Println("err: sql.Open() subscription", err)
    }

    err = db.Ping()
    if err != nil {
        fmt.Println("err: db.Ping() subscription", err)
    }
    if err == nil {
        fmt.Println("init OK..! is connected subscription")
    }

}


func init() {

    start := time.Now()

	go connsubscription()
    time.Sleep(1 * time.Second)

    elapsed := time.Since(start)
    fmt.Printf(" sql subscription time.. :  %s \n", elapsed)
}
