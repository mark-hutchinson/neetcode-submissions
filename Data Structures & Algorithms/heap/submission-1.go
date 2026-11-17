type MinHeap struct {
	Data []int
	Size int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

func (mh *MinHeap) Push(val int) {
	mh.Data = append(mh.Data, val)
	mh.Size++
	index := mh.Size - 1
	mh.bubbleUp(index)
}

func (mh *MinHeap) Pop() int {
	if mh.Size == 0 {
		return -1
	}
	result := mh.Data[0]

	// take last node and move it to start
	mh.Data[0] = mh.Data[mh.Size-1]
	mh.Size--
	// shrink the underlying array
	mh.Data = mh.Data[:mh.Size]

	mh.bubbleDown(0)

	return result
}

func (mh *MinHeap) bubbleDown(i int) {
	leftIndex := left(i)
	rightIndex := right(i)


	smallestIndex := i
	if leftIndex < mh.Size && mh.Data[leftIndex] < mh.Data[smallestIndex] {
		smallestIndex = leftIndex
	}
	if rightIndex < mh.Size && mh.Data[rightIndex] < mh.Data[smallestIndex] {
		smallestIndex = rightIndex
	}
	if i == smallestIndex {
		return
	}

	// swap smallest index and i
	mh.Data[i], mh.Data[smallestIndex] = mh.Data[smallestIndex], mh.Data[i]
	mh.bubbleDown(smallestIndex)
}

func (mh *MinHeap) bubbleUp(i int) {
	if i == 0 {
		return // already at top
	}

	parentIndex := parent(i)
	if mh.Data[i] < mh.Data[parentIndex]  {
		// swap i and parent
		mh.Data[i], mh.Data[parentIndex] = mh.Data[parentIndex], mh.Data[i]
		mh.bubbleUp(parentIndex)
	}
}

func left(i int) int {
	return i*2 +1
}

func right(i int) int {
	return i*2 +2
}

func parent (i int) int {
	return (i-1)/2
}

func (mh *MinHeap) Top() int {
	if mh.Size == 0 {
		return -1
	}
	return mh.Data[0]
}

func (mh *MinHeap) Heapify(nums []int) {
	mh.Data = nil
	mh.Size = 0

	for _, num := range nums {
		mh.Push(num)
	}
}
