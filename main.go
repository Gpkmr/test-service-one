package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Starting server...")
	http.HandleFunc("/one", SrvOne)
	http.HandleFunc("/health", Health)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
		fmt.Println("Defaulting to port", port)
	}

	fmt.Println("Listening on port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Error on startup", err)
	}
}

func SrvOne(w http.ResponseWriter, r *http.Request) {
	// get service-two URL from env
	srvTwoURL := os.Getenv("SERVICE_TWO_URL")
	if srvTwoURL == "" {
		log.Fatal("Service two URL missing !\n")
	}

	if r.Method == "GET" {
		//fmt.Fprintf(w, "Welcome to service one !\n")

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}

		resp, err := client.Get(srvTwoURL)
		if err != nil {
			fmt.Printf("Error accessing service 2\n%v", err)
			fmt.Fprint(w, err)
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Fprint(w, string(body))
	}
}

func Health(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "Service one is running !\n")
	}
}
