package main

import "fmt"

func main() {
	fmt.Println("start!")
}

func superadd1(numbers ...int) { //range는 arry(배열)형태를 반복문에서 쓸수 있게 해줌.
	for index := range numbers { //하지만 여기서는 배열의 인덱스를 index에 넣어줌 (인덱스만 받을 떄)
		fmt.Println(index)
	}
}

func superadd2(numbers ...int) { //하지만 저렇게 인덱스와 숫자를 받는 변수를 두개를 써줘야함
	for index, number := range numbers { //인덱스와 숫자를 둘다 받을떄
		fmt.Println(index, number)
	}
}

func superadd3(numbers ...int) { //하지만 저렇게 인덱스와 숫자를 받는 변수를 두개를 써줘야함
	for _, number := range numbers { //숫자만 받을 때
		fmt.Println(number)
	}
}

//이외 우리가 알던 와일 for 처럼 for를 사용 가능
