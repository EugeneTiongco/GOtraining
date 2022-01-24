package cards

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// the given index exists in the slice or not.
func GetItem(slice []int, index int) (int, bool) {
	if index >= len(slice) || index < 0 {
		return 0, false
	}
	return slice[index], true

}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index >= len(slice) || index < 0 {
		slice = append(slice, value)
		return slice
	} else {
		slice[index] = value
	}
	return slice
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {

	if length > 0 {
		slice := make([]int, length)
		for i := 0; i < length; i++ {
			slice[i] = value
		}
		return slice
	}
	var slice []int
	return slice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index < len(slice) && index >= 0 {
		tail := slice[index+1:]
		slice = slice[:index]
		slice = append(slice, tail...)

		return slice
	}
	return slice
}
