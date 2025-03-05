package main

import (
	"projects.com/todo/task"
	"projects.com/todo/todolist"
)

func main() {
	var tasks todolist.TodoList = todolist.TodoList{Tasks: []task.Task{}}
	tasks.AddTask()

	tasks.PrintAllTasks()
}
