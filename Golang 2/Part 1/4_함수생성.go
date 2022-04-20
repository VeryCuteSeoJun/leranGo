package main

import (
	"fmt"
	"strings"
)

func main() {
	a1, b1 := lenAndupper("SeoJun") //이런식으로 리턴값을 받음 만약 리턴값을 하나를 생략하고 싶다면
	a2, _ := lenAndupper("SeoJun")  //이런식으로 생략하고 싶은 변수명을 _로 대체 하면 됨
	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(a2)
	arryinput(1, 2, 3, 4)

}

func multiply(a int, b int) int { //입력 받는 변수의 타입, 리턴값의 타입을 적어야함 만약 입력이나 리턴이 없으면 공백
	return a * b
}

func lenAndupper(name string) (int, string) { //리턴하는 값이 여러개(배열)일 경우 리턴하는 값의 타입을 () 안에 넣어 여려개로
	return len(name), strings.ToUpper(name) //이런식으로 ,로 구분해서 여려개(배열)로 리턴
}

func arryinput(numbers ...int) { //이런식으로 받는 변수명과 타입 사이에 ...을 넣어 배열값을 변수로 받을 수 있음
	fmt.Println(numbers)
}

func lenAndupper2(name string) (namelen int, upname string) { // 리턴하는 값의 타입 앞에 리턴하는 변수명을 넣어서 리턴하는 변수를 함수를 정의할때 같이 정의 가능
	defer fmt.Println("끗") //앞에 defer를 붙혀서 이 함수가 끝나면 이 명령이 실행되게 설정이 가능함

	namelen = len(name)
	upname = strings.ToUpper(name)
	return
}
