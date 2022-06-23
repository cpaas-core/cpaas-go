package account

import "sync"

// Define the Account type here.
type Account struct {
	mutex  sync.Mutex
	amount int64
	opened bool
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{amount: amount, opened: true}
}

func (a *Account) Balance() (int64, bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if !a.opened {
		return 0, false
	}
	return a.amount, a.opened
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if !a.opened {
		return 0, false
	}
	if a.amount+amount < 0 {
		return 0, false
	}
	a.amount += amount
	return a.amount, true
}

func (a *Account) Close() (int64, bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if !a.opened {
		return 0, false
	}
	a.opened = false
	return a.amount, !a.opened
}
