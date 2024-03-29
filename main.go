package main

import (
	"encoding/csv"
	"encoding/json"
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

func main() {
	// read file
	dat, err := ioutil.ReadFile("resources/engineering_project_businesses.csv")
	if err != nil {
		panic(err)
	}
	inString := string(dat)
	businessList := csv.NewReader(strings.NewReader(inString))

	businessList.Read()
	for {
		record, err := businessList.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		business := new(Business)
		id, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatal(err)
		}
		business.Id = int(id)
		business.Uuid = record[1]
		business.Name = record[2]
		business.Address = record[3]
		business.Address2 = record[4]
		business.City = record[5]
		business.State = record[6]
		zip, _ := strconv.Atoi(record[7])
		if err != nil {
			log.Fatal(err)
		}
		business.Zip = int(zip)
		business.Country = record[8]
		phone, _ := strconv.Atoi(record[9])
		if err != nil {
			log.Fatal(err)
		}
		business.Phone = int(phone)
		business.Website = record[10]
		createdAt, _ := time.Parse("2006-01-02 15:04:05", record[11])
		business.Created_at = createdAt
		businessSlice = append(businessSlice, *business)
	}

	//create request handling

	router := httprouter.New()
	router.GET("/businesses", BusinessList)
	router.GET("/business/:id", BusinessGet)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func BusinessGet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	businessId, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		//log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Business ID invalid"))
	} else {
		if businessId < 0 || businessId > len(businessSlice)-1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Business ID out of range"))
		} else {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			response := businessSlice[businessId]
			json.NewEncoder(w).Encode(response)
		}
	}
}

func BusinessList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pageNumberStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("perPage")
	var pageNumber int
	var pageSize int

	if pageNumberStr == "" {
		pageNumber = 1
		pageNumberStr = "1"
	} else {
		var err error
		pageNumber, err = strconv.Atoi(pageNumberStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Page number invalid"))
			return
		}
	}
	if pageSizeStr == "" {
		pageSize = 50
		pageSizeStr = "50"
	} else {
		var err error
		pageSize, err = strconv.Atoi(pageSizeStr)
		if err != nil || pageSize < 1 || pageSize > len(businessSlice) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Page size invalid"))
			return
		}
	}

	totalPages := len(businessSlice) / pageSize
	if len(businessSlice)%pageSize != 0 {
		totalPages++
	}

	if pageNumber < 1 || pageNumber > totalPages {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Page number out of range"))
		return
	}

	lastPageStr := strconv.Itoa(totalPages)

	var nextPage int
	var prevPage int

	if pageNumber == totalPages {
		nextPage = totalPages
	} else {
		nextPage = pageNumber + 1
	}

	if pageNumber == 1 {
		prevPage = 1
	} else {
		prevPage = pageNumber - 1
	}

	nextPageStr := strconv.Itoa(nextPage)
	prevPageStr := strconv.Itoa(prevPage)

	firstString := "http://localhost:8080/businesses?page=1&perPage=" + pageSizeStr
	prevString := "http://localhost:8080/businesses?page=" + prevPageStr + "&perPage=" + pageSizeStr
	nextString := "http://localhost:8080/businesses?page=" + nextPageStr + "&perPage=" + pageSizeStr
	lastString := "http://localhost:8080/businesses?page=" + lastPageStr + "&perPage=" + pageSizeStr

	lastBusiness := pageNumber * pageSize
	if lastBusiness > len(businessSlice)-1 {
		lastBusiness = len(businessSlice) - 1
	}
	businessList := BusinessListResponse{
		businessSlice[(pageNumber-1)*pageSize : lastBusiness+1],
		Links{
			firstString,
			prevString,
			nextString,
			lastString,
		},
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(businessList)
}
