package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Post struct {
	ID       int    `json:"ID"`
	Name     string `json:"Name"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Phone    string `json:"Phone"`
	Country  string `json:"Country"`
	City     string `json:"City"`
	Address  string `json:"Address"`
}

func main() {
	var client = &http.Client{Timeout: 10 * time.Second}
	url := "http://127.0.0.1:8081/users"

	resp, err := client.Get(url)

	if err != nil {
		panic(err.Error)
	}

	var post []Post

	err = json.NewDecoder(resp.Body).Decode(&post)

	if err != nil {
		panic(err.Error)
	}
	for d := range post {
		fmt.Println(post[d])
	}
}
