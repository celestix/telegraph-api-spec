package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/anonyindian/telegraph-api-spec/telegraph"

	"github.com/anaskhan96/soup"
)

const API_URL = "https://telegra.ph/api"

func main() {
	resp, err := soup.Get(API_URL)
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	telegraph.FindMethodsAndTypes(doc)
	result := telegraph.MakeResult(doc)
	WriteFile("telegraph.json", MakeJSON(result))
}

/*
UTILITIES
*/

func MakeJSON(result telegraph.Result) []byte {
	b, err := json.MarshalIndent(&result, "", "	")
	if err != nil {
		panic(err.Error())
	}
	return b
}

func WriteFile(filename string, data []byte) (*os.File, error) {
	_, err := os.Open(filename)
	if err == nil {
		os.Remove(filename)
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(filename)
	return f, err
}
