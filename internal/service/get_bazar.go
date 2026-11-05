package service

import (
	"encoding/json"
	"os"
)

func GetBazars(fileName string) ([]map[string]int, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	var bazars []map[string]int

	err = json.Unmarshal(data, &bazars)

	if err != nil {
		return nil, err
	}

	return bazars, nil
}
