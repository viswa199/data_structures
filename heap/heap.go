package heap

type Heap struct {
	Array []int
}

func (heap *Heap) MinHeap() {
	i := len(heap.Array) / 2
	for j := i - 1; j >= 0; j-- {
		if 2*j+1 < 2*i {
			if heap.Array[2*j+1] < heap.Array[j] {
				heap.Array[2*j+1], heap.Array[j] = heap.Array[j], heap.Array[2*j+1]
			}
		}
		if 2*j+2 < 2*i {
			if heap.Array[2*j+2] < heap.Array[j] {
				heap.Array[2*j+2], heap.Array[j] = heap.Array[j], heap.Array[2*j+2]
			}
		}
	}
}

func (heap *Heap) MaxHeap() {
	i := len(heap.Array) / 2
	for j := i - 1; j >= 0; j-- {
		if 2*j+1 < 2*i {
			if heap.Array[2*j+1] > heap.Array[j] {
				heap.Array[2*j+1], heap.Array[j] = heap.Array[j], heap.Array[2*j+1]
			}
		}
		if 2*j+2 < 2*i {
			if heap.Array[2*j+2] > heap.Array[j] {
				heap.Array[2*j+2], heap.Array[j] = heap.Array[j], heap.Array[2*j+2]
			}
		}
	}
}
