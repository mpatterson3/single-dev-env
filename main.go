package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl = template.New("foo")

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func order(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.RawQuery)
	fmt.Fprintf(w, `This is the order page`)
}
func about(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.RawQuery)
	fmt.Fprintf(w, `This is the about page
					tell customers who we are`)
}
func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.RawQuery)
	fmt.Fprintf(w, `This is the contact page
	tell customers how to contact us`)
}
func subscribe(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.RawQuery)
	fmt.Fprintf(w, `This is the subcribe page. get info from subscriber add to database`)
}
func handler(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println(r.URL.RawQuery)
	// 	fmt.Fprintf(w, `
	//           ##         .
	//     ## ## ##        ==
	//  ## ## ## ## ##    ===
	// /"""""""""""""""""\___/ ===
	// {                       /  ===-
	// \______ O           __/
	//  \    \         __/
	//   \____\_______/

	// Hello from Docker!

	//`)
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/order", order)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/subscribe", subscribe)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
