package collections

import "fmt"

func OrderedMaps() {
	orderedMap := NewOrderedMap[string, string]()

	orderedMap.Set("me", "Abhishek Ghosh")
	orderedMap.Set("friend-1", "Nasim Molla")
	orderedMap.Set("friend-2", "Sayan Mandal")
	orderedMap.Set("friend-3", "Abhishek Pal")
	orderedMap.Set("friend-4", "Bishal Mukherjee")

	// deleteing the friend 2
	orderedMap.Delete("friend-2")

	iterator := orderedMap.Iterator()

	for {
		index, key, value := iterator()
		if index == nil {
			break
		}
		fmt.Printf("ordered map index : %d, key : %s, value : %s  \n", *index, *key, value)
	}
}

func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		store: map[K]V{},
		keys:  []K{},
	}
}

// key will always be comparable but value can be anything
type OrderedMap[K comparable, V any] struct {
	store map[K]V
	keys  []K
}

// Get will return the value associated with the key.
// If the key doesn't exist, the second return value will be false.
func (orderedMap *OrderedMap[K, V]) Get(key K) (V, bool) {
	val, exists := orderedMap.store[key]
	return val, exists
}

// Set will store a key-value pair. If the key already exists,
// it will overwrite the existing key-value pair.
func (orderedMap *OrderedMap[K, V]) Set(key K, val V) {
	// exists value will be false if the key does not exist in the keys list
	if _, exists := orderedMap.store[key]; !exists {
		// adding to the existing keys list
		orderedMap.keys = append(orderedMap.keys, key)
	}
	// overriding the map[key]
	orderedMap.store[key] = val
}

// Delete will remove the key and its associated value.
func (orderedMap *OrderedMap[K, V]) Delete(key K) {
	// first delete the key from the map
	delete(orderedMap.store, key)
	// Find the index of the key in the keys list
	idx := -1
	for i, val := range orderedMap.keys {
		if val == key {
			idx = i
			break
		}
	}
	// if key exists then we have to remove the key from the keys list
	// list... is the spread operator, which will spread the list into items
	if idx != -1 {
		orderedMap.keys = append(orderedMap.keys[:idx], orderedMap.keys[idx+1:]...)
	}
}

// Iterator is used to loop through the stored key-value pairs.
// it is returning a closure or an anonymous function thatreturns the index, key and value.
func (orderedMap *OrderedMap[K, V]) Iterator() func() (*int, *K, V) {
	var keys = orderedMap.keys
	index := 0
	return func() (_ *int, _ *K, _ V) {
		if index > len(keys)-1 {
			return
		}
		row := keys[index]
		index++
		return &[]int{index - 1}[0], &row, orderedMap.store[row]
	}
}
