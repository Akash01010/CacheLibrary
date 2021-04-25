package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	pq := make(PriorityQueue, 0)
	myAppCache := ApplicationCache{4, map[string]MapValue{}, &pq}
	myAppCache.Add("a", "1", myCostFunc)
	myAppCache.Add("b", "2", myCostFunc)
	myAppCache.Add("c", "3", myCostFunc)
	myAppCache.Add("d", "4", myCostFunc)
	myAppCache.Get("d")
	myAppCache.Update("b", "2.2")
	myAppCache.Delete("c")
	fmt.Printf("%s\n", myAppCache.Get("a"))
	fmt.Printf("%s\n", myAppCache.Get("b"))

	for i, s := range pq {
		fmt.Println(i, s)
	}
}
