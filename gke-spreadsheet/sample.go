package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

var (
	sheetID            = ""
	credentialFileDir  = "credential"
	credentialFilePath = ""
)

type stockInfo struct {
	Code   string
	Name   string
	Status string
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("%s environment variable not set.", k)
	}
	log.Printf("%s environment variable set.", k)
	//log.Println("this is env", k, ":", v)
	return v
}

func fileMustExists(name string) string {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		log.Fatalf("file '%s' does not exists", name)
	}
	return name
}

func init() {
	sheetID = mustGetenv("TRADEDERBY_SHEETID")

	// testまたはdebug用
	cdir := os.Getenv("CREDENTIALFILE_DIR")
	if cdir != "" {
		credentialFileDir = cdir
	}
	// Sheetをみにいくためのserviceaccountの置き場所
	// ローカルでもGKEでも同じになるようにvolumeのmountPathなどを調節している
	credentialFilePath = fileMustExists(credentialFileDir + "/gke-trade-derby-serviceaccount.json")

	log.Println("set env infos")
}

func main() {
	stockInfos := []stockInfo{
		stockInfo{"1417", "ミライトHD", "信用売"},
		stockInfo{"6088", "シグマクシス", "現物買"},
		stockInfo{"6367", "ダイキン", "現物買"},
		stockInfo{"9759", "NSD", "信用売"},
	}

	var sIfs [][]interface{}
	for _, s := range stockInfos {
		var sIf []interface{}
		sIf = append(sIf, s.Code)
		sIf = append(sIf, s.Name)
		sIf = append(sIf, s.Status)
		sIfs = append(sIfs, sIf)
	}
	log.Println(sIfs)

	// spreadsheetのclientを取得
	srv, err := getSheetClient()
	if err != nil {
		log.Fatalf("failed to get sheet client. err: %v", err)
	}
	log.Println("succeeded to get sheet client")

	log.Println("trying to write sheet")
	if err := clearAndWriteSheet(srv, sheetID, "trade-derby", sIfs); err != nil {
		log.Fatalf("failed to clearAndWriteSheet. %v", err)
	}
	log.Println("succeeded to write sheet")
}

// spreadsheets clientを取得
func getSheetClient() (*sheets.Service, error) {
	// googleAPIへのclientをリクエストから作成
	client, err := getClientWithJSON()
	if err != nil {
		return nil, fmt.Errorf("failed to getClientWithJSON: %v", err)
	}
	// spreadsheets clientを取得
	srv, err := sheets.New(client)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve Sheets Client: %v", err)
	}
	return srv, nil
}

func getClientWithJSON() (*http.Client, error) {
	data, err := ioutil.ReadFile(credentialFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read client secret file. path: '%s', %v", credentialFilePath, err)
	}
	conf, err := google.JWTConfigFromJSON(data, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		return nil, fmt.Errorf("failed to parse client secret file to config: %v", err)
	}
	return conf.Client(context.Background()), nil
}

func clearSheet(srv *sheets.Service, sid string, sname string) error {
	// clear stockprice rate spreadsheet:
	resp, err := srv.Spreadsheets.Values.Clear(sid, sname, &sheets.ClearValuesRequest{}).Do()
	if err != nil {
		return fmt.Errorf("failed to clear value: %v", err)
	}
	status := resp.ServerResponse.HTTPStatusCode
	if status != 200 {
		return fmt.Errorf("HTTPstatus error. %v", status)
	}
	return nil
}

// sheetのID, sheet名と対象のデータ（[][]interface{}型）を入力値にとり、
// Sheetにデータを記入する関数
func writeSheet(srv *sheets.Service, sid string, sname string, records [][]interface{}) error {
	valueRange := &sheets.ValueRange{
		MajorDimension: "ROWS",
		Values:         records,
	}
	// Write stockprice rate spreadsheet:
	resp, err := srv.Spreadsheets.Values.Append(sid, sname, valueRange).ValueInputOption("USER_ENTERED").InsertDataOption("INSERT_ROWS").Do()
	if err != nil {
		return fmt.Errorf("failed to write value. %v", err)
	}
	status := resp.ServerResponse.HTTPStatusCode
	if status != 200 {
		return fmt.Errorf("HTTPstatus error. %v", status)
	}
	return nil
}

// SheetのClearとWriteを行う関数
func clearAndWriteSheet(srv *sheets.Service, sid string, sname string, records [][]interface{}) error {
	if err := clearSheet(srv, sid, sname); err != nil {
		return fmt.Errorf("failed to clearSheet. sheetID: %s, sheetName: %s, %v", sid, sname, err)
	}

	// writeSheetに渡す
	if err := writeSheet(srv, sid, sname, records); err != nil {
		return fmt.Errorf("failed to writeSheet. sheetID: %s, sheetName: %s, error data: [%v], %v", sid, sname, records, err)
	}
	return nil
}
