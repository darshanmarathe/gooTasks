package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
	"io"
	"strings"
)

func main() {

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		f, err := os.Open("public" + req.URL.Path)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}
		defer f.Close()
		contenttype :="text/plain" 
		switch {
			case strings.HasSuffix(req.URL.Path , "css"):
				contenttype ="text/css"
			case strings.HasSuffix(req.URL.Path , "html"):
				contenttype ="text/html"
		

		}
		res.Header().Add("Content-type" , contenttype)
		io.Copy(res , f)

	})
	fmt.Println("Listingng on port 3000")
	http.ListenAndServe(":3000", nil)
}
