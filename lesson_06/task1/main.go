package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Host       string              `json:"host: "`
	UserAgent  string              `json:"user_agent: "`
	RequestUri string              `json:"request_uri: "`
	Headers    map[string][]string `json:"headers: "`
}

func (u User) ToJson() string {
	j, _ := json.Marshal(u)
	return string(j)
}

func currentUser(r *http.Request) User {
	var u User
	u.Headers = make(map[string][]string)
	u.Host = r.Host
	u.UserAgent = r.UserAgent()
	u.RequestUri = r.RequestURI
	u.Headers["Accept"] = r.Header["Accept"]
	u.Headers["User-Agent"] = r.Header["User-Agent"]

	return u
}

func handler(w http.ResponseWriter, r *http.Request) {
	u := currentUser(r)
	fmt.Fprintf(w, u.ToJson())
}

func main() {
	const port = 8080

	fmt.Printf("Launching server on port: %d \n\n", port)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}