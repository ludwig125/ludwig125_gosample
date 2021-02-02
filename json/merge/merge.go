package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// JSON EncodeしてDecodeすると中身が微妙に変わる

// Sample is struct for json parse.
type Sample struct {
	Price    float64        `json:",string"`
	Contents map[string]int `json:"contents,omitempty"`
}

// // Contents is struct.
// type Contents struct {
// 	Name map[string]int `json:"name"`
// }

func main() {
	s := `{"price":"1.0"}`
	fmt.Printf("before %+v\n\n", s)
	fn1(s)
	fn2(s)
}

func fn1(s string) {
	var x Sample
	err := json.Unmarshal([]byte(s), &x)
	if err != nil {
		log.Fatalf("%+v err: %v\n", x, err)
	}
	fmt.Printf("after Unmarshal %+v\n", x)

	b, err := json.Marshal(&x)
	if err != nil {
		log.Fatalf("%+v err: %v\n", x, err)
	}
	fmt.Printf("after Marshal %+v\n", string(b))
}

func fn2(s string) {
	var x Sample
	err := json.Unmarshal([]byte(s), &x)
	if err != nil {
		log.Fatalf("%+v err: %v\n", x, err)
	}
	fmt.Printf("after Unmarshal %+v\n", x)

	m := map[string]int{
		"key1": 10,
		"key2": 20,
		"key3": 300,
		"key4": 400,
	}
	x.Contents = m

	b, err := json.Marshal(&x)
	if err != nil {
		log.Fatalf("%+v err: %v\n", x, err)
	}
	fmt.Printf("after Marshal %+v\n", string(b))
}
