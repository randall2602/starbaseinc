package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

// Receive gets data from JS websocket and prints to std out
func Receive(ws *websocket.Conn) {
	var err error
	for {
		var reply string
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}
		fmt.Println("got data")
		fmt.Println(reply)
		Send(ws)
	}
}

// Send send data to JS websocket
func Send(ws *websocket.Conn) {
	if err := websocket.Message.Send(ws, "this is server data"); err != nil {
		fmt.Println("Can't send")
		return
	}
	fmt.Println("sent data")
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/ws", websocket.Handler(Receive))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
