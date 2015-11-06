package pp_binary_search

// CmpFunc takes an index into a collection and returns
// < 0 if the query value is less than the collection value
// > 0 if the query value is greater than the collection value
// and == 0 if the values are equal
type CmpFunc func(i int) int

func Search(colLen int, f CmpFunc) (int, bool) {
	if colLen == 0 {
		return -1, false
	}

	left := 0
	right := colLen - 1

	for right-left >= 0 {
		mid := left + ((right - left) / 2)
		dif := f(mid)
		if dif < 0 {
			right = mid - 1
		} else if 0 < dif {
			left = mid + 1
		} else {
			return mid, true
		}
	}

	return -1, false
}
