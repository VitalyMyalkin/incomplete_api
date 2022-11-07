package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

type Balance struct {
	Balance int `json:"balance"`
}

type test_struct struct {
	ID string `json:"id"`
}

func get_balance(rw http.ResponseWriter, req *http.Request) {
		decoder := json.NewDecoder(req.Body)
    var t test_struct
    err := decoder.Decode(&t)
    if err != nil {
        panic(err)
    }
		fmt.Println(t.ID)
}

func main() {
  http.HandleFunc("/get_balance", get_balance)
  http.ListenAndServe(":8080", nil)
}
