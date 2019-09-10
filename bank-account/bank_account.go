package account

// Act represents a bank account
type Act struct {
	balance int64
	closed  bool
}

//Open simulates opening a bank account
func Open(amt int64) *Act {
	if amt < 0 {
		return nil
	}
	var a = Act{amt, false}
	return &a
}

//Balance returns the balance of the account
func (a Act) Balance() (int64, bool) {
	if a.closed {
		return 0, false
	}
	return a.balance, true
}

//Close closes an account
func (a *Act) Close() (int64, bool) {
	if a.closed {
		return 0, false
	}
	a.closed = true
	return a.balance, a.closed
}

//Deposit puts money into the account
func (a *Act) Deposit(amt int64) (int64, bool) {
	if a.closed || a.balance+amt < 0 {
		return 0, false
	}
	a.balance += amt
	return a.balance, true
}
