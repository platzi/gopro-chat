package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

const (
	USER     = "platzi"
	PASSWORD = "platzi"
	NAME     = "platzi_test"
)

func main() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		USER, PASSWORD, NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	fmt.Println("--- Insertando valores ---")

	var lastInsertId int
	err = db.QueryRow("INSERT INTO users(username,email,created_at) VALUES($1,$2,$3) returning id;", "platzi_user", "email@platzi.com", "2020-12-09").Scan(&lastInsertId)
	checkErr(err)
	fmt.Println("Último id =", lastInsertId)

	fmt.Println("--- Actualizando ---")
	stmt, err := db.Prepare("update users set username=$1 where id=$2")
	checkErr(err)

	res, err := stmt.Exec("platzi_update", lastInsertId)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect, "Filas afectadas")

	fmt.Println("--- Consultando ---")
	rows, err := db.Query("SELECT * FROM users")
	checkErr(err)

	for rows.Next() {
		var id int
		var username string
		var email string
		var created time.Time
		err = rows.Scan(&id, &username, &email, &created)
		checkErr(err)
		fmt.Println("id | username | email | created_at ")
		fmt.Printf("%3v | %8v | %6v | %6v\n", id, username, email, created)
	}

	fmt.Println("# Deleting")
	stmt, err = db.Prepare("delete from users where id=$1")
	checkErr(err)

	res, err = stmt.Exec(lastInsertId)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect, "rows changed")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
