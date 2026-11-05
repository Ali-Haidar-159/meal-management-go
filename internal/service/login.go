package service

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Login(chanelLoginFunc chan error) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\n\nEnter your user name : ")
	un, _ := reader.ReadString('\n')

	fmt.Print("Enter your password : ")
	pass, _ := reader.ReadString('\n')

	userName := strings.TrimSpace(un)
	password := strings.TrimSpace(pass)

	allUser, err := GetAllUsers(os.Getenv("USERS_FILE"))

	if err != nil {
		chanelLoginFunc <- err
		return
	}

	userFound := false

	for _, user := range allUser {

		if user.UserName == userName {

			userFound = true

			if user.Password == password {

				file, errFileCreate := CreateFile(os.Getenv("LOGIN_FILE"))

				if errFileCreate != nil {
					chanelLoginFunc <- errFileCreate
					return
				}

				os.WriteFile(file.Name(), []byte(user.Id), 0644)

				if err != nil {
					chanelLoginFunc <- err
					return
				}

				file.Close()

				chanelLoginFunc <- nil
				return
			}

			chanelLoginFunc <- fmt.Errorf("password is incorrect")
			return
		}
	}

	if !userFound {
		chanelLoginFunc <- fmt.Errorf("username not found")
	}
}
