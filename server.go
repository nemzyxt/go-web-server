// Author : Nemuel Wainaina

package main

import (
	"fmt"
	"log"
	"net/http"
)

func contactHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Error : %v", err)
		return
	}

	if r.Method == "GET" {
		http.ServeFile(w, r, "./static/contact.html")
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	msg := r.FormValue("msg")
	fmt.Fprintf(w, "Name : %v\n", name)
	fmt.Fprintf(w, "Ëmail : %v\n", email)
	fmt.Fprintf(w, "Message : %v", msg)
}

func abtHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" {
		http.Error(w, "Requested page not found !", http.StatusNotFound)
	}
	if r.Method != "GET" {
		http.Error(w, "Method not supported !", http.StatusBadRequest)
	}

	http.ServeFile(w, r, "./static/about.html")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/about", abtHandler)

	fmt.Println("[*] Server started on 127.0.0.1:80 ...")

	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}