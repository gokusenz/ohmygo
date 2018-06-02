package ohmygo

// fibonacci is a function that returns
// a function that returns an int.
// https://tour.golang.org/moretypes/26
func fibonacci() func() int {
	return func() int {
		return 0
	}
}
