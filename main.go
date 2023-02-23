package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Students struct {
	ID   int
	Name string
}

type Book struct {
	ID          int
	Name        string
	StudentName string
}

type Teacher struct {
	ID   int
	Name string
}

func main() {

	connect := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "1234", "golang")
	db, err := sql.Open("postgres", connect)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	student := &Students{
		Name: "Sindarov Jorabek",
	}

	err = db.QueryRow("INSERT INTO student (name) Values ($1) RETURNING id", student.Name).Scan(&student.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(student)

	student.Name = "kimdir"

	result, err := db.Exec("UPDATE student SET name=$1 where id=$2", student.Name, student.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(student)
	fmt.Println(result.RowsAffected())
}