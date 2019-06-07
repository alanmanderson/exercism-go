package dominoes

import "fmt"

// Domino contains a top and a bottom number
type Domino [2]int

// MakeChain determines if a valid chain can be made
func MakeChain(stones []Domino) ([]Domino, bool) {
	stack := make([]Domino, 0)
	if len(stones) == 0 {
		return nil, true
	}
	stack = append(stack, stones[0])
	stack = append(stack, flipDomino(stones[0]))
	fmt.Printf("%v", stack)
	return nil, true
}

func flipDomino(stone Domino) (out Domino) {
	return Domino{stone[1], stone[0]}
}
