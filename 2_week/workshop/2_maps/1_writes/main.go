package main

import (
	"fmt"
	"sync"
)

// func main(){
// 	var storage map[int]int
// 	wg := sync.WaitGroup{}
// 	writes := 1000

// 	wg.Add(writes)
// 	for i := 0; i < writes; i++{
// 		i := i
// 		go func(){
// 			defer wg.Done()
// 			storage[i] = i
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println(storage)
// }

func main(){
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	writes := 1000
	storage := make(map[int]int, writes)

	wg.Add(writes)
	for i := 0; i < writes; i++{
		i := i 
		go func(){
			defer wg.Done()
			mu.Lock()
			storage[i] = i
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(storage)
}