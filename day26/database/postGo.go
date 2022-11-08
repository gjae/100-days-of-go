package main
import (
    "fmt"
    "math/rand"
    "time"
    database "gjae/database/db"
)

var MIN = 0
var MAX = 26
func random(min, max int) int {
    return rand.Intn(max-min) + min
}
func getString(length int64) string {
    startChar := "A"
    temp := ""
    var i int64 = 1
    for {
        myRand := random(MIN, MAX)
        newChar := string(startChar[0] + byte(myRand))
        temp = temp + newChar
        if i == length {
            break
        }
        i++
    }
    return temp
}

func main() {
    database.Hostname = "localhost"
    database.Port = 5432
    database.Username = "mtsouk"
    database.Password = "pass"
    database.Database = "go"

    data, err := database.ListUsers()
    if err != nil {
        fmt.Println(err)
        return
    }
    for _, v := range data {
        fmt.Println(v)
    }
    SEED := time.Now().Unix()
    rand.Seed(SEED)
    random_username := getString(5)
    t := database.Userdata{
        Username:    fmt.Sprintf("Gio_%s", random_username),
        Name:        "Mihalis",
        Surname:     "Tsoukalos",
        Description: "This is me!"}
    id := database.AddUser(t)
    if id == -1 {
        fmt.Println("There was an error adding user", t.Username)
    }
    err = database.DeleteUser(id)
    if err != nil {
        fmt.Println(err)
    }
    // Trying to delete it again!
    err = database.DeleteUser(id)
    if err != nil {
        fmt.Println(err)
    }

	id = database.AddUser(t)
    if id == -1 {
        fmt.Println("There was an error adding user", t.Username)
    }
    t = database.Userdata{
        Username:    fmt.Sprintf("Gio_%s", random_username),
        Name:        "Mihalis",
        Surname:     "Tsoukalos",
        Description: "This might not be me!"}
    err = database.UpdateUser(t)
    if err != nil {
        fmt.Println(err)
    }

    t, err = database.RetrieveUser(t.Username)

    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("User: ", t)

    // database.DeleteAllRecords()
}
