package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func HomePageShow(reader *bufio.Reader, chanelHome chan string) {
	fmt.Printf("\n===== WELCOME TO %s =====\n", os.Getenv("APP_NAME"))

	fmt.Println("1. Signup ")
	fmt.Println("2. Login ")
	fmt.Println("0. Exit From Program")

	fmt.Print("\nEnter user choice : ")
	goal, _ := reader.ReadString('\n')

	choice := strings.TrimSpace(goal)

	chanelHome <- choice
}

func UserDashboard(reader *bufio.Reader, chanelDashboard chan string) {
	fmt.Printf("\n===== WELCOME TO %s =====\n", os.Getenv("APP_NAME"))

	fmt.Println("1. Add Bazaar")
	fmt.Println("2. Add Meal Count")
	fmt.Println("0. Logout")

	fmt.Print("\nEnter user choice : ")
	goal, _ := reader.ReadString('\n')

	choice := strings.TrimSpace(goal)

	chanelDashboard <- choice

}
