// returns a string indicating one for a person and one for me
package twofer

import "fmt"

// Specifies with whom we are sharing the item
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
