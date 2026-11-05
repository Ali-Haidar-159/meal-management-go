package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"ali.com/playground/modules"
	"ali.com/playground/utils"
	"github.com/joho/godotenv"
)

func main() {

	// CONFIGURATION
	errToLoad := godotenv.Load()
	if errToLoad != nil {
		fmt.Println("Error loading .env file:", errToLoad)
		return
	}

	// DECLARATION PART
	var chanelHome = make(chan string)
	var chanelDashboard = make(chan string)
	var chanelErrWriteFunc = make(chan error)
	var chanelLoginFunc = make(chan error)

	reader := bufio.NewReader(os.Stdin)

	// MAIN WORK
home:
	go utils.HomePageShow(reader, chanelHome)
	goal := <-chanelHome

	switch goal {
	case "1":

		fmt.Print("Enter user username :")
		username, errName := reader.ReadString('\n')
		fmt.Print("Enter your password : ")
		password, errAge := reader.ReadString('\n')
		fmt.Print("Enter user department :")
		department, errDepartment := reader.ReadString('\n')

		if errDepartment != nil || errName != nil || errAge != nil {
			fmt.Println("Find error to input data.Error is : ", errDepartment, errName, errAge)
			return
		}

		person := utils.Person{
			Id:         modules.CreateId(),
			UserName:   strings.TrimSpace(username),
			Password:   strings.TrimSpace(password),
			Department: strings.TrimSpace(department),
		}
		go modules.WriteFile(os.Getenv("USERS_FILE"), person, chanelErrWriteFunc)

		err1 := <-chanelErrWriteFunc
		if err1 != nil {
			fmt.Println("Find error to write data : ", err1)
		} else {
			fmt.Println("Data Stored Successfully")
		}

		goto home

	case "2":
		go modules.Login(chanelLoginFunc)
		errLogin := <-chanelLoginFunc

		if errLogin != nil {
			fmt.Println(errLogin)
		} else {
			fmt.Println("Login Successful")
		dashboard:
			go utils.UserDashboard(reader, chanelDashboard)
			choice := <-chanelDashboard
			modules.DashboardHandler(choice)
			goto dashboard
		}

	case "0":
		fmt.Println("Exit From App Successfully")
		break
	default:
		fmt.Println("Enter a right command!")
		goto home
	}
}
