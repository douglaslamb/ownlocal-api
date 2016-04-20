package main

import "time"

//global business list
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

type Links struct {
	First string `json:"first"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
	Last  string `json:"last"`
}

type BusinessListResponse struct {
	BusinessList []Business `json:"businesses"`
	Links        `json:"links"`
}
