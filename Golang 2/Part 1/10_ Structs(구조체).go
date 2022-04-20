package main

import "fmt"

type person struct { // person == 구조체의 이름
	name    string
	age     int
	favFood []string
}

func main() {
	//구조체에 값을 넣는법

	favFood := []string{"pizza", "ramen"}

	// 순서대로 넣기
	a := person{"SeoJun", 15, favFood}

	// 이름으로 넣기
	b := person{name: "SeoJun", age: 15, favFood: favFood}

	fmt.Println(a)
	fmt.Println(b)
}
