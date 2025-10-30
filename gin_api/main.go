package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func demoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	res := map[string]string{
		"code":    "200",
		"message": "Hello World",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
func main() {
	log.Println("Server is starting now....")
	http.HandleFunc("/demo", demoHandler)

	err := http.ListenAndServe("127.0.0.1:8000", nil)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
