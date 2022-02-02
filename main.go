package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func asciiGen(input string,banner string)string{
	return input
}

func ascii(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		tmpl,_ := template.ParseFiles("templates/index.html")
		tmpl.Execute(w,nil)
	} else {
		r.ParseForm()
		asciiInput := r.Form["asciiInput"]
		asciiBanner := r.Form["banner"]
		fmt.Println(asciiInput,asciiBanner)
		asciiRaw := asciiGen(asciiInput[0],asciiBanner[0])
		tmpl,_ := template.ParseFiles("/templates/index.html")
		tmpl.Execute(w,asciiRaw)
	}
}

func main() {
	http.HandleFunc("/", ascii)
	err := http.ListenAndServe(":8080",nil)
	if err != nil{
		log.Fatalf("500 internal server error : %s",err.Error())
	}
}