package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Id          int
	Title       string
	Description string
}

var id int = 0

var tasks []Task

func AddTask() {
	reader := bufio.NewReader(os.Stdin)
	id++
	fmt.Println("Enter the Title:")
	title, err := reader.ReadString('\n')
	checkNilError(err)

	fmt.Println("Enter the Title:")
	description, err := reader.ReadString('\n')
	checkNilError(err)
	task := Task{id, title, description}
	tasks = append(tasks, task)
}

func ShowTask() {

	for i := 0; i < len(tasks); i++ {
		fmt.Println("Task", tasks[i].Id)
		fmt.Print("Title :", tasks[i].Title)
		fmt.Print("Description :", tasks[i].Description)
		fmt.Println("")
	}
}

func DeleteTask(id int) {
	// var task1 []Task
	// var task2 []Task
	// task1 = append(tasks[:id])
	// task2 = append(tasks[id+1:])

	// var newTask []Task
	// newTask = append(task1, task2...)
	// tasks = newTask

	var newTask []Task

	for i := 0; i < len(tasks); i++ {
		if id-1 != i {
			newTask = append(newTask, tasks[i])
		}
	}

	tasks = newTask
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var option int = 0

	for option != 4 {
		fmt.Println("1.Add Task\n2.Delete Task\n3.Show Tasks\n4.Exit\n")
		fmt.Println("Choose the option:")

		input, err := reader.ReadString('\n')
		checkNilError(err)
		optionFloat, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
		checkNilError(err)
		option = int(optionFloat)

		switch option {
		case 1:
			AddTask()
			fmt.Println("Added Task")
		case 2:
			fmt.Print("Enter the id of the task to be deleted:")
			id1, err := reader.ReadString('\n')
			checkNilError(err)
			idTobePassed, err := strconv.ParseFloat(strings.TrimSpace(id1), 64)
			id2 := int(idTobePassed)
			DeleteTask(id2)
			fmt.Println("Deleted Task")
		case 3:
			ShowTask()
			fmt.Println("Showed Tasks")
		case 4:

		default:
			fmt.Println("Plz enter the proper option")
		}
	}

	// fmt.Println("Tasks:\n", tasks)

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
