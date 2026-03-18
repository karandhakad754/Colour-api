package main

import (
	"encoding/json"
	"net/http"
	"os"
)

type Result struct {
	Number int    `json:"number"`
	Color  string `json:"color"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	res := Result{Number: 7, Color: "green"}
	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/result", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}

	http.ListenAndServe(":"+port, nil)
}
