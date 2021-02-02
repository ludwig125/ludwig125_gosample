package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

/** JSONデコード用に構造体定義 */
type documents struct {
	ID     string `json:"id"`
	Fields fields `json:"fields"`
}

type fields struct {
	Account int    `json:"account"`
	Item    string `json:"item"`
	Local   int    `json:"local"`
}

const dumpFile = "testFile"

func main() {
	if err := writeFile("http://localhost:8080"); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	if err := readFile(); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	// name, err := fetchRequiredName("c")
	// fmt.Println(name, err)
}
func writeFile(url string) error {
	//url := "http://localhost:8080"
	//url := "http://localhost:8080/json?type=documents"
	//url := "http://testhost:8080"
	resp, err := getBody(url)
	if err != nil {
		return fmt.Errorf("failed to getBody. url: %s, %v", url, err)
	}

	// ローカルのストレージに dump しておく
	f, err := os.Create(dumpFile)
	if err != nil {
		return fmt.Errorf("failed to create dump file: %v", err)
	}
	defer f.Close()

	if _, err := io.Copy(f, resp); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	log.Println("succeeded to write dump file")
	return nil
}

func readFile() error {
	f, err := os.Open(dumpFile)
	if err != nil {
		return fmt.Errorf("failed to open file. %v", err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		// line includes '\n'.
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		fmt.Println(strconv.Quote(line))
	}
	return nil
}

func jsonParse(line string) (*documents, error) {
	line = strings.TrimRight(line, "]")
	line = strings.TrimRight(line, ",")

	data := new(documents)

	if err := json.Unmarshal([]byte(line), data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return data, err
	}

	//fmt.Printf("data %T\n", data)
	// fmt.Println(data.ID)
	// fmt.Println(data.Fields)

	return data, nil
}

func getBody(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url + "/json?type=documents")
	if err != nil {
		return nil, fmt.Errorf("failed to http.Get, %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
		return nil, fmt.Errorf("response status code is not 200 but '%d', url: %s", resp.StatusCode, url)
	}

	return resp.Body, nil
}
