//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func longestCommonPrefix(strs []string) string {
//	if len(strs) == 0 {
//		return ""
//	}
//	prefix, count := strs[0], len(strs)
//	for i := 1; i < count; i++ {
//		prefix = lcp(prefix, strs[i])
//		if prefix == "" {
//			break
//		}
//	}
//	return prefix
//}
//
//func lcp(str1, str2 string) string {
//	length := min(len(str1), len(str2))
//	index := 0
//	for i := 0; i < length; i++ {
//		if str1[i] == str2[i] {
//			index++
//		}
//	}
//	return str1[:index]
//}
//
//func min(x, y int) int {
//	if x < y {
//		return x
//	}
//	return y
//}
//
//func test(ch chan bool) {
//	//go func() {
//	//	//time.Sleep(1 * time.Second)
//	//	<-ch
//	//	fmt.Print("hello")
//	//}()
//	select {
//	case <-ch:
//		fmt.Println("hahaha...")
//	}
//}
//
//func main() {
//	ch := make(chan bool)
//	test(ch)
//	ch <- true
//	time.Sleep(2 * time.Second)
//	//lcp := longestCommonPrefix([]string{"cir", "car"})
//	//log.Printf("lcp: %s", lcp)
//}

package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	ID   string
	Conn *websocket.Conn
	Send chan []byte
}

type Message struct {
	SenderID   string `json:"sender_id"`
	ReceiverID string `json:"receiver_id"`
	Content    string `json:"content"`
}

var clients = make(map[string]*Client)
var register = make(chan *Client)
var unregister = make(chan *Client)
var broadcast = make(chan Message)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	clientID := r.URL.Query().Get("id")
	client := &Client{ID: clientID, Conn: conn, Send: make(chan []byte)}

	register <- client

	go handleMessages(client)

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			unregister <- client
			break
		}
		broadcast <- msg
	}
}

func handleMessages(client *Client) {
	for {
		msg := <-client.Send
		client.Conn.WriteMessage(websocket.TextMessage, msg)
	}
}

func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	http.HandleFunc("/ws", handleConnections)

	go manageClients()

	log.Println("http server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func manageClients() {
	for {
		select {
		case client := <-register:
			clients[client.ID] = client
			fmt.Println("Client registered:", client.ID)
		case client := <-unregister:
			delete(clients, client.ID)
			fmt.Println("Client unregistered:", client.ID)
		case msg := <-broadcast:
			if receiver, ok := clients[msg.ReceiverID]; ok {
				receiver.Send <- []byte(msg.Content)
			}
		}
	}
}
