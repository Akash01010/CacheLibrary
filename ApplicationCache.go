package main

import (
	"container/heap"
	"fmt"
	"sync"
)

var timeVar int = 0

type costFunc func(int) int

func myCostFunc(par int) int {
	timeVar = par + 1
	return timeVar
}

type MapValue struct {
	value string
	ref   *Item
}

type ApplicationCache struct {
	mu       sync.Mutex
	size     int
	mapStore map[string]MapValue
	pq       *PriorityQueue
}

func (cache ApplicationCache) Add(key string, value string, myFunc costFunc) {
	cache.mu.Lock()
	_, ok := cache.mapStore[key]
	if ok {
		fmt.Printf("Key already present in the cache\n")
	} else {
		priori := myFunc(timeVar)
		item := &Item{
			key:      key,
			priority: priori,
		}
		fmt.Printf("Adding %s to the cache with pri %d\n", key, priori)
		cache.mapStore[key] = MapValue{value, item}
		heap.Push(cache.pq, item)
		if len(cache.mapStore) > cache.size {
			removedItem := heap.Pop(cache.pq).(*Item)
			fmt.Printf("Removing %s from the cache with priority %d with index %d\n", removedItem.key, removedItem.priority, removedItem.index)
			delete(cache.mapStore, removedItem.key)
		}
	}
	cache.mu.Unlock()
}

func (cache ApplicationCache) Get(key string) string {
	cache.mu.Lock()
	value, ok := cache.mapStore[key]
	if ok {
		fmt.Printf("Cache Hit!\n")
		cache.pq.update(value.ref, value.ref.key, myCostFunc(timeVar))
		defer cache.mu.Unlock()
		return value.value
	} else {
		fmt.Printf("Cache Miss!\n")
		//maybe throw error
		defer cache.mu.Unlock()
		return "-1"
	}
}

func (cache ApplicationCache) Update(key, value string) {
	cache.mu.Lock()
	val, ok := cache.mapStore[key]
	if ok {
		cache.mapStore[key] = MapValue{value, val.ref}
		cache.pq.update(val.ref, val.ref.key, myCostFunc(timeVar))
		fmt.Printf("Updating %s key in cache with value %s\n", key, value)
	} else {
		fmt.Printf("%s key not found in cache\n", key)
	}
	cache.mu.Unlock()
}

func (cache ApplicationCache) Delete(key string) {
	cache.mu.Lock()
	value, ok := cache.mapStore[key]
	if ok {
		cache.pq.update(value.ref, value.ref.key, -1)
		removedItem := heap.Pop(cache.pq).(*Item)
		fmt.Printf("Removing %s from cache\n", removedItem.key)
		delete(cache.mapStore, removedItem.key)
	} else {
		fmt.Printf("Key not present in the cache\n")
	}
	cache.mu.Unlock()
}
