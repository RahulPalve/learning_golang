package main
import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)
func log_err(err error){
	if err!=nil{
		fmt.Println(err)
		panic(err)
	}
}
func main(){
	connection_string := "user=vaibhav dbname=packager password=vbhv3301 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection_string)
	log_err(err)

	err = db.Ping()
	log_err(err)
	defer db.Close()
}