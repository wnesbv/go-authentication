package chat

import (
	"time"
	"log"
	"fmt"
    //"os"
	"net/http"
	
	"go_authentication/options"
	"github.com/gorilla/websocket"
)


var us_upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		fmt.Println("user upgrader")
		return true
	},
}

var (
	us_clients    = make(map[*websocket.Conn]bool)
	us_register   = make(chan *websocket.Conn)
	us_unregister = make(chan *websocket.Conn)
	us_broadcast  = make(chan Message)
)


// go run
func userCh() {

	for {
		select {

		case client := <-us_register:
			us_clients[client] = true
			fmt.Println("Add client", us_clients[client])

		case client := <-us_unregister:
			delete(us_clients, client)
			client.Close()
			fmt.Println("Del client")

		case message := <-us_broadcast:
			fmt.Println("msg..", message)

			for client := range us_clients {

				if err := client.WriteJSON(message); err != nil {
					
					fmt.Println("message err:", err)
					delete(us_clients, client)
					client.Close()
				}
			}
		}
	}
}


func UsMsg(w http.ResponseWriter, r *http.Request) {

    cls,err := options.WhoisWho(w,r)
    if err != nil {
        return
    }
    id,err := options.IdUrl(w,r)
    fmt.Println("id..", id)
    if err != nil {
        return
    }

	conn,err := us_upgrade.Upgrade(w,r, nil)
	fmt.Println("Upgrade..", cls.Email)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		fmt.Println("Close..", cls.Email)
		us_unregister <- conn
		conn.Close()
	}()

	us_register <- conn

	sqlStatement := `INSERT INTO msguser (coming, owner, to_user, completed, created_at) VALUES ($1,$2,$3,$4,$5)`

	for {
		var message Message
		err := conn.ReadJSON(&message)
		if err != nil {
			fmt.Println("err message", err)
			break
		}
		fmt.Println("message..", message.Message)

    	_,err = db.Exec(sqlStatement, message.Message, cls.User_id, id, true, time.Now())

		if err != nil {
			fmt.Println("err Exec()", err)
			break
		}

		us_broadcast <- message
	}
}
