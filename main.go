package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}

func addCache(appCache *ApplicationCache, iterations int, readChan, updateChan, deleteChan chan string) {
	fmt.Println("Inside addCache")
	for i := 0; i < iterations; i++ {
		key := randomString(5)
		value := randomString(2)
		appCache.Add(key, value, myCostFunc)
		fmt.Printf("Added key %s with value %s \n", key, value)
		switch randomInt(0, 2) {
		case 0:
			readChan <- key
		case 1:
			updateChan <- key
		case 2:
			deleteChan <- key
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func readCache(appCache *ApplicationCache, readChan chan string) {
	fmt.Println("Inside readCache")
	for i := range readChan {
		fmt.Printf("%s key has value %s\n", i, appCache.Get(i))
	}
}

func updateCache(appCache *ApplicationCache, updateChan chan string) {
	fmt.Println("Inside updateCache")
	for i := range updateChan {
		value := randomString(2)
		appCache.Update(i, value)
		fmt.Printf("Update %s key with value %s\n", i, value)
	}
}

func deleteCache(appCache *ApplicationCache, deleteChan chan string) {
	fmt.Println("Inside deleteCache")
	for i := range deleteChan {
		appCache.Delete(i)
		fmt.Printf("Deleted %s key\n", i)
	}
}

func main() {
	pq := make(PriorityQueue, 0)
	myAppCache := ApplicationCache{
		size:     4,
		mapStore: map[string]MapValue{},
		pq:       &pq}

	readChan := make(chan string, 10)
	updateChan := make(chan string, 10)
	deleteChan := make(chan string, 10)
	iterations := 10

	go addCache(&myAppCache, iterations, readChan, updateChan, deleteChan)
	go readCache(&myAppCache, readChan)
	go updateCache(&myAppCache, updateChan)
	go deleteCache(&myAppCache, deleteChan)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		if s == "exit" {
			return
		}
	}
}
