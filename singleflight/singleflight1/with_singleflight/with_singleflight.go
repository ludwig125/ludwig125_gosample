package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/sync/singleflight"
)

// ref https://medium.com/@vCabbage/go-avoid-duplicate-requests-with-sync-singleflight-311601b3068b
// 事前に別のターミナルで go run mock_server/server.goを実行しておく必要がある

func main() {
	// We need a Group to use singleflight.
	var requestGroup singleflight.Group

	http.HandleFunc("/github", func(w http.ResponseWriter, r *http.Request) {
		// This time we'll wrap the githubStatus() call with singleflight's Group.Do()
		// Do takes a key (more on this later) and a function that returns a interface{} and an error.
		v, err, shared := requestGroup.Do("github", func() (interface{}, error) {
			// githubStatus() returns Status, error, which statifies interface{}, error, so we can return the result directly.
			return githubStatus()
		})
		// Do returns an interface{}, error, and a bool which indicates whether multiple calls to the function shared the same result.

		// Check the error, as before.
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// We know that v will be a string, so we'll use a type assertion.
		status := v.(Status)

		// Update the log statement so we can see if the results were shared.
		log.Printf("/github handler requst: status %q, shared result %t", status.Status, shared)

		fmt.Fprintf(w, "GitHub Status: %q, Request Count: %d", status.Status, status.ReqCnt)
	})

	http.ListenAndServe("127.0.0.1:8080", nil)
}

// githubStatus retrieves GitHub's API status
func githubStatus() (Status, error) {
	// No changes made to this function other than removing the comments for brevity.
	log.Println("Making request to GitHub API")
	defer log.Println("Request to GitHub API Complete")

	// // Atrificially delay this function to emulate a long running task
	// time.Sleep(1 * time.Second)

	// Make a request to the GitHub Status API
	resp, err := http.Get("http://localhost:9000/status")
	if err != nil {
		return Status{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return Status{}, fmt.Errorf("github response: %s", resp.Status)
	}

	//r := struct{ Status string }{}
	var r Status

	// Decode the JSON response
	err = json.NewDecoder(resp.Body).Decode(&r)

	return r, err
}

// Status is server status.
type Status struct {
	Status string `json:"status"`
	ReqCnt int    `json:"count,omitempty"`
}
