package cli

import "fmt"

func TodoMenu(choice int16) int {
	fmt.Println("\n       <----Menu---->\n")
	fmt.Println("Shoose options using number(1-0)")
	fmt.Println("1. **Create task**\n2. **Show my tasks**\n3. **Show completed tasks**\n4. **Update the task's status**\n5. **Delete task**\n6. **Search tasks**\n7. **Log out**\n8. **Exit program**")
	fmt.Scan(&choice)
	return int(choice)
}

func SignIn_menu(choice int16) int {
	fmt.Println("Sign up if You are new here\nLogin if You already have been registrated")
	fmt.Println("Print the option using numbers 1-2\n1. Sign up\n2. Login\n3. Delete my account\n4. Exit")
	fmt.Scan(&choice)
	return int(choice)
}
