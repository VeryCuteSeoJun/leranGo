package main

import (
	"fmt"
	"time"
)

func main() {
	go SayMessage("SeoJun") // go 키워드를 이용해서 이런식으로 병렬적으로 프로그램을 실행 가능함
	SayMessage("PPAP")      // 하지만 go 키워드는 main함수가 죽으면 자동으로 종료됨
}

func SayMessage(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello " + name)
		time.Sleep(time.Second)
	}
}
