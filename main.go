package main

import (
"fmt"
"log"
"net/http"
)

func order(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.URL.RawQuery)
	fmt.Fprintf(w,`This is the order page`)
}
func about(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.URL.RawQuery)
	fmt.Fprintf(w,`This is the about page
					tell customers who we are`)
}
func contact(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.URL.RawQuery)
	fmt.Fprintf(w,`This is the contact page
	tell customers how to contact us`)
}
func subscribe(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.URL.RawQuery)
	fmt.Fprintf(w,`This is the subcribe page. get info from subscriber add to database`)
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.RawQuery)
	fmt.Fprintf(w, `
          ##         .
    ## ## ##        ==
 ## ## ## ## ##    ===
/"""""""""""""""""\___/ ===
{                       /  ===-
\______ O           __/
 \    \         __/
  \____\_______/

	
Hello from Docker!

`)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/order", order)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/subscribe", subscribe)
	log.Fatal(http.ListenAndServe(":8080", nil))
	
}


