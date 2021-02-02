package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	token := mustGetenv("CIRCLE_API_USER_TOKEN")
	defer func() {
		err := requestCircleci(token, "delete_gke_cluster")
		//err := requestCircleci(token, "list_gke_cluster")
		if err != nil {
			log.Fatalf("failed to requestCircleci: %v", err)
		}
		log.Println("requestCircleci successfully")
	}()
	log.Println("do task")
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("%s environment variable not set.", k)
	}
	log.Printf("%s environment variable set.", k)
	return v
}

func requestCircleci(token, job string) error {
	// 参考
	// https://circleci.com/docs/ja/2.0/api-job-trigger/
	// https://circleci.com/docs/api/#trigger-a-new-job
	client := &http.Client{}
	circleciURL := "https://circleci.com/api/v1.1/project/github/ludwig125/gke-test/tree/master"
	j := fmt.Sprintf(`{"build_parameters": {"CIRCLE_JOB": "%s"}}`, job)
	req, err := http.NewRequest("POST", circleciURL, bytes.NewBuffer([]byte(j)))
	req.SetBasicAuth(token, "")
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// circleci APIを呼び出すと、201 Created が返ってくるのでチェック
	if resp.StatusCode != 201 {
		return fmt.Errorf("status code Error. %v", resp.StatusCode)
	}

	// レスポンス本文が見たい場合はここのコメントアウトを外す
	// body, err := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))
	return nil
}

// --上の書き方は以下と同じ--

// BuildParams is params for circleci API
type BuildParams struct {
	CircleciJobs CircleciJob `json:"build_parameters"`
}

// CircleciJob designate circleci job name
type CircleciJob struct {
	JobName string `json:"CIRCLE_JOB"`
}

func requestCircleci2(token, job string) error {
	client := &http.Client{}

	params := BuildParams{
		CircleciJobs: CircleciJob{JobName: job},
	}
	jsonBytes, err := json.Marshal(params)
	if err != nil {
		return fmt.Errorf("failed to Marshal: %v", err)
	}
	circleciURL := "https://circleci.com/api/v1.1/project/github/ludwig125/gke-test/tree/master"
	req, err := http.NewRequest("POST", circleciURL, bytes.NewBuffer(jsonBytes))
	req.SetBasicAuth(token, "")
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		return fmt.Errorf("status code Error. %v", resp.StatusCode)
	}

	return nil
}
