package main

import "fmt"

func sub() {
	for {
		fmt.Println("sub loop")
	}
}

func main() {
	go sub() // ゴルーチン開始 sub loopとmain loopが交互に出力される
	for {
		fmt.Println("main loop")
	}
}
