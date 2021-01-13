package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

var host = "0.0.0.0"
var port = 8000

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		panic(err)
	}
	server := &http.Server{
		Handler: handler{},
	}
	fmt.Printf("Starting HTTP server on %s:%d", host, port)
	if err := server.Serve(listen); err != nil {
		panic(err)
	}
}

type handler struct {
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\n/-------------- HTTP request received --------------/")
	fmt.Printf("From: %s\n", r.RemoteAddr)
	fmt.Printf("Request URI: %s\n", r.RequestURI)
	fmt.Printf("Method: %s\n", r.Method)

	fmt.Printf("\nHeaders:\n")
	ValuesPrettyPrint(5, r.Header)

	fmt.Printf("\nQuery Params:\n")
	ValuesPrettyPrint(5, r.URL.Query())

	r.ParseForm()
	fmt.Printf("\nPost Params:\n")
	ValuesPrettyPrint(5, r.Form)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading request body:", err)
	}
	defer r.Body.Close()

	fmt.Printf("\nRAW Body:\n %s", body)

	fmt.Println("\n/-------------- END --------------/\n")
}

func ValuesPrettyPrint(indentCount int, values map[string][]string) {
	for k, v := range values {
		fmt.Printf("%s %s => %v\n", strings.Repeat(" ", indentCount), k, v)
	}
}
