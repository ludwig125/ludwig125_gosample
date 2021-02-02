package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/** JSONデコード用に構造体定義 */
type person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	jsonData := `
	[
	  	{"id":1,"name":"a"},
  		{"id":2,"name":"b"},
 	 	{"id":3,"name":"c"},
 	 	{"id":4,"name":"d"}
	]
	`
	jsonBytes := ([]byte)(jsonData)

	// JSONデコード
	var people []person
	if err := json.Unmarshal(jsonBytes, &people); err != nil {
		log.Fatal(err)
	}
	// デコードしたデータを表示
	for _, p := range people {
		fmt.Printf("%d : %s\n", p.ID, p.Name)
	}
}
