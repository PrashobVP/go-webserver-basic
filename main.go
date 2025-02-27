package main

import (

	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
 if err := r.ParseForm(); err !=nil{
	fmt.Fprintf(w, "ParseForm() err: %v", err)
	return
 }
 fmt.Fprintf(w, "POST request Successful")
 name := r.FormValue("name")
 address := r.FormValue("address")
 fmt.Fprintf(w, "Name = %s\n", name)
 fmt.Fprintf(w, "Address =%s\n", address)
}


func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET"{
		http.Error(w, "method is not support", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "hello!")
}

func main(){

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Statring server at port 8082\n")
	if err:=http.ListenAndServe(":8082", nil); err!=nil{
		log.Fatal(err)
	}
}