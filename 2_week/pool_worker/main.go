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
	inputData := []int{1,2,3,4,5,6,7,8,9,10,11,22,}
	maxBuffer := 100
	workerConst := 3

	tasks := make(chan Task, maxBuffer) // jobs
	results := make(chan Task, maxBuffer)

	for w := 0; w < workerConst; w++{
		go taskWorker(w, tasks, results)
	}

	for _, v := range inputData{
		tasks <- Task{
			Id: v,
	
		}
	}
	close(tasks)

	for j := 0; j < len(inputData); j++{
		<- results
	}
}


func taskWorker(id int, jobs <- chan Task, results chan <- Task){
	for j := range jobs{
		fmt.Println("worker: ", id, "started job: ", j)
		j.Index = makeIndex(j.Id)
		time.Sleep(2*time.Second)
		fmt.Println("worker: ", id, "finished job:", j)
		results <- j
	}
}

func makeIndex(id int) int {
	return id - 10 
}