package main

import (
	database "github.com/resistance22/rdbms/Database"
)

// func setup() {
// 	rand.Seed(time.Now().UnixNano())
// }

func main() {
	database.SaveData1("./Data/data", []byte("Hello World"))
}
