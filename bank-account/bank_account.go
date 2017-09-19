package account

import (
	"sync"
)

const testVersion = 1

type Account struct {
	m       sync.Mutex
	balance int64
	closed  bool
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}

	return &Account{balance: initialDeposit}
}

func (a *Account) Close() (payout int64, ok bool) {
	a.m.Lock()
	defer a.m.Unlock()
	if a.closed {
		return 0, false
	}

	a.closed = true
	return a.balance, true
}

func (a *Account) Balance() (int64, bool) {
	if a.closed {
		return -1, false
	}

	return a.balance, true
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	if a.closed {
		return -1, false
	}

	a.m.Lock()
	defer a.m.Unlock()
	newBalance = a.balance + amount
	if newBalance < 0 {
		return a.balance, false
	}
	a.balance = newBalance
	return newBalance, true
}
