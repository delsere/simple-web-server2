package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	//Parse url parameters passed, then parse the response packet for the POST body (request body)
	// attention: If you do not call ParseForm method, the following data can not be obtained form
	r.ParseForm()

	// print information on server side.
	fmt.Println(r.Form)

	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	fmt.Fprintf(w, "Hello astaxie!") // write data to response
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method

	if r.Method == "GET" {
		t, err := template.ParseFiles("template/login.gtpl")
		if err != nil {
			log.Fatal("Parsing login.gtpl failed: ", err)
		}
		t.Execute(w, nil)
	} else {

		r.ParseForm()

		// logic part of log in
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method

	if r.Method == "GET" {
		t, err := template.ParseFiles("register.gtpl")
		if err != nil {
			log.Fatal("Parsing register.gtpl failed: ", err)
		}
		t.Execute(w, nil)
	} else {

		r.ParseForm()

		// logic part of log in
		fmt.Println("nome:", r.Form["name"])
		fmt.Println("cognome:", r.Form["cognome"])
		fmt.Println("email:", r.Form["email"])
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		/*
				nome := r.Form["name"]
				cognome := r.Form["cognome"]

				fmt.Fprintf(w, []byte(nome))


					email := r.Form["username"]
					password := r.Form["password"]

			lenght := len(nome)

			if len(nome) == 0 {
				fmt.Fprintf(w, "Nome non inserito "+string(lenght))
			} else {
				fmt.Fprintf(w, "Nome OK!"+string(lenght))
			}
			if len(cognome) == 0 {
				fmt.Fprintf(w, "Cognome non inserito")
			}
		*/
	}
}

func main() {
	//Imposto le HandleFunc
	setupHandlerFunctions()
	startWebServer()
}

func setupHandlerFunctions() {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
}

func startWebServer() {
	// setting listening port
	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
