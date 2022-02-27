package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Create an HTTP handler
	http.HandleFunc("/github", func(w http.ResponseWriter, r *http.Request) {
		// Retrieve GitHub's API status
		status, err := githubStatus()
		if err != nil {
			// Send an internal error response if we were unable to retrieve the status.
			// note: not a great idea to raw errors to a client, but this is just a demonstration.
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Log the status so we can see what's happening from the server's perspective.
		log.Printf("/github handler requst: status %q", status)

		// Write the response to the client.
		fmt.Fprintf(w, "GitHub Status: %q", status)
	})

	http.ListenAndServe("127.0.0.1:8080", nil)
}

// githubStatus retrieves GitHub's API status
func githubStatus() (string, error) {
	// Log the start and end of the function so we can see how many times it's called.
	log.Println("Making request to GitHub API")
	defer log.Println("Request to GitHub API Complete") // The defer causes this to be logged after the function's return statement.

	// // Atrificially delay this function to emulate a long running task
	// time.Sleep(1 * time.Second)

	// Make a request to the GitHub Status API
	// resp, err := http.Get("https://status.github.com/api/status.json")
	resp, err := http.Get("http://localhost:9000/status")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check that we got a good response.
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("github response: %s", resp.Status)
	}

	// Anonymous struct to extract the status from the response
	var r Status

	// Decode the JSON response
	err = json.NewDecoder(resp.Body).Decode(&r)

	return r.Status, err
}

// Status is server status.
type Status struct {
	Status string `json:"status"`
	ReqCnt int    `json:"count,omitempty"`
}
