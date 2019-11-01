package account

// Act represents a bank account
type Act struct {
	balance chan int64
}

//Open simulates opening a bank account
func Open(amt int64) *Act {
	if amt < 0 {
		return nil
	}
	var a = Act{make(chan int64)}
	a.balance <- amt
	return &a
}

//Balance returns the balance of the account
func (a Act) Balance() (int64, bool) {
	if <-a.closed {
		return 0, false
	}
	amt, ok := <-a.balance
	return amt, ok
}

//Close closes an account
func (a *Act) Close() (int64, bool) {
	var amt = 0
	amt, ok := <-a.balance
	close(a.balance)

	if !ok {
		return 0, false
	}
	return amt, ok
}

//Deposit puts money into the account
func (a *Act) Deposit(amt int64) (int64, bool) {
	balance, ok := <-a.balance - amt
	if <-a.closed || balance < 0 {
		return 0, false
	}
	a.balance <- balance
	return <-a.balance, true
}
