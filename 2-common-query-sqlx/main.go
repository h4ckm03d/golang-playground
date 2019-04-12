package main

import "fmt"

var columnsUsers = []string{}

func ColumnUser() []string {
	if len(columnsUsers) == 0 {
		columnsUsers, _ = GetColumns(User{}, "db")
	}
	return columnsUsers
}

type User struct {
	ID        int64  `db:"id"`
	Age       int    `db:"age"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
}

func main() {
	for i := 0; i < 1000; i++ {
		fmt.Println(ColumnUser())
	}
}
