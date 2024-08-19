package main

import "fmt"

type Task struct {
	name  string
	time  int
	event string
}

var tasks []Task = []Task{
	{name: "main", time: 0, event: "start"},
	{name: "subTask1", time: 5, event: "start"},
	{name: "subTask1", time: 10, event: "end"},
	{name: "subTask2", time: 15, event: "start"},
	{name: "subTask2", time: 20, event: "end"},
	{name: "main", time: 25, event: "end"},
}

func calculateExecutionTime(tasks []Task) map[string]int {
	taskExecTimes := make(map[string]int)

	for _, task := range tasks {
		if value, ok := taskExecTimes[task.name]; ok {
			taskExecTimes[task.name] = task.time - value
		} else {
			taskExecTimes[task.name] = task.time
		}
	}

	return taskExecTimes
}

func printTaskTimes() {
	for key, value := range calculateExecutionTime(tasks) {
		fmt.Printf("%v: %v \n", key, value)
	}
}
