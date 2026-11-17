type LRUCache struct {
    Capacity int
	Queue []int // contains the keys, least recently used at front
	Data map[int]int // key -> value
}


func Constructor(capacity int) LRUCache {
    cache := LRUCache{
		Capacity: capacity,
		Data: make(map[int]int),
	}
	return cache
}

func (this *LRUCache) Get(key int) int {
	value, exists := this.Data[key]
	if !exists {
		return -1
	}
	
	// Move to the end of LRU queue
	lruIndex := 0
	for lruIndex < len(this.Queue) {
		if this.Queue[lruIndex] == key {
			break // we found the index
		}
		lruIndex++
	}

	this.Queue = append(this.Queue[:lruIndex], this.Queue[lruIndex+1:]...)
	this.Queue = append(this.Queue, key)
	return value
}

func (this *LRUCache) Put(key int, value int) {
	_, exists := this.Data[key]
	if exists {
		_ = this.Get(key) // this updates the LRU
		this.Data[key] = value
		return
	}
	
	this.Data[key] = value
	
	if len(this.Queue) < this.Capacity {
		// Simply add it to the queue
		this.Queue = append(this.Queue, key)
	} else {
		// Evict an entry first
		evicted := this.Queue[0]
		delete(this.Data, evicted)
		this.Queue = append(this.Queue[1:], key)
	}
	
}
