package main

import "fmt"

func main() {
	const name = "Hello!"
	// 이렇게 만든 애는 상수임 변경 X

	var var1 string = "Hello!"
	fmt.Printf(var1)
	// 기본형, 타입을 써줘야
	// 함수 밖에서도 사용 가능

	var2 := "Hello!"
	fmt.Printf(var2)
	// 축약형, 애는 타입을 안써도 지가 찿아줌
	// 함수 안에서만 사용 가능
}
