package main

import "fmt"

func main() {
	input3()
}

func input1() {
	var a string   //사용할 변수 생성
	fmt.Scanln(&a) //입력받아서 변수에 저장
	fmt.Println(a) // 출력
}

func input2() {
	var a string //사용할 변수 생성
	var b string //사용할 변수 생성

	fmt.Scanln(&a, &b) //입력받아서 변수에 저장 (입력시 뛰어쓰기로 구분)
	fmt.Println(a, b)  // 출력
}

func input3() {
	var n1, n2, n3, n4 int                                  //변수 생성
	fmt.Scanf("%d.%d.%d.%d", &n1, &n2, &n3, &n4)            //특정 포멧으로 입력받기(이 경우에는 .을 구분으로 입력)
	fmt.Printf("입력한 IPv4주소는 %d.%d.%d.%d\n", n1, n2, n3, n4) //포멧으로 출력
	// %d: 정수
	// %f: 소수(float)
	// %s: 문자열
	// %t: 불(bool)
	// %v: 배열
	// %v: 맵(map)
	// %v: 구족체
	// %p: 포인터
	// %v: 인터페이스
}
