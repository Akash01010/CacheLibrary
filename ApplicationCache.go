package main

import (
	"container/heap"
	"fmt"
)

var time int = 0

type costFunc func(int) int

func myCostFunc(par int) int {
	time = par + 1
	return time
}

type MapValue struct {
	value string
	ref   *Item
}

type ApplicationCache struct {
	size     int
	mapStore map[string]MapValue
	pq       *PriorityQueue
}

func (cache ApplicationCache) Add(key string, value string, myFunc costFunc) {
	_, ok := cache.mapStore[key]
	if ok {
		fmt.Printf("Key already present in the cache\n")
	} else {
		priori := myFunc(time)
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
}

func (cache ApplicationCache) Get(key string) string {
	value, ok := cache.mapStore[key]
	if ok {
		fmt.Printf("Cache Hit!\n")
		cache.pq.update(value.ref, value.ref.key, myCostFunc(time))
		return value.value
	} else {
		fmt.Printf("Cache Miss!\n")
		//maybe throw error
		return "-1"
	}
}

func (cache ApplicationCache) Update(key, value string) {
	val, ok := cache.mapStore[key]
	if ok {
		cache.mapStore[key] = MapValue{value, val.ref}
		cache.pq.update(val.ref, val.ref.key, myCostFunc(time))
		fmt.Printf("Updating %s key in cache with value %s\n", key, value)
	} else {
		fmt.Printf("%s key not found in cache\n", key)
	}

}

func (cache ApplicationCache) Delete(key string) {
	value, ok := cache.mapStore[key]
	if ok {
		cache.pq.update(value.ref, value.ref.key, -1)
		removedItem := heap.Pop(cache.pq).(*Item)
		fmt.Printf("Removing %s from cache\n", removedItem.key)
		delete(cache.mapStore, removedItem.key)
	} else {
		fmt.Printf("Key not present in the cache\n")
	}
}
