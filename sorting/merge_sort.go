package sorting

func TopDownSplitMerge(a []int, iBegin, iEnd int) {
	if iEnd-iBegin <= 0 {
		return
	}
	iMiddle := (iEnd + iBegin) / 2
	TopDownSplitMerge(a, iBegin, iMiddle)
	TopDownSplitMerge(a, iMiddle+1, iEnd)
	MergeSortedArrays(a, iBegin, iMiddle, iEnd)
}

func MergeSortedArrays(a []int, iBegin, iMiddle, iEnd int) {
	i := iBegin
	j := iMiddle + 1

	var res []int
    
	//merge two sorted list a[iBegin:iMiddle] and a[iMiddle+1:iEnd]
	for i <= iMiddle && j <= iEnd {
		if a[i] <= a[j] {
			res = append(res, a[i])
			i += 1
		} else {
			res = append(res, a[j])
			j += 1
		}
	}
	for i <= iMiddle {
		res = append(res, a[i])
		i += 1
	}
	for j <= iEnd {
		res = append(res, a[j])
		j += 1
	}
    
	//now copy the elements of slice to actual array.
	k := 0
	for i = iBegin; i <= iEnd; i++ {
		a[i] = res[k]
		k++
	}

}