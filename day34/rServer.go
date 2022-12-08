package main

import (
	"fmt"
	"time"
	"net/http"
	"log"
	"io"
	"encoding/json"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}


var user User

var DATA = make(map[string]string)

var PORT = ":8081"

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("PATH ", r.URL.Path, " from ", r.Host)
	w.WriteHeader(http.StatusNotFound)
	Body := "Thanks for visiting"

	fmt.Fprintf(w, "%s", Body)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("PATH ", r.URL.Path, " from ", r.Host)
	t := time.Now().Format(time.RFC1123)
	Body := "The current time is : "+t+"\n"

	fmt.Fprintf(w, "%s", Body)
}


func addHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("PATH ", r.URL.Path, " from ", r.Host)

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method "+r.Method+" Is not allowed")
		return
	}

	d, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(d, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	if user.Username == "" || user.Password == "" {
		http.Error(w, "Error: ", http.StatusBadRequest)
		return
	}

	DATA[user.Username] = user.Password
	w.WriteHeader(http.StatusCreated)
	log.Println(DATA)
}


func getHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("PATH ", r.URL.Path, " from ", r.Host)

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	d, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		log.Println(err)
		return
	}

	err = json.Unmarshal(d, &user)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}


	
	if _, ok := DATA[user.Username]; !ok && user.Username != "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", d)

}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(defaultHandler))
	mux.Handle("/time", http.HandlerFunc(timeHandler))
	mux.Handle("/user/add", http.HandlerFunc(addHandler))
	mux.Handle("/user/get", http.HandlerFunc(getHandler))

	s := &http.Server {
		Addr:			PORT,
		Handler:		mux,
		IdleTimeout:	10 * time.Second,
		ReadTimeout:	time.Second,
		WriteTimeout:	time.Second,
	}



	fmt.Println("Running server at ", PORT)
	err := s.ListenAndServe()

	if err != nil {
		fmt.Println(err)
		return
	}
}