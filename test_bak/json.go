package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

/** JSONデコード用に構造体定義 */
type personalName struct {
	Name string `json:"name"`
}

func main() {
	//	name, err := fetchRequiredName("c")
	name, err := fetchRequiredName()
	fmt.Println(name, err)
}

//func fetchRequiredName(s string) (string, error) {
func fetchRequiredName() (string, error) {
	s := "c"

	url := "http://localhost:8080/json"
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to http.Get, %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("response status code is not 200 but '%d'", resp.StatusCode)
	}

	// JSONデコード
	var names []personalName
	// if err := json.NewDecoder(resp.Body).Decode(&names); err != nil {
	// 	return "", fmt.Errorf("failed to JSON decode, %v %v", err, resp.Body)
	// }

	decoder := json.NewDecoder(resp.Body)
	for {
		if err := decoder.Decode(&names); err == io.EOF {
			break
		} else if err != nil {
			return "", fmt.Errorf("failed to JSON decode, %v %v", err, resp.Body)
		}
	}

	// デコードしたデータを表示
	for _, n := range names {
		if strings.Contains(n.Name, s) {
			return n.Name, nil
		}
	}
	return "", fmt.Errorf("failed to fetch. not match %s", s)
}
