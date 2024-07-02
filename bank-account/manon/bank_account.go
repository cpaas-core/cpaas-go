package account

import "sync"

// Account type
type Account struct {
	sync.RWMutex
	amount int64
	closed bool
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{amount: amount, closed: false}
}

func (a *Account) Balance() (int64, bool) {
	a.RLock()
	defer a.RUnlock()

	if a.closed {
		return 0, false
	}
	return a.amount, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.Lock()
	defer a.Unlock()

	if a.closed {
		return 0, false
	}
	if a.amount+amount < 0 {
		return 0, false
	}
	a.amount += amount
	return a.amount, true
}

func (a *Account) Close() (int64, bool) {
	a.Lock()
	defer a.Unlock()

	if a.closed {
		return a.amount, false
	}
	balance := a.amount
	a.closed = true
	a.amount = 0
	return balance, true
}
