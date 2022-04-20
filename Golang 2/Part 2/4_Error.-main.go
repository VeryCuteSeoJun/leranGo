package main

import (
	"fmt"
)

func main() {
	account := accounts.NewAccount("SeoJun") //지갑 구조체를 생성
	account.Deposit(10)                      // 10만큼 증가
	fmt.Println(account.Balance())           // 지갑의 잔액을 출력
	err := account.Withdraw(20)              // 20만큼 차감 (하지만 지갑의 잔고는 10이므로 에러 발생)
	if err != nil {                          // 에러가 nil이 아니면 == 에러가 발생했다면
		fmt.Println(err) // 에러를 출력
	}
	fmt.Println(account.Balance())
}
