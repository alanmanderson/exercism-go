// Package listops provides several array functions on the new array type of IntList
package listops

type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

// IntList list of integers
type IntList []int

// Foldr apply the function from the initial value on all elements from the right
func (arr IntList) Foldr(f binFunc, initial int) int {
	for i := len(arr) - 1; i >= 0; i-- {
		initial = f(arr[i], initial)
	}
	return initial
}

// Foldl apply the function from the initial value on all elements from the left
func (arr IntList) Foldl(f binFunc, initial int) int {
	for i := len(arr) - 1; i >= 0; i-- {
		initial = f(initial, arr[i])
	}
	return initial
}

// Length return the length of the IntList
func (arr IntList) Length() int {
	length := 0
	for range arr {
		length++
	}
	return length
}

// Filter return the elements that return true from the passed in function
func (arr IntList) Filter(f predFunc) IntList {
	newList := make(IntList, 0)
	for _, element := range arr {
		if f(element) {
			newList = append(newList, element)
		}
	}
	return newList
}

// Map execute the function on all of the array elements
func (arr IntList) Map(f unaryFunc) IntList {
	for index, element := range arr {
		arr[index] = f(element)
	}
	return arr
}

// Reverse return an array with the elements reversed
func (arr IntList) Reverse() IntList {
	for i := len(arr) - 1; i >= len(arr)/2; i-- {
		tmp := arr[i]
		index := len(arr) - i - 1
		arr[i] = arr[index]
		arr[index] = tmp
	}
	return arr
}

// Append the elements of the input array to the array
func (arr IntList) Append(app IntList) IntList {
	newList := make([]int, len(arr)+len(app))
	initialLength := len(arr)
	for index, value := range arr {
		newList[index] = value
	}
	for index, value := range app {
		newList[index+initialLength] = value
	}
	return newList
}

// Concat append the arrays passed in from the arrays
func (arr IntList) Concat(concat []IntList) IntList {
	newList := arr[:]
	for _, value := range concat {
		newList = newList.Append(value)
	}
	return newList
}
