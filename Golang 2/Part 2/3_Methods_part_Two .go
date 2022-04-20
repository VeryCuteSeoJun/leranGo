package accounts

// 실질적으로 (원본)구조체의 값을 건드는 방법.

// 계정 구조체
type Account struct {
	owner   string
	balance int
}

// 게정에 돈에 x만큼 더하기
func (a Account) DepositOld(amount int) { // 하지만 이건 구조체의 복사본을 받아 건들이는거기 때문에 구조체의 값이 변경되지 않음
	a.balance += amount
}

// 게정에 돈에 x만큼 더하기
func (a *Account) DepositNew(amount int) { // 하지만 이건 *를 달아 받은 구조체의 원본을 건들이기 때문에 실제 값이 변경됨
	a.balance += amount
}
