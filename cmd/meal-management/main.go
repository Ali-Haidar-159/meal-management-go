package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"ali.com/playground/internal/cli"
	"ali.com/playground/internal/domain"
	"ali.com/playground/internal/service"
	"github.com/joho/godotenv"
)

func main() {

	if err := loadConfig(); err != nil {
		fmt.Println("Error loading configuration:", err)
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
	go cli.HomePageShow(reader, chanelHome)
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

		person := domain.Person{
			Id:         service.CreateId(),
			UserName:   strings.TrimSpace(username),
			Password:   strings.TrimSpace(password),
			Department: strings.TrimSpace(department),
		}
		go service.WriteFile(os.Getenv("USERS_FILE"), person, chanelErrWriteFunc)

		err1 := <-chanelErrWriteFunc
		if err1 != nil {
			fmt.Println("Find error to write data : ", err1)
		} else {
			fmt.Println("Data Stored Successfully")
		}

		goto home

	case "2":
		go service.Login(chanelLoginFunc)
		errLogin := <-chanelLoginFunc

		if errLogin != nil {
			fmt.Println(errLogin)
		} else {
			fmt.Println("Login Successful")
		dashboard:
			go cli.UserDashboard(reader, chanelDashboard)
			choice := <-chanelDashboard
			service.DashboardHandler(choice)
			if choice != "0" {
				goto dashboard
			} else {
				goto home
			}
		}

	case "0":
		fmt.Println("Exit From App Successfully")
		break
	default:
		fmt.Println("Enter a right command!")
		goto home
	}
}

func loadConfig() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	for {
		envPath := filepath.Join(dir, ".env")
		if _, err := os.Stat(envPath); err == nil {
			if err := godotenv.Load(envPath); err != nil {
				return err
			}
			return os.Chdir(dir)
		} else if !os.IsNotExist(err) {
			return err
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return fmt.Errorf(".env file not found in the current directory or any parent directory")
		}
		dir = parent
	}
}
