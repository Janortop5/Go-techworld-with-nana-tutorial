package main

import "fmt"

//	var shortGolang =  "Watch Go crash course"
//	var FullGolang  = "Watch Nana's Golang Full Course"
//	var rewardDessert = "Reward myself with a donut"

	//rewardDessert := "Reward myself with a donut"
	// var maxItemsInGroup = 20

//	var taskItems = [3]string {shortGolang, FullGolang, rewardDessert }

func main() {

		fmt.Println("####### Welcome to our Todolist App! #######")

		var shortGolang =  "Watch Go crash course"
		var FullGolang  = "Watch Nana's Golang Full Course"
		var rewardDessert = "Reward myself with a donut"
	// rewardDessert := "Reward myself with a donut"
  //  var maxItemsInGroup = 10

		var taskItems = []string {shortGolang, FullGolang, rewardDessert }

		printTasks(taskItems)
		println()
		addTask(taskItems, "Go for a run")
		println()
		printTasks(taskItems)
		// fmt.Println(shortGolang)
		// fmt.Println(FullGolang)
		//fmt.Println(rewardDessert)
		// fmt.Println(5)
		//fmt.Println(999999)

		// fmt.Println()
		// fmt.Println("Tutorials")
		// fmt.Println(shortGolang)
		// fmt.Println(FullGolang)

		// fmt.Println()
		// fmt.Println("Rewards") 
		// fmt.Println(rewardDessert)

		// fmt.Println()
		// fmt.Println("My Project")
		// fmt.Println(FullGolang)
}

func printTasks(taskItems []string,) {
		fmt.Println("list of my Todos")
		// fmt.Println("Tasks:", taskItems)
		for index, task := range taskItems {
			  // fmt.Println(index+1, ".", task)
				// fmt.Printf("%d; %s\n", index+1, task)
				// fmt.Printf("%d. %s\n", index+1, task)
        fmt.Printf("%d: %s\n", index+1, task)

		}
}

func addTask(taskItem []string, newTask string) {
		var updatedTaskItems = append(taskItem, newTask)
		printTasks(updatedTaskItems)
		
}