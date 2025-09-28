package main

import (
	"fmt"
	"sync"
)

// func main(){
// 	counter := 20
// 	for i := 0; i < counter; i++{
// 		go func(){
// 			fmt.Println(i * i)
// 		}()
// 	}

// 	time.Sleep(time.Second)
// }

func main(){
	counter := 20
	wg := sync.WaitGroup{}

	for i := 0; i < counter; i++{
		wg.Add(1)
		go func(in int){
			defer wg.Done()
			fmt.Println(in * in)
		}(i)
		
	}

	wg.Wait()
}
