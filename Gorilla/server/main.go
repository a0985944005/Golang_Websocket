package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = &websocket.Upgrader{
	//如果有 cross domain 的需求，可加入這個，不檢查 cross domain
	// cross domain 跨來源資源共用，例如:img內的src引用其他網頁的圖片
	CheckOrigin: func(r *http.Request) bool { return true },
}

func main() {

	http.HandleFunc("/echo", websockethandle)
	log.Println("server start at :5407")
	log.Fatal(http.ListenAndServe(":5407", nil))
}

func websockethandle(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer func() {
		log.Println("disconnect !!")
		c.Close()
	}()
	for {
		mtype, msg, err := c.ReadMessage() //ReadMessage 一直取值
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("receive: %s\n", msg)
		err = c.WriteMessage(mtype, msg) //WriteMessage  回傳值
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
