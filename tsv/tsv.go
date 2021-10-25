package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
	"time"
)

// https://deeeet.com/writing/2016/10/25/go-interface-testing/

func main() {
	db := NewDB()
	server := NewServer()
	if err := run(db, time.Now(), server); err != nil {
		log.Fatal(err)
	}
}

func run(db DB, now time.Time, srv Server) error {
	tables := []string{"TableA", "TableB", "TableC"}
	for _, table := range tables {
		partitions, err := db.ReadPartition(table)
		if err != nil {
			return err
		}
		// log.Println(table, "partition is", partitions)
		if !containsToday(partitions, now) {
			return fmt.Errorf("table %s does not have today's partition. partitions: %s", table, partitions)
		}
	}

	res, err := db.GetTSV()
	if err != nil {
		return err
	}

	out, err := ConvertDataFromTSV(strings.NewReader(res))
	if err != nil {
		return err
	}
	fmt.Println(out)

	if err := srv.Delete("tmp_params"); err != nil {
		return err
	}

	currentData, err := srv.Read("tmp_params")
	if err != nil {
		return err
	}
	fmt.Println(currentData)

	return nil
}

type DB interface {
	ReadPartition(string) (string, error)
	GetTSV() (string, error)
}

type database struct {
}

func NewDB() DB {
	return database{}
}

func (d database) ReadPartition(table string) (string, error) {
	return "", nil
}

func (d database) GetTSV() (string, error) {
	return "", nil
}

func containsToday(target string, now time.Time) bool {
	format := "20060102"
	return strings.Contains(target, now.Format(format))
}

type Data struct {
	A string
	B string
}

// ref: https://ankurraina.medium.com/reading-a-simple-csv-in-go-36d7a269cecd
// https://text.baldanders.info/golang/encode-csv-or-tsv-data/

func ConvertDataFromTSV(in io.Reader) ([]Data, error) {
	// Parse the file
	r := csv.NewReader(in)
	//r := csv.NewReader(bufio.NewReader(csvfile))

	// デリミタ設定(TSVなら\t, CSVなら,)
	r.Comma = '\t'

	firstLine := true
	var res []Data
	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			// 残りのデータが空だったらここに入る
			// 最初からデータが空の時もここに入る
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error occured in read tsv: %v", err)
		}
		if firstLine { // １行目は項目名なので捨てる
			firstLine = false
			continue
		}

		if len(record) != 2 {
			return nil, fmt.Errorf("tsv data does not have exactly two fields. data: %v", record)
		}
		res = append(res, Data{A: record[0], B: record[1]})
	}
	return res, nil
}

type Server interface {
	Delete(string) error
	Read(string) (string, error)
}

type server struct {
}

func NewServer() Server {
	return server{}
}

func (s server) Delete(params string) error {
	return nil
}

func (s server) Read(params string) (string, error) {
	return "", nil
}
