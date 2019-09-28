package paasio

//NewWriteCounter stuff
type NewWriteCounter struct {
}

//WriteCount stuff
func (c NewWriteCounter) WriteCount() (n int64, nops int) {
	return n, nops
}

//NewReadCounter stuff
type NewReadCounter struct {
}

//ReadCount stuff
func (c NewReadCounter) ReadCount() (n int64, nops int) {
	return n, nops
}

//NewReadWriteCounter stuff
type NewReadWriteCounter struct {
	r ReadCounter
	w WriteCounter
}
