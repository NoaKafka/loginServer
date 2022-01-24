package chat

import (
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

// wss://pubwss.bithumb.com/pub/ws
//var addr = flag.String("addr", "localhost:8080", "http service address")

func sendMessege(conn *websocket.Conn, msg string) {
	err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		log.Fatal(err)
	}
}

func Chat(id string) error {
	flag.Parse()
	log.SetFlags(0)
	log.Println(id)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "ws"}
	log.Printf("connecting to %s", u.String())
	uu := u.String() + "?id=" + id
	log.Println(uu)
	c, _, err := websocket.DefaultDialer.Dial(uu, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	// 메세지를 읽는 비동기 함수
	go func() {
		defer close(done)

		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)

		}
	}()

	// go func() {
	// 	buffer := bufio.NewReader(os.Stdin)
	// 	for {
	// 		msg, err := buffer.ReadString('\n')
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		} else {
	// 			sendMessege(c, msg)
	// 		}
	// 	}
	// }()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		log.Println("Start")

		select {

		case <-done:
			return nil

		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return err
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return nil
		}
	}
}
