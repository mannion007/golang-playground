package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

type Reponse struct {
	Code   int    `json:"code"`
	Result interface{} `json:"result"`
}

type developer struct {
	Name       string
	Manager    string
	Department string
	Team       string
	Floor      int
}

var developers = map[int]developer{
	1: developer{
		Name:       "James",
		Manager:    "Rudolf",
		Department: "Development",
		Team:       "DAS",
		Floor:      5,
	},
	2: developer{
		Name:       "Kamil",
		Manager:    "Rudolf",
		Department: "Development",
		Team:       "DAS",
		Floor:      5,
	},
	3: developer{
		Name:       "Chris",
		Manager:    "Charlie",
		Department: "Development",
		Team:       "Product",
		Floor:      5,
	},
	4: developer{
		Name:       "Andrew",
		Manager:    "Charlie",
		Department: "Development",
		Team:       "Product",
		Floor:      5,
	},
	5: developer{
		Name:       "Gijs",
		Manager:    "Frazer",
		Department: "Project Management",
		Team:       "Product",
		Floor:      5,
	},
	6: developer{
		Name:       "Sorin",
		Manager:    "Edina",
		Department: "Development",
		Team:       "SmartShop",
		Floor:      5,
	},
	7: developer{
		Name:       "Bomany",
		Manager:    "Simon",
		Department: "Devops",
		Team:       "SmartShop",
		Floor:      5,
	},
	8: developer{
		Name:       "Peter",
		Manager:    "David",
		Department: "Development",
		Team:       "Web",
		Floor:      4,
	},
}

func sendResponse(w http.ResponseWriter, code int, result interface{}) {
	response := Reponse{
		Code:   code,
		Result: result,
	}

	json, err := json.Marshal(response)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func hello(w http.ResponseWriter, r *http.Request) {
	code := 200
	sendResponse(w, code, developers)
}

func main() {
	var port = flag.String("port", "8000", "The port the webserver runs on")
	http.HandleFunc("/", hello)
	http.ListenAndServe(fmt.Sprintf(":%s", *port), nil)
}
