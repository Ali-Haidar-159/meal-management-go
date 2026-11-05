package service

import (
	"encoding/json"
	"os"

	"ali.com/playground/internal/domain"
)

func GetAllUsers(file string) ([]domain.Person, error) {

	data, err := os.ReadFile(file)

	if err != nil {
		return nil, err
	}

	var users []domain.Person

	err = json.Unmarshal(data, &users)

	if err != nil {
		return nil, err
	}

	return users, nil
}
