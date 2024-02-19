package collections

import "fmt"

func OrderedMaps() {
	m := NewOrderedMap[int, string]()

	m.Set(1, "string1")
	m.Set(2, "string2")
	m.Set(3, "string3")
	m.Set(4, "string4")
	m.Delete(3)
	m.Delete(8)

	iterator := m.Iterator()

	for {
		i, k, v := iterator()
		if i == nil {
			break
		}
		fmt.Println(*i, *k, v+" is a string")
	}

}

type OrderedMap[K comparable, V any] struct {
	store map[K]V
	keys  []K
}

// Go 1 Constructor
func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		store: map[K]V{},
		keys:  []K{},
	}
}

// Get will return the value associated with the key.
// If the key doesn't exist, the second return value will be false.
func (o *OrderedMap[K, V]) Get(key K) (V, bool) {
	val, exists := o.store[key]
	return val, exists
}

// Set will store a key-value pair. If the key already exists,
// it will overwrite the existing key-value pair.
func (o *OrderedMap[K, V]) Set(key K, val V) {
	if _, exists := o.store[key]; !exists {
		o.keys = append(o.keys, key)
	}
	o.store[key] = val
}

// Delete will remove the key and its associated value.
func (o *OrderedMap[K, V]) Delete(key K) {
	delete(o.store, key)

	// Find key in slice
	idx := -1
	for i, val := range o.keys {
		if val == key {
			idx = i
			break
		}
	}
	if idx != -1 {
		o.keys = append(o.keys[:idx], o.keys[idx+1:]...)
	}
}

// Iterator is used to loop through the stored key-value pairs.
// The returned anonymous function returns the index, key and value.
func (o *OrderedMap[K, V]) Iterator() func() (*int, *K, V) {
	var keys = o.keys
	j := 0
	return func() (_ *int, _ *K, _ V) {
		if j > len(keys)-1 {
			return
		}

		row := keys[j]
		j++

		return &[]int{j - 1}[0], &row, o.store[row]
	}
}
