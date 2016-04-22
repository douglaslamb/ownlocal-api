package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestBusinessListRequest(t *testing.T) {
	t.Log("Requesting entire business list")
	response, err := http.Get("http://localhost:8080/businesses?page=1&perPage=50000")
	defer response.Body.Close()
	if err != nil {
		t.Error("GET generated an error")
	}
	jsonDataFromHttp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Error("Conversion of response to string generated an error")
	}

	var jsonData BusinessListResponse

	err = json.Unmarshal([]byte(jsonDataFromHttp), &jsonData)
	if err != nil {
		t.Error("Conversion of response to JSON generated an error")
	}

	//check first business
	firstBusinessTime, _ := time.Parse("2006-01-02 15:04:05", "2012-12-10 16:17:58")

	firstBusiness := Business{
		0,
		"2859d6e0-1cb9-4fe9-bc00-97823a9fa4cb",
		"Yundt-Flatley",
		"1386 Lim Brooks",
		"Suite 517",
		"Lake Betsy",
		"IA",
		19416,
		"US",
		4034880719,
		"http://www.halvorson.com/",
		firstBusinessTime,
	}

	if !reflect.DeepEqual(jsonData.BusinessList[0], firstBusiness) {
		t.Error("business with id 0 as returned by API is incorrect")
	}

	//check last business
	lastBusinessTime, _ := time.Parse("2006-01-02 15:04:05", "2014-02-24 00:03:02")

	lastBusiness := Business{
		49999,
		"4be0c534-4d8e-4654-a92b-e833fe62ab8b",
		"Kuhic, Bayer and Johns",
		"48933 Cormier Islands",
		"",
		"Lake Margarettside",
		"NV",
		85524,
		"US",
		5270180703,
		"http://www.gleason.com/",
		lastBusinessTime,
	}

	if !reflect.DeepEqual(jsonData.BusinessList[49999], lastBusiness) {
		t.Error("business with id 49999 as returned by API is incorrect")
	}

	//check business in middle
	// id 5729
	middleBusinessTime, _ := time.Parse("2006-01-02 15:04:05", "2012-04-27 06:22:24")

	middleBusiness := Business{
		5729,
		"de8099f9-5fca-4c00-873d-71e6bb0047fa",
		"Bins Group",
		"42826 Geneva Shores",
		"",
		"Feeneyfurt",
		"IL",
		50873,
		"US",
		1948817012,
		"http://www.wintheiser-bosco.info/",
		middleBusinessTime,
	}

	if !reflect.DeepEqual(jsonData.BusinessList[5729], middleBusiness) {
		t.Error("business with id 5729 as returned by API is incorrect")
	}
}

func TestOneBusinessRequest(t *testing.T) {
	//just testing a few businesses with a few ids
	t.Log("Requesting three businesses with different id")

	var responseBusiness0 Business
	var responseBusiness40988 Business
	var responseBusiness49999 Business

	// business0

	response0, err := http.Get("http://localhost:8080/business/0")
	if err != nil {
		t.Error("GET generated an error")
	}
	jsonDataFromHttp0, err := ioutil.ReadAll(response0.Body)
	if err != nil {
		t.Error("Conversion of response to string generated an error")
	}

	err = json.Unmarshal([]byte(jsonDataFromHttp0), &responseBusiness0)
	if err != nil {
		t.Error("Conversion of response to JSON generated an error")
	}
	defer response0.Body.Close()

	//business40988

	response40988, err := http.Get("http://localhost:8080/business/40988")
	if err != nil {
		t.Error("GET generated an error")
	}
	jsonDataFromHttp40988, err := ioutil.ReadAll(response40988.Body)
	if err != nil {
		t.Error("Conversion of response to string generated an error")
	}

	err = json.Unmarshal([]byte(jsonDataFromHttp40988), &responseBusiness40988)
	if err != nil {
		t.Error("Conversion of response to JSON generated an error")
	}
	defer response40988.Body.Close()

	//business49999

	response49999, err := http.Get("http://localhost:8080/business/49999")
	if err != nil {
		t.Error("GET generated an error")
	}
	jsonDataFromHttp49999, err := ioutil.ReadAll(response49999.Body)
	if err != nil {
		t.Error("Conversion of response to string generated an error")
	}

	err = json.Unmarshal([]byte(jsonDataFromHttp49999), &responseBusiness49999)
	if err != nil {
		t.Error("Conversion of response to JSON generated an error")
	}
	defer response49999.Body.Close()

	firstBusinessTime, _ := time.Parse("2006-01-02 15:04:05", "2012-12-10 16:17:58")
	business0 := Business{
		0,
		"2859d6e0-1cb9-4fe9-bc00-97823a9fa4cb",
		"Yundt-Flatley",
		"1386 Lim Brooks",
		"Suite 517",
		"Lake Betsy",
		"IA",
		19416,
		"US",
		4034880719,
		"http://www.halvorson.com/",
		firstBusinessTime,
	}

	lastBusinessTime, _ := time.Parse("2006-01-02 15:04:05", "2014-02-24 00:03:02")
	business49999 := Business{
		49999,
		"4be0c534-4d8e-4654-a92b-e833fe62ab8b",
		"Kuhic, Bayer and Johns",
		"48933 Cormier Islands",
		"",
		"Lake Margarettside",
		"NV",
		85524,
		"US",
		5270180703,
		"http://www.gleason.com/",
		lastBusinessTime,
	}

	middleBusinessTime, _ := time.Parse("2006-01-02 15:04:05", "2012-07-28 04:00:09")
	business40988 := Business{

		40988,
		"ab3a7b23-a6d5-48a1-9603-3cb51ff7e000",
		"Lebsack and Sons",
		"85372 Prosacco Stream Suite 540",
		"",
		"West Vertieport",
		"LA",
		76144,
		"US",
		8419926954,
		"http://www.ortiz.com/",
		middleBusinessTime,
	}

	if !reflect.DeepEqual(responseBusiness0, business0) {
		t.Error("business with id 0 as returned by API is incorrect")
	}
	if !reflect.DeepEqual(responseBusiness40988, business40988) {
		t.Error("business with id 40988 as returned by API is incorrect")
	}
	if !reflect.DeepEqual(responseBusiness49999, business49999) {
		t.Error("business with id 49999 as returned by API is incorrect")
	}
}
