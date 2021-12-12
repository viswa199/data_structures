package sorting

func Quicksort(a []int, lo, hi int) {
	if lo >= 0 && hi >= 0 && lo < hi {
		p := Partition(a, lo, hi)
		Quicksort(a, lo, p-1)
		Quicksort(a, p+1, hi)
	}
}

func Partition(a []int, lo, hi int) int {
	j := lo - 1
	pivot := a[hi]
	for i := lo; i <= hi; i++ {
		if a[i] <= pivot {
			j += 1
			temp := a[i]
			a[i] = a[j]
			a[j] = temp
		}
	}
	return j
}
