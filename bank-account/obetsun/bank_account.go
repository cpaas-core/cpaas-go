package account

import (
	"sync"
)

type Account struct {
	sync.RWMutex
	balance int64
	open    bool
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{balance: amount, open: true}
}

func (a *Account) Balance() (int64, bool) {
	a.RLock()
	defer a.RUnlock()
	if a == nil || !a.open {
		return 0, false
	}
	return a.balance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.Lock()
	defer a.Unlock()
	if a == nil || !a.open {
		return 0, false
	}

	if a.balance+amount < 0 {
		return a.balance, false
	}
	a.balance += amount
	return a.balance, true
}

func (a *Account) Close() (balance int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.open {
		balance = a.balance
		a.open = false
		ok = true
	}
	return balance, ok
}
