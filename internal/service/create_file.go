package service

import (
	"fmt"
	"os"
)

func CreateFile(fileName string) (*os.File, error) {

	if fileName != "" {

		file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		defer file.Close()

		if err != nil {
			return nil, err
		} else {
			return file, nil
		}

	} else {
		return nil, fmt.Errorf("File name is required")
	}

}
