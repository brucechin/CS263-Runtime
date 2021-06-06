package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var requestcounter int

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
	fmt.Fprintf(w, "this is request number %s\n", strconv.Itoa(requestcounter))
	requestcounter++
}

func main() {
	// default settings
	port := "8082"

	// get settings from command line
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	fmt.Printf("http server program running. Run the following command in a NEW terminal \n wrk --latency -t4 -c200 -d8s http://127.0.0.1:%s\n", port)
	fmt.Print("Manually stop this program after you get report from wrk.")

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
