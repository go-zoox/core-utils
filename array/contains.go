package array

// Contains returns true if the array contains the value.
// It is a wrapper for the `Includes` function.
func Contains[K comparable](values []K, value K) bool {
	return Includes(values, value)
}
