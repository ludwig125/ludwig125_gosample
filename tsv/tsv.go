package main

import (
	"encoding/csv"
	"fmt"
	"io"
)

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

		if len(record) != 2 {
			return nil, fmt.Errorf("tsv data does not have exactly two fields. data: %v", record)
		}
		res = append(res, Data{A: record[0], B: record[1]})
	}
	return res, nil
}

// func main() {
// 	// テスト用文字列
// 	str := "test\tテスト\nHello\tこんにちは"

// 	// CSVのReaderを用意
// 	r := csv.NewReader(strings.NewReader(str))

// 	// デリミタ(TSVなら\t, CSVなら,)設定
// 	r.Comma = '\t'

// 	// コメント設定(なんとコメント文字を指定できる!)
// 	r.Comment = '#'

// 	// 全部読みだす
// 	records, err := r.ReadAll()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// 各行でループ
// 	for _, v := range records {
// 		// 1列目
// 		fmt.Print(v[0])

// 		fmt.Print(" | ")

// 		// 2列目
// 		fmt.Println(v[1])
// 	}
// }
