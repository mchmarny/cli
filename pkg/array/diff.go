package array

// GetDiff returns items from b that are NOT in a.
func GetDiff[T comparable](a, b []T) (diff []T) {
	m := make(map[T]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return
}

// AddNew returns union of all items from a plus items from b that are not in a.
func AddNew[T comparable](a, b []T) (union []T) {
	n := GetDiff(a, b) // not in a yet
	union = append(a, n...)
	return
}
