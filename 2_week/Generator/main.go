package main

import "fmt"


type Task struct {
	Id int 
	Index int
}

func main(){
	for task := range taskGnerator(1, 5){
		fmt.Println(task.Id, task.Index)
	}
}


func taskGnerator(start int, end int) <-chan Task{
	ch := make(chan Task)

	go func (ch chan Task) { 
		for j := start; j <= end; j ++{
			// Блокируем операцию
			ch <- Task{
				Id: j,
				Index: makeIndex(j),
			}
		}
		close(ch)
	}(ch)
	return ch
}


func makeIndex(id int) int {
	return id - 10 
}