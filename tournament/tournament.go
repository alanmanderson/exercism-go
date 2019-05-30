package tournament

import (
	"fmt"
	"io"
	"strings"
)

// Tally returns the tally of wins by team
func Tally(r io.Reader, w io.Writer) error {
	var sb strings.Builder
	b := make([]byte, 8)
	for {
		bytes, err := r.Read(b)
		fmt.Printf("%d %v %v\n", bytes, err, string(b))
		if err == io.EOF {
			break
		}
		sb.WriteString(string(b))
	}
	return nil
}

func main() {

	fmt.Printf("Hello World")

}
