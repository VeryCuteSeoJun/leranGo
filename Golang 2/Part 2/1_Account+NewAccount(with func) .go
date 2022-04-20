package accounts

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
