package main

import "fmt"

func main() {
	m := make(map[int]string)
	m[1] = "US"
	m[81] = "Japan"
	m[86] = "China"

	fmt.Println(m)

	// makeを使わずリテラルで定義した場合
	m = map[int]string{1: "Taro", 2: "Hanako", 3: "Jiro"}
	fmt.Println(m)

	m = map[int]string{
		1: "Taro",
		2: "Hanako",
		3: "Jiro", // カンマが必要
	}
	fmt.Println(m)

	m = map[int]string{1: "A", 2: "B", 3: "C"}
	fmt.Println(m)
	s, ok := m[1]
	fmt.Printf("%s, %v\n", s, ok) // A, true
	s, ok = m[9]
	fmt.Printf("%s, %v\n", s, ok) // , false
	_, ok = m[3]
	fmt.Printf("%v\n", ok) // true

	if _, ok = m[1]; ok { // m[1]の結果が取得できたらOKはtrue
		fmt.Println("OK")
	}
	if _, ok = m[9]; ok { // m[9]の結果が取得できない場合はOKはfalse
		fmt.Println("OK")
	} else {
		fmt.Println("NG")
	}

	// for文でkey, valueを出力
	for k, v := range m {
		fmt.Printf("%d, %s\n", k, v)
	}

	// delete
	delete(m, 2) // key=2を削除
	for k, v := range m {
		fmt.Printf("%d, %s\n", k, v)
	}
}
