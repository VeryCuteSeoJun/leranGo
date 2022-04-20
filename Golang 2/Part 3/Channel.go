package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)       // 이런식으로 채널을 만듬
	go SayMessage("SeoJun", c) // go 키워드를 이용해서 이런식으로 병렬적으로 프로그램을 실행 가능함
	go SayMessage("PPAP", c)   // 하지만 go 키워드는 main함수가 죽으면 자동으로 종료됨
	fmt.Println(<-c)           // 하지만 채널이라는 파이프를 통해 go키워드가 다 실행될때까지 main함수가 죽지 않도록 기다리고 go 키워드로 실행된 함수의 리턴값을 받을 수 있음
	fmt.Println(<-c)           // 채널로 정보를 받음
}

func SayMessage(name string, c chan bool) { // bool값을 주고 받는 채널을 매개변수로 받음
	for i := 0; i < 10; i++ {
		fmt.Println("Hello " + name)
		time.Sleep(time.Second)
	}
	c <- true // 채널로 정보를 보냄
}
