package main

import (
	"io"
	"os"
	"log"
	"github.com/emicklei/go-restful"
	"fmt"
	"time"
	"net/http"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	Id int
	Name string
	Author string
	Isbn uint64
}

func pingTime(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, fmt.Sprintf("%s", time.Now()))
}

func DB(driver, path string) (*sql.DB, error) {
	return sql.Open(driver, path)
}

func createTable(db *sql.DB) {
	f, _ := os.ReadFile("./table.sql")
	statement, err := db.Prepare(string(f))

	if err != nil {
		panic(err)
	}

	statement.Exec()
	log.Println("BOOKS table has been created successfuly")
}

func DbOperations(db *sql.DB) {
	// Insert operation
	statement, _ := db.Prepare("INSERT INTO books(name, author, isbn) VALUES(?, ?, ?)")

	statement.Exec("A table of two cities", "Charles Dickens", 140430547)

	// Read operation
	rows, _ := db.Query("SELECT id, name, author, isbn FROM books")

	for rows.Next() {
		var tmpBook Book
		rows.Scan(&tmpBook.Id, &tmpBook.Name, &tmpBook.Author, &tmpBook.Isbn)
		log.Println("A book has been readed")
		log.Printf("Book id='%d' , Book name='%s', Book author='%s', ISBN=%d\n", tmpBook.Id, tmpBook.Name, tmpBook.Author, tmpBook.Isbn)
	}
	rows.Close()

	// Update operation
	statement, _ = db.Prepare("UPDATE books SET name=? WHERE id=?")

	statement.Exec("Golang is cool", 1)
	log.Println("Book has been updated")

	// Delete operation
	statement, _ = db.Prepare("DELETE FROM books WHERE id = ?")
	statement.Exec(1)
}

func main() {

	db, err := DB("sqlite3", "./books.db")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	createTable(db)
	DbOperations(db)

	webservice := new(restful.WebService)

	webservice.Route(webservice.GET("/ping").To(pingTime))

	restful.Add(webservice)

	http.ListenAndServe(":8000", nil)
}