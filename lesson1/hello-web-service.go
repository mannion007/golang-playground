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
	code := 400
	result := "Name cannot be empty"
	name := r.FormValue("name")

	if len(name) != 0 {
		code = 200
		result = fmt.Sprintf("%s %s", "hello", r.FormValue("name"))
	}

	sendResponse(w, code, result)
}

func main() {
	var port = flag.String("port", "8000", "The port the webserver runs on")
	http.HandleFunc("/", hello)
	http.ListenAndServe(fmt.Sprintf(":%s", *port), nil)
}
