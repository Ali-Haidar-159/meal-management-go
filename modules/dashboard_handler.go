package modules

import (
	"fmt"
	"os"
)

func DashboardHandler(choice string) {
	switch choice {
	case "1":

		userId, err := os.ReadFile(os.Getenv("LOGIN_FILE"))

		if err != nil {
			fmt.Println("find error to read the login id")
		} else {
			AddBazar(string(userId))
		}

	case "2":
		AddMeal()
	}

}
