package storage

package storage

import "sync"

type Buffer struct {
	mu       sync.Mutex
	data     []map[string]interface{}
	capacity int
}

// PUBLIC

func NewBuffer(capacity int) *Buffer {
	if capacity <= 0 {
		capacity = 100 // default safety
	}

	return &Buffer{
		data:     make([]map[string]interface{}, 0, capacity),
		capacity: capacity,
	}
}

func (b *Buffer) Add(d map[string]interface{}) {
	if d == nil {
		return
	}

	b.mu.Lock()
	defer b.mu.Unlock()

	// prevent memory explosion
	if len(b.data) >= b.capacity {
		// drop oldest (simple ring behavior)
		b.data = b.data[1:]
	}

	b.data = append(b.data, d)
}

func (b *Buffer) Flush() []map[string]interface{} {
	b.mu.Lock()
	defer b.mu.Unlock()

	if len(b.data) == 0 {
		return nil
	}

	out := make([]map[string]interface{}, len(b.data))
	copy(out, b.data)

	// reset buffer
	b.data = b.data[:0]

	return out
}

func (b *Buffer) Peek() []map[string]interface{} {
	b.mu.Lock()
	defer b.mu.Unlock()

	out := make([]map[string]interface{}, len(b.data))
	copy(out, b.data)

	return out
}

func (b *Buffer) Size() int {
	b.mu.Lock()
	defer b.mu.Unlock()

	return len(b.data)
}