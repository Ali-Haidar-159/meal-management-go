package modules

import (
	"encoding/json"
	"os"

	"ali.com/playground/utils"
)

func WriteFile(file string, data utils.Person, chanelErrWriteFunc chan error) {
	users, _ := GetAllUsers(os.Getenv("USERS_FILE"))
	users = append(users, data)

	jsonData, errMarshal := json.MarshalIndent(users, "", "  ")

	if errMarshal != nil {
		chanelErrWriteFunc <- errMarshal
	} else {
		err := os.WriteFile(file, jsonData, 0644)

		if err != nil {
			chanelErrWriteFunc <- err
		} else {
			chanelErrWriteFunc <- nil
		}
	}

}
