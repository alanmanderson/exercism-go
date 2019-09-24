package erratum

import "errors"

//Use stuff
func Use(o ResourceOpener, input string) error {
	return errors.New("error")
}
