package main

import (
	"net/http"
	"io"
	"log"
)

func helloServer(w http.ResponseWriter, req *http.Request)  {
	io.WriteString(w, "hello, world!\n")
}

func welcomeServer(w http.ResponseWriter, req *http.Request)  {
	io.WriteString(w, "welcome!\n")
}

func main()  {
	http.HandleFunc("/hello", helloServer)
	http.HandleFunc("/welcome", welcomeServer)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}