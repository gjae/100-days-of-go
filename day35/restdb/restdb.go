package main

import (
	"fmt"
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

var (
	Host 		= "localhost"
	Port 		= 5432
	Username	= "mtsouk"
	Password 	= "pass"
	Database 	= "restapi"
)

type User struct {
	ID			int
	Username	string
	Password	string
	Lastlogin	int
	admin		int
	active		int
}

func ConnectPostgres() *sql.DB {
	conn : = fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", Host, Port, Username, Password, Database)

	db, err = sql.Open("postgres", conn)
	if err != nil {
		log.Println(err)
		return nil
	}

	return db
}

func FindUserByID(ID int) User {
	db := ConnectPostgres()
	if db == nil {
		fmt.Println("No connect")
		return User{}
	}

	defer db.Close()

	rows, err := db.Query(
		fmt.Sprintf("SELECT * FROM users WHERE id=%d", ID)
	)

	if err != nil {
		fmt.Println("Error select query")
		return User{}
	}

	var ID, lastlogin, admin, active int
	var username, password string

	for rows.Next() {
		err = rows.Scan(&ID, &username, &password, &lastlogin, &admin, &active)

		if err != nil {
			log.Println(err)
			return User{}
		}

		return User{ID, username, password, lastlogin, admin, active}
	}

	return User{}
}

func DeleteUser(ID int) bool {
	db := ConnectPostgres()
	if db == nil {
		log.Println("Cannot connect to PostgreSQL!")
		return false
	}

	defer db.Close()

	t := FindUserByID(ID)

	if t.ID == 0 {
		log.Println("User ", ID, " does not exists")
		return false
	}

	stmt, err := db.Prepare("DELETE FROM users WHERE id = %1")

	if err != nil {
		log.Println(err)
		return false
	}

	_, err = stmt.Exec(ID)
	if err != nil {
		log.Println("Delete error ", error)
		return false
	}

	return true
}

func ListAllUsers() []User {
	db := ConnectPostgres()
	if db == nil {
		fmt.Println("Cannot connect to postresql!")
		return []User{}
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users\n")
	if err != nil {
		log.Println(err)
		return []User{}
	}

	all := []User{}
	var c1 int
	var c2, c3 string
	var c4 int64
	var c5, c6 int

	for rows.Next() {
		err = rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6)
		temp := User{c1, c2, c3, c4, c5, c6}
		all = append(all, temp)
	}

	log.Println(all)
	return all
}

func IsUserValid(u User) bool {
	db := ConnectPostgres()
	if db == nil {
		log.Println("Cannot connect to postgrsql")
		return false
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users WHERE Username = $1", u.Username)

	if err != nil {
		log.Println(err)
	}

	var temp User
	var ID, lastlogin, admin, active int
	var username, password string

	for rows.Next() {
		err = rows.Scan(&ID, &username, &password, &lastlogin, &admin, &active)

		if err != nil {
			log.Println(err)
			return User{}
		}

		temp = User{ID, username, password, lastlogin, admin, active}
	}

	if u.Username == temp.Username && u.Password == tmp.Password {
		return true
	}

	return false

}