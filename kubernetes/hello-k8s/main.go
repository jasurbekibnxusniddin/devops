package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello k8s!")
	})

	log.Printf("Listening on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
