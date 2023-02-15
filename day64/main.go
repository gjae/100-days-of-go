package main

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var connecteds = 0

type Client struct {
	Ctx         *gin.Context
	ID          int
	Message     chan string
	UserCounter chan int
}

type ClientList []*Client

var clients ClientList

type ServerSide struct {
	Clients ClientList

	Router *gin.Engine
}

func HomeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Contador de usuarios"})
}

func (client *Client) UsersCounterNotify() {
	exit := false
	for !exit {
		select {
		case msg := <-client.Message:
			if msg != "exit" {
				log.Printf("Notificando: %s", msg)
				client.Ctx.SSEvent("notification", msg)
			} else {
				client.Ctx.SSEvent("close", true)
				exit = true
			}
		case counter := <-client.UserCounter:
			log.Print("Enviando counter")
			data, _ := json.Marshal(gin.H{"counter": counter})
			client.Ctx.SSEvent("counter", string(data))
		}
		time.Sleep(time.Second * 10)

	}

}

func DashboardCounterHandler(c *gin.Context) {
	client := &Client{ID: rand.Intn(200), Ctx: c, Message: make(chan string), UserCounter: make(chan int)}

	clients = append(clients, client)

	log.Print("Nuevo usuario conectado")

	c.Stream(func(w io.Writer) bool {

		go client.UsersCounterNotify()
		client.Ctx.SSEvent("client", client.ID)
		time.Sleep(time.Second * 3)
		data, _ := json.Marshal(gin.H{"counter": connecteds})
		client.Ctx.SSEvent("counter", string(data))
		time.Sleep(time.Second * 3)
		return true
	})
}

func main() {
	server := NewServer()

	server.Router.GET("/", UserCounterMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"error": false})
	})

	server.Router.GET("/dashboard/", HomeHandler)
	server.Router.GET("/dashboard/counter", SeverSideEventMiddleware(), DashboardCounterHandler)
	server.Router.POST("/dashboard/:user/disconnect", func(c *gin.Context) {
		userId, _ := strconv.Atoi(c.Param("user"))
		aux := make(ClientList, 0)
		for _, ob := range clients {
			if ob.ID == userId {
				log.Printf("Desconectando ... %d", userId)
				c.JSON(http.StatusOK, gin.H{"error": false})
				log.Print("Ok ", clients)
			} else {
				aux = append(aux, ob)
			}
		}

		clients = aux
	})

	server.Router.Run(":8000")
}

func NewServer() *ServerSide {
	server := &ServerSide{
		Router:  gin.Default(),
		Clients: make([]*Client, 0),
	}

	server.Router.LoadHTMLGlob("templates/*")

	return server
}

func SeverSideEventMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "Keep-alive")
		c.Writer.Header().Set("Transfer-Encoding", "chunked")

		c.Next()
	}
}

func UserCounterMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		connecteds++
		log.Print("Notificando nuevo usuario")
		for _, client := range clients {
			log.Printf("Usuario ID: %d", client.ID)
			client.Message <- "Nuevo usuario conectado"
			client.UserCounter <- connecteds
		}

		ctx.Next()
	}
}
