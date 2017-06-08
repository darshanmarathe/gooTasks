package main

import (
	"fmt"
	"net/http"
)

type myHandler struct {
	greeting string
}

func (mh myHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte(fmt.Sprintf("%v world", mh.greeting)))
}

func main() {
	http.Handle("/", &myHandler{greeting: "Hello .."})
	http.HandleFunc("/Contact", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Contact us"))
	})
	fmt.Println("Listingng on port 3000")
	http.ListenAndServe(":3000", nil)
}
