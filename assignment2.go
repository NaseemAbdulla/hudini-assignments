package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Main function to display the menu and handle user input

// Create a structure Task having name and done

type Task struct {
	name        string
	description string
	isDone      bool
}

var tasks []Task

func AddTask(tasks *[]Task, reader *bufio.Reader) {
	fmt.Print("Enter the title of the Task: ")
	title, _ := reader.ReadString('\n')

	fmt.Print("Enter the description of the Task: ")
	desc, _ := reader.ReadString('\n')

	newTask := Task{
		name:        title,
		description: desc,
		isDone:      false,
	}

	*tasks = append(*tasks, newTask)
}

func ListTasks(tasks []Task) {
	// Iterate over all tasks and print

	for i, value := range tasks {
		fmt.Println("Task id: ", i+1)
		fmt.Print("Task name: ", value.name)
		fmt.Println("Task desc: ", value.description)
		fmt.Println("Task done: ", value.isDone)

	}
}

func MarkAsDone(tasks []Task, reader *bufio.Reader) {
	fmt.Print("Enter the Task id :")
	val, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(val))

	tasks[id-1].isDone = true

	fmt.Printf("Task %d is done", id)
}

func delete_at_index(slice []Task, index int) []Task {
	return append(slice[:index], slice[index+1:]...)
}

func DeleteItem(tasks *[]Task, reader *bufio.Reader) {
	fmt.Print("Enter the Task id to delete :")
	val, _ := reader.ReadString('\n')
	dlt, _ := strconv.Atoi(strings.TrimSpace(val))

	*tasks = delete_at_index(*tasks, dlt-1)
	fmt.Printf("The Task %d is deleted \n", dlt)
}
func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Mark Task as Done")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")
		fmt.Print("Enter choice: ")

		input, _ := reader.ReadString('\n')
		choice, err := strconv.Atoi(strings.TrimSpace(input))

		switch choice {
		case 1:
			AddTask(&tasks, reader)
		case 2:
			ListTasks(tasks)
		case 3:
			MarkAsDone(tasks, reader)
		case 4:
			DeleteItem(&tasks, reader)
		case 5:
			os.Exit(0)
		}

		if err != nil {
			fmt.Println("Invalid choice, please try again.")
			continue
		}

	}
}
