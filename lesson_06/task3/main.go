package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "index.html")
	case "POST":
		name := r.FormValue("name")
		address := r.FormValue("address")

		cookie := &http.Cookie {
			Name:  "token",
			Value: name+":"+address,
		}
		http.SetCookie(w, cookie)

		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Address = %s\n", address)
		fmt.Fprintf(w, "Token = %s", cookie.Value)

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	const port = 8080
	fmt.Printf("Launching server on port: %d \n\n", port)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
