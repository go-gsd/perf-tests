package main

import (
    "fmt"
    "net/http"
	"strconv"
)

func add(w http.ResponseWriter, req *http.Request) {

	fmt.Println("Incomming request:", req)
	
	switch req.Method {
		case "GET":
			val1, err := strconv.Atoi(req.URL.Query().Get("val1"))
			val2, err := strconv.Atoi(req.URL.Query().Get("val2"))
					
			fmt.Println("Incomming request: ", val1, ",", val2, ",", err)
			
			fmt.Fprintf(w, strconv.Itoa(val1+val2))
		default:
			fmt.Fprintf(w, "error")
	}
	
    
}

func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {

    http.HandleFunc("/add", add)
    http.HandleFunc("/headers", headers)

	port := ":8090"

	fmt.Println("Starting web app on port: ", port)
    http.ListenAndServe(port, nil)
}