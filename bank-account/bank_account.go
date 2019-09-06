package account

// Act represents a bank account
type Act struct {
	balance float32
	closed  bool
}

//Open simulates opening a bank account
func Open(amt float32) *Act {
	var a = Act{amt, false}
	return &a
}

//Balance returns the balance of the account
func (a Act) Balance() (float32, error) {
	return a.balance, nil
}
