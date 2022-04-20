package main

import "fmt"

func main() {
	// 맵 타입 선언법 : map[keytype]valuetype{}
	Maps := map[string]string{"name": "SeoJun", "age": "15"}

	// 맵을 반복에 사용하기
	for key, value := range Maps {
		fmt.Println(key, value)
	}
}
