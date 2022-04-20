package main

import "fmt"

func main() {
	fmt.Println("카톡 좆같게 옾정")
}

func IF1(number int) bool { // IF만 사용시
	if number == 1 {
		fmt.Println("숫자는 1")
		return true
	}
	return false
}

func IF2(number int) bool { // IF와 else if 사용시
	if number == 1 {
		fmt.Println("숫자는 1")
		return true
	} else if number == 2 {
		fmt.Println("숫자는 2")
		return true
	}
	return false
}

func IF3(number int) bool { // IF와 else if, else 사용시
	if number == 1 {
		fmt.Println("숫자는 1")
		return true
	} else if number == 2 {
		fmt.Println("숫자는 2")
		return true
	} else {
		fmt.Println("숫자는 3")
		return false
	}
}
