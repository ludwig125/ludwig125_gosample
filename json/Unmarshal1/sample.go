package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// ref https://text.baldanders.info/golang/operate-json/

// Booksdata is struct.
type Booksdata struct {
	Books []BooksdataBook `json:"books"`
}

// BooksdataBook is struct.
type BooksdataBook struct {
	Creators []BooksdataBookCreator `json:"Creators"`
	ID       string                 `json:"ID"`
	Type     string                 `json:"Type"`
}

// BooksdataBookCreator is struct.
type BooksdataBookCreator struct {
	Name  string  `json:"Name"`
	Role  string  `json:"Role"`
	Price float64 `json:"Price"`
}

func main() {
	sampleJSON := `{
		"books": [
				{
				"Type": "hoge",
				"ID": "56839",
				"Creators": [
					{
					"Name": "AAAA",
					"Role": "翻訳",
					"Price": 1.123456789123456789123456789123456789123456789
					}
				]
				},
				{
				"Type": "fuga",
				"ID": "56840",
				"Creators": [
					{
					"Name": "BBBB",
					"Role": "翻訳",
					"Price": 1.123456789123456789123456789123456789123456789
					}
				]
				}
			]
		}`

	fn1(sampleJSON)
	fmt.Println()
	fn2(sampleJSON)
}

func fn1(s string) {
	var x Booksdata
	err := json.Unmarshal([]byte(s), &x)
	if err != nil {
		log.Fatalf("%+v err: %v\n", x, err)
	}
	fmt.Printf("after Unmarshal %+v\n", x)

	fmt.Printf("Creators %+v\n", x.Books[0].Creators)
	fmt.Printf("Price %+v\n", x.Books[0].Creators[0].Price)
}

func fn2(s string) {
	var x map[string]interface{}
	err := json.Unmarshal([]byte(s), &x)
	if err != nil {
		log.Fatalf("%+v err: %v\n", x, err)
	}
	fmt.Printf("after Unmarshal %+v\n", x)

	book := x["books"].([]interface{})
	creators := book[0].(map[string]interface{})["Creators"].([]interface{})
	price := creators[0].(map[string]interface{})["Price"].(float64)

	fmt.Printf("Creators %+v\n", creators...)
	fmt.Printf("Price %+v\n", price)
}
