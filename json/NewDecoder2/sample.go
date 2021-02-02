package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Product is struct for json parse.
type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price,string"`
}

func main() {
	s := `{"name":"Galaxy Nexus","price":"3460.00"}`
	var pro Product
	err := json.NewDecoder(strings.NewReader(s)).Decode(&pro)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(pro)
}
