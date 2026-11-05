package service

import (
	"fmt"
	"os"
)

func DashboardHandler(choice string) {

	userId, err := os.ReadFile(os.Getenv("LOGIN_FILE"))
	if err != nil {
		fmt.Println("find error to read the login id")
		return
	}
	switch choice {
	case "1":
		AddBazar(string(userId))

	case "2":
		AddMeal(string(userId))
	case "0":
		fmt.Println("Logout successful")
		break
	default:
		fmt.Println("Enter a correct command !")

	}

}
