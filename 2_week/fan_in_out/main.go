package main

import (
	"fmt"
	"time"
	"sync"
)

type Task struct{
	Id int
	Index int
}

func main(){
	in := build([]int{1,2,3})
	out1 := fillIndex(in)
	// in2 := build([]int{3,5,6})
	out2 := fillIndex(in)

	for task := range mergeTasks(out1, out2) {
		fmt.Println(task.Id, task.Index)
	}
}


// Fan-in
func mergeTasks(in ...<-chan Task) <- chan Task {
	// ... неопределенное кол-во каналлов  []
	wg := sync.WaitGroup{}
	out := make(chan Task)

	// Для каждого входного канала in 
	// output копирует значение из с в out
	// до тех пор пока c не закрыт

	output := func(c <-chan Task){
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(in))
	for _, v := range in {
		go output(v)
	}

	go func(){
		wg.Wait()
		close(out)
	}()

	return out

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


