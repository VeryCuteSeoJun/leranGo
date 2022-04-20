package main

import "fmt"

func main() {
	arry1 := [5]string{"a", "b", "c", "d", "e"} //길이가 5 이고 타입이 string인 배열을 생성
	arry2 := []string{"a", "b", "c", "d", "e"}  //길이 제한이 없고 타입이 string인 배열을 생성
	arry2 = append(arry2, "f")                  //배열에 "f"를 추가함
	fmt.Println(arry1)
	fmt.Println(arry2)
}
