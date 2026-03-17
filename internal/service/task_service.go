package service

import (
	"Actium_Todo/internal/cli"
	"Actium_Todo/internal/models"
	"Actium_Todo/internal/repository"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var t = models.Task{}
var reader = bufio.NewReader(os.Stdin)

func CheckTheStatus(status string) bool {
	return models.ValidStatuses[status]
}

func SearchTasks(choice, creatorId int, userN string) {
	switch choice {
	case 1:
		fmt.Println("Enter the Task's ID")
		var taskId int
		fmt.Scan(&taskId)
		fmt.Printf("Loading all tasks %s\n", userN)
		time.Sleep(1 * time.Second)
		foundTasks, err := repository.SearchTask_byId(creatorId, taskId)
		if err != nil {
			log.Fatal(err)
		}
		if len(foundTasks) == 0 {
			fmt.Println("No tasks found with the given ID.")
		} else {
			for _, t := range foundTasks {
				fmt.Printf("ID: %d | Title: %s | Description: %s | Status: %s | Created at: %s | Deadline: %s\n",
					t.Id, t.Title, t.Description, t.Status, t.CreatedAt.Format("2006-01-02"), t.Deadline)
			}
		}
	case 2:
		fmt.Println("Enter the Task's Title")
		taskTitle, _ := reader.ReadString('\n')
		taskTitle = strings.TrimSpace(taskTitle)
		fmt.Printf("Loading all tasks %s\n", userN)
		time.Sleep(1 * time.Second)
		foundTasks, err := repository.SearchTask_byTitle(creatorId, taskTitle)
		if err != nil {
			log.Fatal(err)
		}
		if len(foundTasks) == 0 {
			fmt.Println("No tasks found with the given Title.")
		} else {
			for _, t := range foundTasks {
				fmt.Printf("ID: %d | Title: %s | Description: %s | Status: %s | Created at: %s | Deadline: %s\n",
					t.Id, t.Title, t.Description, t.Status, t.CreatedAt.Format("2006-01-02"), t.Deadline)
			}
		}
	case 3:
		fmt.Println("Enter the Task's Status")
		taskStatus, _ := reader.ReadString('\n')
		taskStatus = strings.TrimSpace(taskStatus)
		fmt.Printf("Loading all tasks %s\n", userN)
		time.Sleep(1 * time.Second)
		foundTasks, err := repository.SearchTask_byStatus(creatorId, taskStatus)
		if err != nil {
			log.Fatal(err)
		}
		if len(foundTasks) == 0 {
			fmt.Println("No tasks found with the given Status.")
		} else {
			for _, t := range foundTasks {
				fmt.Printf("ID: %d | Title: %s | Description: %s | Status: %s | Created at: %s | Deadline: %s\n",
					t.Id, t.Title, t.Description, t.Status, t.CreatedAt.Format("2006-01-02"), t.Deadline)
			}
		}
	case 4:
		return
	default:
		fmt.Println("Invalid option")
	}
}

func Run() {

	for {
		var choice int16
		input := cli.SignIn_menu(choice)
		switch input {
		case 1: //Sign Up
			fmt.Println("Enter Your username")
			userN, _ := reader.ReadString('\n')
			userN = strings.TrimSpace(userN)
			fmt.Println("Enter Your password")
			userP, _ := reader.ReadString('\n')
			userP = strings.TrimSpace(userP)
			if userN != "" && userP != "" {
				SignUp(userN, userP)
			} else {
				fmt.Println("Username and password cannot be empty. Please try again.")
			}
		case 2: //Login
			fmt.Println("Enter Your username")
			userN, _ := reader.ReadString('\n')
			userN = strings.TrimSpace(userN)
			fmt.Println("Enter Your password")
			userP, _ := reader.ReadString('\n')
			userP = strings.TrimSpace(userP)
			userId, ok := Login(userN, userP)
			if ok {
				fmt.Printf("\n\nSuccessfull login\nAccess allowed\nWelcome back %s\n\n", userN)
				creatorId := userId
				var choice int
			MenuLoop:
				for {
					m := cli.TodoMenu(int16(choice))
					switch m {
					case 1: //Add Task
						var (
							title       string
							description string
							status      string
							deadline    string
						)

						fmt.Println("Enter the task's title")
						title, _ = reader.ReadString('\n')
						title = strings.TrimSpace(title)
						fmt.Println("Enter the task's Description")
						description, _ = reader.ReadString('\n')
						description = strings.TrimSpace(description)
						fmt.Println("\nEnter the task's Status\nWrite the first letter in uppercase\n")
						fmt.Println("Use the tags given down below\n\n'New' for the new tasks awaiting to be accomplished\n'In progress' for the tasks in process\n'Completed' for the completed tasks\n'Canceled' for the canceled tasks")

						for { //Checking the given status by user

							status, _ = reader.ReadString('\n')
							status = strings.TrimSpace(status)
							if !CheckTheStatus(status) {
								fmt.Println("\nUse the given examples of status")
								fmt.Println("'New' for the new tasks awaiting to be accomplished\n'In progress' for the tasks in process\n'Completed' for the completed tasks\n'Canceled' for the canceled tasks")
								continue
							}
							break
						}

						fmt.Println("Enter the task's Deadline")
						fmt.Println("Use the format givven bellow:\nYYYY-MM-DD")
						deadline, _ = reader.ReadString('\n')
						deadline = strings.TrimSpace(deadline)
						t.AddTask(title, description, status, deadline, creatorId)

						err := repository.CreateTask(title, description, status, creatorId, deadline)
						if err != nil {
							log.Fatal(err)
						}
						fmt.Printf("\nTask has been created.\nTitle: %s\nDescription: %s.\nStatus: %s.\nDeadline: %s\n", title, description, status, deadline)

					case 2: // Show all tasks
						fmt.Printf("Loading all tasks %s\n", userN)
						time.Sleep(1 * time.Second)
						gettasks, err := repository.GetTasksByCreator(creatorId)
						if err != nil {
							log.Fatal(err)
						}
						if len(gettasks) == 0 {
							fmt.Println("No tasks have been created!")
						} else {
							//otprint with local numbering for user
							for i, t := range gettasks {
								fmt.Printf("Task #%d | ID: %d | Title: %s | Description: %s | Status: %s | Created at: %s | Deadline: %s\n",
									i+1, t.Id, t.Title, t.Description, t.Status, t.CreatedAt.Format("2006-01-02"), t.Deadline)
							}
						}

					case 3: // Show all completed tasks
						fmt.Printf("Loading completed tasks %s\n", userN)
						time.Sleep(1 * time.Second)
						tasks, err := repository.ShowCompletedTasks(creatorId)
						if err != nil {
							log.Fatal(err)
						}
						if len(tasks) == 0 {
							fmt.Println("No completed tasks have been created!")
						} else {
							//otprint with local numbering for user
							for i, t := range tasks {
								fmt.Printf("Task #%d | ID: %d | Title: %s | Description: %s | Status: %s | Created at: %s | Deadline: %s\n",
									i+1, t.Id, t.Title, t.Description, t.Status, t.CreatedAt.Format("2006-01-02"), t.Deadline)
							}
						}
					case 4: //Change the task's status
						tasks, err := repository.LoadAllTasks(creatorId)
						if err != nil {
							log.Fatal(err)
						}
						if len(tasks) == 0 {
							fmt.Println("No tasks have been created!")
							continue
						}
						fmt.Println("Enter the id of the Task")
						var taskId int
						fmt.Scan(&taskId)

						found := false
						for _, t := range tasks {
							if taskId == t.Id {
								found = true
								break
							}

						}
						if !found {
							fmt.Println("The task by the given ID has not been found ")
							continue
						}
						fmt.Println("Set new status for the task")
						for { //Checking the given status by user
							fmt.Println("Use the tags given down below\n\n'New' for the new tasks awaiting to be accomplished\n'In progress' for the tasks in process\n'Completed' for the completed tasks\n'Canceled' for the canceled tasks")
							newStatus, _ := reader.ReadString('\n')
							newStatus = strings.TrimSpace(newStatus)
							if !CheckTheStatus(newStatus) {
								fmt.Println("\nUse the given examples of status")
								fmt.Println("'New' for the new tasks awaiting to be accomplished\n'In progress' for the tasks in process\n'Completed' for the completed tasks\n'Canceled' for the canceled tasks")
							}
							repository.UpdateStatus(taskId, newStatus, creatorId)
							break
						}
					case 5: //Delete task
						tasks, err := repository.LoadAllTasks(creatorId)
						if err != nil {
							log.Fatal(err)
						}
						if len(tasks) == 0 {
							fmt.Println("No tasks have been created!")
							continue
						}
						fmt.Println("Enter the id of the Task")
						var taskId int
						fmt.Scan(&taskId)

						found := false
						for _, t := range tasks {
							if taskId == t.Id {
								found = true
								break
							}

						}
						if !found {
							fmt.Println("The task by the given ID has not been found ")
							continue
						}
						repository.DeleteTask(taskId, creatorId)
						fmt.Printf("The task by ID %d has been deleted successfully", taskId)

					case 6: //Search tasks
						tasks, err := repository.LoadAllTasks(creatorId)
						if err != nil {
							log.Fatal(err)
						}
						if len(tasks) == 0 {
							fmt.Println("No tasks have been created!")
							continue
						}

						fmt.Println("Choose the parameter to search with")
						fmt.Println("1. Search tasks by ID\n2. Search tasks by Title\n3. Search tasks by Status\n4. Exit the search menu")
						var searchChoice int
						fmt.Scan(&searchChoice)
						SearchTasks(searchChoice, creatorId, userN)
					case 7: //Log Out
						fmt.Println("Saving all settings")
						time.Sleep(1 * time.Second)
						fmt.Println("Logging out...\n")
						time.Sleep(1 * time.Second)
						break MenuLoop
					case 8: //Exit the programm
						fmt.Println("\nExiting the program...")
						time.Sleep(2 * time.Second)
						fmt.Println("Have a productive day! Good bye.")
						return
					default: //Default text for the wrong option
						log.Println("Invalid option! Try again")
					}
				}
			} else {
				fmt.Println("Incorrect username or Password")
			}

		case 3: //Delete account
			fmt.Println("Enter Your username")
			userN, _ := reader.ReadString('\n')
			userN = strings.TrimSpace(userN)
			fmt.Println("Enter Your password")
			userP, _ := reader.ReadString('\n')
			userP = strings.TrimSpace(userP)
			err := repository.DeleteAllTasksFromUser(userN)
			if err != nil {
				log.Fatal(err)
			}
			err = DeleteMyAccount(userN, userP)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("\nAccount with username %s has been deleted\n\n", userN)
		case 4: //Exit
			fmt.Println("\nExiting the program...")
			time.Sleep(2 * time.Second)
			fmt.Println("Have a productive day! Good bye.")
			return
		default:
			fmt.Println("Invalid option\n")

		}
	}
}
