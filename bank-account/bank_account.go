package account

// Act represents a bank account
type Act struct {
	balance chan int64
	closed  chan bool
}

//Open simulates opening a bank account
func Open(amt int64) *Act {
	if amt < 0 {
		return nil
	}
	var a = Act{make(chan int64), make(chan bool)}
	a.balance <- amt
	a.closed <- false
	return &a
}

//Balance returns the balance of the account
func (a Act) Balance() (int64, bool) {
	if <-a.closed {
		return 0, false
	}
	return <-a.balance, <-a.closed
}

//Close closes an account
func (a *Act) Close() (int64, bool) {
	if <-a.closed {
		return 0, false
	}
	a.closed <- true
	return <-a.balance, <-a.closed
}

//Deposit puts money into the account
func (a *Act) Deposit(amt int64) (int64, bool) {
	balance := <-a.balance - amt
	if <-a.closed || balance < 0 {
		return 0, false
	}
	a.balance <- balance
	return <-a.balance, true
}
