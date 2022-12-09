package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
	"strconv"
	"time"
	"github.com/justinas/alice"
)

type City struct {
	Name string
	Area uint64
}


func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var city City
		decode := json.NewDecoder(r.Body)

		err := decode.Decode(&city)

		if err != nil {
			panic(err)
		}

		defer r.Body.Close()

		fmt.Printf("Got %s city with area %d sq miles!\n", city.Name, city.Area)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("201 - Created"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method not allowed"))
	}
}

func filterContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { 
		log.Println("Currently in the check content middleware")
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("415 - Unsupported Media Type"))

			return
		}

		handler.ServeHTTP(w, r)
	})
}

func setServerTimeoutCookie(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
		tempCookie := http.Cookie{Name: "Server Timout(UTC)", Value: strconv.FormatInt(time.Now().Unix(), 10)}
		http.SetCookie(w, &tempCookie)
	})
}

func main() {

	originalHandler := http.HandlerFunc(postHandler)

	chain := alice.New(filterContentType, setServerTimeoutCookie).Then(originalHandler)
	http.Handle("/city", chain)

	http.ListenAndServe(":8000", nil)
}