package main

import (
	"encoding/json"
	"fmt"

	"github.com/ludwig125/ludwig125_gosample/easyjson/mypackage"
	"github.com/mailru/easyjson"
)


func main() {

    var data mypackage.JSONData
    jsonBlob := `{"Data" : ["One", "Two", "Three"]}`

    err := json.Unmarshal([]byte(jsonBlob), &data)

    if err != nil {
        panic(err)
    }

    fmt.Println(data.Data)


	d := mypackage.JSONData{Data: []string{
		"a",
		"b",
	}}
	dat, err := easyjson.Marshal(d)
	if err != nil {
		panic(err)
	}
	fmt.Println("dat",string(dat))
	v := mypackage.JSONData{}
	if err := easyjson.Unmarshal(dat, &v);err != nil {
		panic(err)
	}
	// fmt.Println(string(b)) // {"B":{"c":[1,2,3],"d":null}}
	fmt.Println(v.Data)


	dat2:=[]byte(`{"Data":["a2","b2"]}`)
	v2 := mypackage.JSONData{}
	if err := easyjson.Unmarshal(dat2, &v2);err != nil {
		panic(err)
	}
	// fmt.Println(string(b)) // {"B":{"c":[1,2,3],"d":null}}
	fmt.Println(v2.Data)
}
