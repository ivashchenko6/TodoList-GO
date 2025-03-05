package todolist

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"projects.com/todo/task"
)

type TodoList struct {
	Tasks []task.Task //Slice of  tasks [TASK, TASK, TASK]
}

func (todos *TodoList) AddTask() bool {
	var newTask task.Task = task.Task{}
	reader := bufio.NewReader(os.Stdin)

	id, err := gonanoid.Generate("abcdef", 6)

	if err != nil {
		panic(err)

	}
	newTask.ID = id

	fmt.Print("Task name: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	if len(text) < 2 {
		fmt.Println("Length must be greater than 2 symbols:")
		return false
	}
	newTask.Title = text
	newTask.IsItDone = false

	todos.Tasks = append(todos.Tasks, newTask)
	return true
}

func (todos *TodoList) RemoveTask(todoNumber int) bool {
	if (todoNumber) > len(todos.Tasks) {
		fmt.Println("ERROR. task number is out of range")
		return false
	}

	todos.Tasks = append(todos.Tasks[:todoNumber-1], todos.Tasks[todoNumber:]...)
	return true
}

func (todos TodoList) PrintAllTasks() {
	for index, task := range todos.Tasks {
		fmt.Printf("%d. Title: %s, Status: %v\n", index+1, task.Title, task.IsItDone)
	}
}
