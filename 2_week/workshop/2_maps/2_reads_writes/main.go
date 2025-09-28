package readswrites

import (
	"fmt"
	"sync"
)


// func main(){ 
// 	storage := make(map[int]int, 1000)

// 	wg := sync.WaitGroup{}
// 	reads := 1000
// 	writes := 1000
// 	mu := sync.Mutex{}

// 	wg.Add(writes)
// 	for i := 0; i < writes; i++{
// 		i := i 
// 		go func() {
// 			defer wg.Done()
// 			mu.Lock()
// 			defer mu.Unlock()
// 			storage[i] = i
// 		}()
// 	}

// 	wg.Add(reads)
// 	for i := 0; i < reads; i++{
// 		i := i 
// 		go func() {
// 			defer wg.Done()
// 			_, _  = storage[i]
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println(storage)
// }


func main(){ 
	storage := make(map[int]int, 1000)

	syncM := sync.Map{} // нет типизации, размера
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	rmu := sync.RWMutex{}
	
	reads := 1000
	writes := 1000
	
	wg.Add(writes)
	for i := 0; i < writes; i++{
		i := i 
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			storage[i] = i
			syncM.Store(i, i)
		}()
	}

	wg.Add(reads)
	for i := 0; i < reads; i++{
		i := i 
		go func() {
			defer wg.Done()

			rmu.RLock() 
   			defer rmu.RUnlock()

			syncM.Load(i)
			_, _  = storage[i]
			
		}()
	}
	wg.Wait()
	fmt.Println(storage)
}