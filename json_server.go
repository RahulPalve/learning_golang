package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Note struct {
	Title       string
	Description string
	User        string
}

type User struct {
	Username string
	Password string
}

type DB struct {
	Users []User
	Notes []Note
}

func LoadDB(file string) *DB {
	db := new(DB)
	data, _ := os.ReadFile(file)
	json.Unmarshal(data, db)
	return db

}
func (db *DB) GetAll(w http.ResponseWriter, r *http.Request) {

	result, _ := json.Marshal(db)
	w.Write(result)
}

func (db *DB) GetByUser(w http.ResponseWriter, r *http.Request) {
	var filtered []Note
	username := r.URL.Query()["username"][0]
	fmt.Printf("%v %T", username, username)
	for _, note := range db.Notes {
		if note.User == username {
			filtered = append(filtered, note)
		}

	}
	result, _ := json.Marshal(filtered)
	w.Write(result)
}

func main() {
	db := LoadDB("db.json")
	http.HandleFunc("/", db.GetAll)
	http.HandleFunc("/search", db.GetByUser)
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		panic(err)
	}
}
