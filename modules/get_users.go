package modules

import (
	"encoding/json"
	"os"

	"ali.com/playground/utils"
)

func GetAllUsers(file string) ([]utils.Person, error) {

	data, err := os.ReadFile(file)

	if err != nil {
		return nil, err
	}

	var users []utils.Person

	err = json.Unmarshal(data, &users)

	if err != nil {
		return nil, err
	}

	return users, nil
}
