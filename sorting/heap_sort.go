package sorting

func HeapSort(a []int, n int) {
	for i := n; i >= 1; i-- {
		heapify(a, 0, i)
	}

}

func heapify(a []int, low, high int) {
	maxheap(a, low, high)
	a[low], a[high-1] = a[high-1], a[low]
}

func maxheap(a []int, low, high int) {
	i := (high - low) / 2
	for j := i - 1; j >= 0; j-- {
		if (2*j + 1) < high {
			if a[j] < a[2*j+1] {
				a[j], a[2*j+1] = a[2*j+1], a[j]
			}
		}
		if (2*j + 2) < high {
			if a[j] < a[2*j+2] {
				a[j], a[2*j+2] = a[2*j+2], a[j]
			}
		}
	}
}
