package erratum

import (
    "errors"
    "fmt"
    )

//Use stuff
func Use(o ResourceOpener, input string) error {
    for {
        r, err := o()
        if err != nil {
            switch err.(type) {
                case TransientError:
                    continue
                default:
                    fmt.Println("default")
                    return errors.New("too awesome")
            }
            return err
        }
        r.Frob(input)
        
        defer func(r Resource) {
            if x := recover(); x != nil {
                r.Defrob(input)
            }
            r.Defrob(input)
        }(r)
        break;
    }
	return nil
}
