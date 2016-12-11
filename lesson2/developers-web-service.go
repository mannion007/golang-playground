package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

type Reponse struct {
	Code   int    `json:"code"`
	Result string `json:"result"`
}

type developer struct {
	name       string
	manager    string
	department string
	team       string
	floor      int
}

var developers = map[int]developer{
	1: developer{
		name:       "James",
		manager:    "Rudolf",
		department: "Development",
		team:       "DAS",
		floor:      5,
	},
	2: developer{
		name:       "Kamil",
		manager:    "Rudolf",
		department: "Development",
		team:       "DAS",
		floor:      5,
	},
	3: developer{
		name:       "Chris",
		manager:    "Charlie",
		department: "Development",
		team:       "Product",
		floor:      5,
	},
	4: developer{
		name:       "Andrew",
		manager:    "Charlie",
		department: "Development",
		team:       "Product",
		floor:      5,
	},
	5: developer{
		name:       "Gijs",
		manager:    "Frazer",
		department: "Project Management",
		team:       "Product",
		floor:      5,
	},
	6: developer{
		name:       "Sorin",
		manager:    "Edina",
		department: "Development",
		team:       "SmartShop",
		floor:      5,
	},
	7: developer{
		name:       "Bomany",
		manager:    "Simon",
		department: "Devops",
		team:       "SmartShop",
		floor:      5,
	},
	8: developer{
		name:       "Peter",
		manager:    "David",
		department: "Development",
		team:       "Web",
		floor:      4,
	},
}

func sendResponse(w http.ResponseWriter, code int, result string) {
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
	//result := "Name cannot be empty"

	theResult, _ := json.Marshal(developers)

	//fmt.Println(string(theResult))

	sendResponse(w, code, string(theResult))
}

func main() {
	var port = flag.String("port", "8000", "The port the webserver runs on")
	http.HandleFunc("/", hello)
	http.ListenAndServe(fmt.Sprintf(":%s", *port), nil)
}
