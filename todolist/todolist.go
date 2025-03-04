package todolist

import "projects.com/todo/task"

type TodoList struct {
	Tasks []task.Task //Slice of  tasks [TASK, TASK, TASK]
}
