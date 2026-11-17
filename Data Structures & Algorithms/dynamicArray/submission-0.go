type DynamicArray struct {
    data []int
}

func NewDynamicArray(capacity int) *DynamicArray {
    da := &DynamicArray{}
    da.data = make([]int, 0, capacity)
    return da
}

func (da *DynamicArray) Get(i int) int {
    return da.data[i]
}

func (da *DynamicArray) Set(i int, n int) {
    da.data[i] = n
}

func (da *DynamicArray) Pushback(n int) {
    // If we call pushback(int n) but the array is full, we should resize() the array first. (increase capacity)
    if da.GetSize() >= da.GetCapacity() {
        da.resize()
    }
    da.data = append(da.data, n)
}

func (da *DynamicArray) Popback() int {
    index := len(da.data) -1
    element := da.data[index]
    da.data = da.data[:index]
    return element
}

func (da *DynamicArray) resize() {
    newCap := cap(da.data) * 2
    if newCap == 0 {
        newCap = 1
    }
    newData := make([]int, len(da.data), newCap)
    for index, val := range da.data {
        newData[index] = val
    }
    da.data = newData
}

func (da *DynamicArray) GetSize() int {
    return len(da.data)
}

func (da *DynamicArray) GetCapacity() int {
    return cap(da.data)
}
