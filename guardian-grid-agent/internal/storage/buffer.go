package storage

import "sync"

type Buffer struct {
	mu       sync.Mutex
	data     map[string][]map[string]interface{}
	capacity int
}

func NewBuffer(capacity int) *Buffer {
	if capacity <= 0 {
		capacity = 100
	}

	return &Buffer{
		data:     make(map[string][]map[string]interface{}),
		capacity: capacity,
	}
}

func (b *Buffer) AddGrouped(key string, d map[string]interface{}) {
	if d == nil {
		return
	}

	b.mu.Lock()
	defer b.mu.Unlock()

	list := b.data[key]

	if len(list) >= b.capacity {
		list = list[1:]
	}

	list = append(list, d)
	b.data[key] = list
}

func (b *Buffer) FlushGrouped() map[string][]map[string]interface{} {
	b.mu.Lock()
	defer b.mu.Unlock()

	out := make(map[string][]map[string]interface{})

	for k, v := range b.data {
		if len(v) > 0 {
			out[k] = v
		}
	}

	// reset
	b.data = make(map[string][]map[string]interface{})

	return out
}
