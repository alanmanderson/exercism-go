package strain

// Ints stuff
type Ints []int

// Strings stuff
type Strings []string

// Lists stuff
type Lists [][]int

// IntPredicate stuff
type IntPredicate func(int) bool

// StringPredicate stuff
type StringPredicate func(string) bool

// ListPredicate stuff
type ListPredicate func([]int) bool

// Keep returns a list with the elements that satisfy the predicate
func (l Lists) Keep(pred ListPredicate) (out Lists) {
	for _, elt := range l {
		if pred(elt) {
			out = append(out, elt)
		}
	}
	return
}

// Discard returns a list with the elements that don't satisfy the predicate
func (l Lists) Discard(pred ListPredicate) Lists {
	return l.Keep(ListPredicate(func(i []int) bool { return !pred(i) }))
}

// Keep returns a list with the elements that satisfy the predicate
func (i Ints) Keep(pred IntPredicate) (out Ints) {
	for _, elt := range i {
		if pred(elt) {
			out = append(out, elt)
		}
	}
	return
}

// Discard returns a list with the elements that don't satisfy the predicate
func (i Ints) Discard(pred IntPredicate) (out Ints) {
	return i.Keep(IntPredicate(func(i int) bool { return !pred(i) }))
}

// Keep returns a list with the elements that satisfy the predicate
func (s Strings) Keep(pred StringPredicate) (out Strings) {
	for _, elt := range s {
		if pred(elt) {
			out = append(out, elt)
		}
	}
	return
}

// Discard returns a list with the elements that don't satisfy the predicate
func (s Strings) Discard(pred StringPredicate) (out Strings) {
	return s.Keep(StringPredicate(func(i string) bool { return !pred(i) }))
}
