package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)
var host string
var port int
var indentLength int
var logger string
	
func parseFlags() {
	flag.StringVar(&host, "host", "0.0.0.0", "Host you want to run on, by default 0.0.0.0")
	flag.IntVar(&port, "port", 8000, "port you want to expose")
	flag.IntVar(&indentLength, "indent_length", 5, "Indent you want for nested values.")
	flag.StringVar(&logger, "logger", "logrus", "Logger library you want to use")
	flag.Parse()
}
func main() {
	parseFlags()
	logger := NewLogger(logger)
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		logger.Printf("error starting tcp listener: %v\n", err)
		os.Exit(1)
	}
	server := &http.Server{
		Handler: handler{logger},
	}
	logger.Printf("Starting HTTP server on %s:%d", host, port)

	if err := server.Serve(listen); err != nil {
		logger.Printf("error in serving handler with tcp server: %v\n", err)
	}
}

type handler struct {
	logger Logger
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Printf("\n/-------------- HTTP request received --------------/\n")
	h.logger.Printf("From: %s\n", r.RemoteAddr)
	h.logger.Printf("Request URI: %s\n", r.RequestURI)
	h.logger.Printf("Method: %s\n", r.Method)

	h.logger.Printf("\nHeaders:\n")
	valuesPrettyPrint(5, r.Header)

	h.logger.Printf("\nQuery Params:\n")
	valuesPrettyPrint(5, r.URL.Query())

	r.ParseForm()
	h.logger.Printf("\nPost Params:\n")
	valuesPrettyPrint(5, r.Form)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.logger.Printf("Error reading request body:%v\n", err)
	}
	defer r.Body.Close()

	h.logger.Printf("\nRAW Body:\n %s\n", body)

	h.logger.Printf("\n/-------------- END --------------/\n")
}


