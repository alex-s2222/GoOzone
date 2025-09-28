package main

import (
	"fmt"
	"time"
)

type Task struct {
	Id int 
	Index int
}

func main(){
	// НЕПРАВИЛЬНО (НЕКРАСИВО)
	//for task := range fillIndex(fillIndex(fillIndex(build([]int{1, 2, 3 })))){
	//	fmt.Println(task.Id, task.Index)
	//}
	in := build([]int{40, 50, 60, 70, 80, 80,2000,2031})
	out := fillIndex(in)
	for task := range out{
		fmt.Println(task.Id, task.Index)
	}
}

func build(in [] int) <- chan Task{
	out := make(chan Task)
	go func() {
		for _, n := range in {
			time.Sleep(1*time.Second)
			out <- Task{
				Id: n,
			}
		}
		close(out)
	}()
	return out
}

func makeIndex(id int) int {
	return id - 10 
}

func fillIndex(in <- chan Task) <- chan Task {
	out := make(chan Task)
	go func() {
		for task := range in {
			task.Index = makeIndex(task.Id)
			out <- task
		}
		close(out)
	}()

	return out
}


