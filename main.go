package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
)

//global business list
var businessSlice []Business = make([]Business, 0, 10)

type Business struct {
	Id         uint32    `json:"id"`
	Uuid       string    `json:"uuid"`
	Name       string    `json:"name"`
	Address    string    `json:"address"`
	Address2   string    `json:"address2"`
	City       string    `json:"city"`
	State      string    `json:"state"`
	Zip        uint16    `json:"zip"`
	Country    string    `json:"country"`
	Phone      uint32    `json:"phone"`
	Website    string    `json:"website"`
	Created_at time.Time `json:"created_at"`
}

type BusinessListResponse struct {
	BusinessList []Business `json:"businesses"`
	//add pagination later
}

func main() {
	// read file
	dat, err := ioutil.ReadFile("resources/engineering_project_businesses.csv")
	if err != nil {
		panic(err)
	}
	inString := string(dat)
	businessList := csv.NewReader(strings.NewReader(inString))

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
	fmt.Println(businessSlice[12])

	//create request handling

	router := httprouter.New()
	router.GET("/businesses/", BusinessList)
	//router.GET("/business/:id", Business)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func BusinessList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	response := BusinessListResponse{businessSlice}
	json.NewEncoder(w).Encode(response)
	//fmt.Println(businessSlice[0])
}
