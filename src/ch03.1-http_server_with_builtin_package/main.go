package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandlerFunc)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func helloHandlerFunc(rw http.ResponseWriter, req *http.Request) {
	res := map[string]any{
		"say": "Hello",
		"to":  "Gophers",
	}
	json.NewEncoder(rw).Encode(res)
}
