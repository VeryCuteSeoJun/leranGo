package main

import (
	"fmt"
	"strconv"
)

// 계정 구조체
type Account struct {
	owner   string
	balance int
}

// 직접적으로 계정을 만드는것보단 구조체를 만들어서 리턴하는 함수를 통해 구조체를 만듬
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Acount 구조체를 print했을때 나오는 메시지를 수정 가능
func (a Account) String() string {
	return "name : " + a.owner + ", money : " + strconv.Itoa(a.balance)
}

func main() {
	a := NewAccount("SeoJun")
	fmt.Println(a)
}
