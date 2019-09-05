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
