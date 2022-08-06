package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Starting server...")
	http.HandleFunc("/one", SrvOne)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
		fmt.Println("Defaulting to port", port)
	}

	fmt.Println("Listening on port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Panic("Error on startup", err)
	}
}

func SrvOne(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//fmt.Fprintf(w, "Welcome to service one !\n")
		resp, err := http.Get("http://127.0.0.1:8082/two")
		if err != nil {
			fmt.Printf("Error accessing service 2\n%v", err)
			fmt.Fprint(w, err)
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Fprint(w, string(body))
	}
}
