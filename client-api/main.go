package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		log.Fatalln(err)
	}

	// fmt.Println(res)

	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
	}

	bodyString := string(bodyBytes)

	fmt.Println(bodyString)

	var todoStruct Todo
	json.Unmarshal(bodyBytes, &todoStruct)

	fmt.Println(todoStruct)

}
