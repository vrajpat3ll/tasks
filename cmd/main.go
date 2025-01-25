// tasks add
// tasks list
// tasks complete

package main

import (
	"fmt"
	"os"
	"strconv"
	"todo-cli"
)

const todoFile = "C:/commands/.todos.json"

func RED(msg string) string {
	return "\u001b[31m" + msg + "\u001b[0m"
}

func ParseArgs() []string {
	const NUMBER_OF_FEATURES uint8 = 4
	// ret = []string {add, list, complete, delete}
	ret := make([]string, NUMBER_OF_FEATURES)
	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		if arg == "add" {
			ret[0] = os.Args[i+1]
		} else if arg == "list" {
			ret[1] = "true"
		} else if arg == "complete" {
			ret[2] = os.Args[i+1]
		} else if arg == "delete" {
			ret[3] = os.Args[i+1]
		}
	}
	return ret
}
func Usage() {
	fmt.Println("Usage of tasks.exe:")
	fmt.Println("\u001b[1madd\u001b[0m string:")
	fmt.Println("    - description of the task at hand")
	fmt.Println("")
	fmt.Println("\u001b[1mlist\u001b[0m bool:")
	fmt.Println("    - flag to list all the tasks")
	fmt.Println("")
	fmt.Println("\u001b[1mcomplete\u001b[0m int:")
	fmt.Println("    - index of todo to complete")
	fmt.Println("")
	fmt.Println("\u001b[1mdelete\u001b[0m int:")
	fmt.Println("    - index of todo to delete")
}
func main() {
	if len(os.Args) < 2 {
		Usage()
		os.Exit(1)
	}
	
	// args is a [4]string wherin
	// 1st element is the task to be added
	// 2nd element is to tell us whether to display the todos
	// 3rd element is an id in string format to tell us which todo is completed
	// 4th element is an id in string format to tell us which todo is to be deleted
	args := ParseArgs()
	// fmt.Println(args)
	todos := &tasks.TODOs{}
	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	task := args[0]

	show := false
	if args[1] != "" {
		show = true
	}
	
	completeID := 0
	if args[2] != "" {
		id := args[2]
		val, err := strconv.Atoi(id)
		if err != nil || val <= 0 {
			fmt.Fprintln(os.Stderr, RED("ID should be a positive integer!"))
			os.Exit(1)
		}
		completeID = val
	}
	
	deleteID := 0
	if args[3] != "" {
		id := args[3]
		val, err := strconv.Atoi(id)
		if err != nil || val <= 0 {
			fmt.Fprintln(os.Stderr, RED("ID should be a positive integer!"))
			os.Exit(1)
		}
		deleteID = val
	}
	
	switch {
	case task != "":
		// add a task here
		println("Add task: ", task)
		todos.Add(task)
	
	case show:
		// list all tasks here
		todos.List()
	
	case completeID != 0:
		// complete all your tasks here
		println("Complete ID: ", completeID)

		if err := todos.Complete(completeID); err != nil {
			fmt.Println(RED("Could not complete task!"))
			os.Exit(1)
		}
	
	case deleteID != 0:
		// delete your tasks here
		println("Delete ID: ", deleteID)

		if err := todos.Delete(deleteID); err != nil {
			fmt.Println(RED("Could not delete task!"))
			os.Exit(1)
		}
	
	default:
		fmt.Fprintln(os.Stdout, RED("invalid command"))
		os.Exit(1)
	}
	
	if err := todos.Store(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
