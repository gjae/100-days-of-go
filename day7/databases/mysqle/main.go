package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

/**
* Prepare function is used for write operation
* require call to Exec function (from stmt returned value)
*
* While Query funcion is used for Read Operations (SELECT)
* no require make call to Exec function and return rows and error values
* rows is a cursor so that require Next call for walk throught results
* rows.Scan assign columns to vars receiveds as pointers
*/
func main() {
	db, err := sql.Open("mysql", "root:root@/godb")
	defer db.Close()

	checkErr(err)

	// insert
	stmt, err := db.Prepare("INSERT userinfo SET  username=?, departname=?, created=?")
	checkErr(err)

	res, err := stmt.Exec("Giovanny", "Informatica", "2022-10-05")
	checkErr(err)

	// Get data
	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	// Update info
	stmtUpdate, err := db.Prepare("UPDATE userinfo SET departname=? WHERE uid=?")
	checkErr(err)

	resUpdate, err := stmtUpdate.Exec("SISTEMAS", id)
	checkErr(err)

	affected, err := resUpdate.RowsAffected()
	checkErr(err)

	fmt.Println(affected)

	// Query
	rows, err := db.Query("SELECT uid, username, departname FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username, departname string
		
		// Shoud be passed in the order of select columns is indicated
		err = rows.Scan(&uid, &username, &departname)
		checkErr(err)

		fmt.Println(uid, username, departname)
	}

	// Delete
	delete, err := db.Prepare("DELETE FROM userinfo WHERE uid=?")
	checkErr(err)

	_, err = delete.Exec(id)
	checkErr(err)

}