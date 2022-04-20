package accounts

import "strconv"

// 메소드를 만드는 법

// 계정 구조체
type Account struct {
	owner   string
	balance int
}

// 게정에 돈에 x만큼 더하기 (func과 사이와 함수의 이름 사이에 (a <구조체명>) 이런식으로 리시버를 달아 메소드를 만들 수 있음. 여기서 a는 파이선에서 self의 역활을 함)
func (a Account) Deposit(amount int) { // 하지만 이건 구조체의 복사본을 받아 건들이는거기 때문에 구조체의 값이 변경되지 않음
	a.balance += amount
}

// balance값을 리턴하는 함수
func (a Account) Balance() int {
	return a.balance
}

// Acount 구조체를 print했을때 나오는 메시지를 수정 가능 (&{SeoJun 0} -> name : SeoJun, money : 0)
func (a Account) String() string { // 메소드 명이 String이기 때문에 가능!
	return "name : " + a.owner + ", money : " + strconv.Itoa(a.balance)
}
