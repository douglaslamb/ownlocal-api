package main

import (
	"encoding/csv"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
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
	// read file
	dat, err := ioutil.ReadFile("resources/engineering_project_businesses.csv")
	if err != nil {
		panic(err)
	}
	inString := string(dat)
	businessList := csv.NewReader(strings.NewReader(inString))
	businessSlice := make([]Business, 0, 10)

	fmt.Println(businessList.Read())
	for {
		record, err := businessList.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		business := new(Business)
		id, _ := strconv.Atoi(record[0])
		business.Id = uint32(id)
		business.Uuid = record[1]
		business.Name = record[2]
		business.Address = record[3]
		business.Address2 = record[4]
		business.City = record[5]
		business.State = record[6]
		zip, _ := strconv.Atoi(record[7])
		business.Zip = uint16(zip)
		business.Country = record[8]
		phone, _ := strconv.Atoi(record[9])
		business.Phone = uint32(phone)
		business.Website = record[10]
		createdAt, _ := time.Parse("2006-01-02 15:04:05", record[11])
		business.Created_at = createdAt
		businessSlice = append(businessSlice, *business)
	}
	fmt.Println(businessSlice[10])

	//create request handling

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	//log.Fatal(http.ListenAndServe(":8080", nil))
}
