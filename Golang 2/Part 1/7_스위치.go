package main

func main() {

}

func switch1(numbers int) bool { //스위치를 시작할떄 변수명을 작성할 경우
	switch numbers {
	case 1: //numbers가 1이라면
		return true
	case 2: //numbers가 2이라면
		return true
	case 3: //numbers가 3이라면
		return true
	}
	return false //아무것도 해당이 안되면 false를 리턴
}

func switch2(numbers int) bool { //스위치 안에서 변수명을 쓸 경우
	switch {
	case numbers == 1: //numbers가 1이라면
		return true
	case numbers == 2: //numbers가 2이라면
		return true
	case numbers == 3: //numbers가 3이라면
		return true
	default: //아무것도 해당 안되면
		return false
	}
}
