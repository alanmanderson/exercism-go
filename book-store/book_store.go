package bookstore

//Cost determine the cost of the shopping cart
func Cost(books []int) int {
	count := 0
	for range books {
		count++
	}
	return count * 800
}
