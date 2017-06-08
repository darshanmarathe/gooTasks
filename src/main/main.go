package main 

import (
	"net/http"
)

func main(){
	println("starting...")
http.HandleFunc("/", func (w http.ResponseWriter , r *http.Request)  {
	w.Write([]byte("Home page"))
	
})

http.HandleFunc("/contact" , func (w http.ResponseWriter , r *http.Request)  {
	w.Write([]byte("Login page"))
})

   println("go app running on port 3000")
	http.ListenAndServe(":3000" , nil)

}