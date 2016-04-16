package main

import (
	//"encoding/csv"
	"fmt"
	"html"
	//"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Business struct {
	Id         uint32
	Uuid       string
	Name       string
	Address    string
	Address2   string
	City       string
	State      string
	Zip        uint16
	Country    string
	Phone      uint32
	Website    string
	Created_at time.Time
}

func main() {
	dat, err := ioutil.ReadFile("resources/engineering_project_businesses.csv")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dat))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
