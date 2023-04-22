package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var connections = make(map[*websocket.Conn]bool)

var upgrader = websocket.Upgrader{
	WriteBufferSize: 4096,
	ReadBufferSize:  4096,
	CheckOrigin: func(req *http.Request) bool {
		return true
	},
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"greeting": "WELCOME TO BASIC WS CHAT"})
	})

	router.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{})
			log.Fatal(err)
		}
		defer conn.Close()

		log.Printf("New user connected: %v\n", conn.RemoteAddr().String())

		connections[conn] = true
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					connections[conn] = false
					log.Printf("User %s disconnected\n", conn.RemoteAddr().String())
				}
				break
			}

			log.Printf("User message: %v\n", string(message))

			for activeCon := range connections {
				if connections[activeCon] {
					activeCon.WriteMessage(websocket.TextMessage, message)
				}
			}
		}
	})

	log.Println("Server running on port 8000")

	log.Fatal(router.Run(":8000"))
}
