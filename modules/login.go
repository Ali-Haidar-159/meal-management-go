package modules

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
		fmt.Println("Find error to get all users.Error is : ", err)
		chanelLoginFunc <- err
	} else {
		for _, user := range allUser {
			// fmt.Println("user name : ", user.UserName)
			if user.UserName == userName {
				if user.Password == password {

					file, errFileCreate := CreateFile(os.Getenv("LOGIN_FILE"))

					if errFileCreate != nil {
						chanelLoginFunc <- errFileCreate
					} else {
						os.WriteFile(file.Name(), []byte(user.Id), 0644)
						chanelLoginFunc <- nil
					}
				} else {
					chanelLoginFunc <- fmt.Errorf("Auth Failed.")
				}
			} else {
				chanelLoginFunc <- fmt.Errorf("Auth Failed.")
			}
		}
	}

}
