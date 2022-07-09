package main

import (
	"context"
	"github.com/gorilla/websocket"
	"log"
	"markdown/utils/config"
	"markdown/utils/redis"
	"net/http"
	"os"
)

var conns = make(map[string]*websocket.Conn)
var upgrader = websocket.Upgrader{} // use default options

/*
TODO:
1、线程不安全，存在多个全局变量【Doing】
2、需要支持多页面，目前只有 1 个页面
3、每次返回全量文本，开销比较大，需要实现一些 diff 算法
4、数据要持久存储，目前存在变量里，进程重启会丢失【DONE】
5、解决多人编辑的冲突问题
6、改一下前端 UI,有点儿丑
7、数据库配置放入 config 文件【DONE】
*/

func socketHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade our raw HTTP connection to a websocket based one
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection upgradation:", err)
		return
	}
	conns[conn.RemoteAddr().String()] = conn
	defer func() {
		delete(conns, conn.RemoteAddr().String())
		conn.Close()
	}()

	// 从 redis get 最后一次的文本信息，返给新建立的连接
	ctx := context.Background()
	lastMessage, err := redis.NewClient().Get(ctx, "lastMessage").Result()
	if err != nil {
		panic(err)
	}

	log.Println("lastMessage:", lastMessage)

	err = conn.WriteMessage(websocket.TextMessage, []byte(lastMessage))
	if err != nil {
		log.Println("Error during message writing:", err)
	}

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error during message reading:", err)
			break
		}
		log.Printf("Received: %s", message)

		//数据写入 redis
		err = redis.NewClient().Set(ctx, "lastMessage", message, 0).Err()
		if err != nil {
			panic(err)
		}

		//将上一次的文本信息，循环写给所有的连接
		for _, c := range conns {
			err = c.WriteMessage(messageType, message)
			if err != nil {
				log.Println("Error during message writing:", err)
				break
			}
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	body, err := os.ReadFile("statics/index.html")
	if err != nil {
		panic(err)
	}
	w.Write(body)
}

func init() {
	config.LoadConfig()
}

func main() {
	http.HandleFunc("/socket", socketHandler)
	http.HandleFunc("/", home)
	//http.HandleFunc("/:room_id", home)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
