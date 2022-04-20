package accounts

import "errors"

// 계정 구조체
type Account struct {
	owner   string
	balance int
}

// 지갑에서 x만큼 돈을 차감하기
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount { //만약 에러가 발생하면
		return errors.New("지갑에 돈이 없습니다!") // 에러메시지를 리턴
	}
	a.balance -= amount
	return nil // 위에서 error타입을 리턴한다고 했으니 아무것도 리턴을 안할순 없음 따라서 nil(파이선값은곳에선 null)을 리턴해서 에러가 없음을 나타냄
}
