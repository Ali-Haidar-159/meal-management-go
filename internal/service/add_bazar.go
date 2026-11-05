package service

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func AddBazar(userID string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the bazar costs (100 25 80) :")
	bazar, err := reader.ReadString('\n')
	bazar = strings.TrimSpace(bazar)

	if err != nil {
		fmt.Println("Find error to get bazar input : ", err)
		return
	} else {
		eachBazar := strings.Split(bazar, " ")

		totalBazarCost := 0
		for _, item := range eachBazar {
			itemInt, _ := strconv.Atoi(item)
			totalBazarCost += itemInt
		}

		fileName := fmt.Sprintf("%s_bazar.json", time.Now().Month())
		_, errFile := CreateFile(fileName)

		if errFile != nil {
			fmt.Println("Find error to create the bazar file", errFile)
			return
		} else {
			var found int = 0
			foundObj := make(map[string]int)
			// create new bazar data
			bazars, _ := GetBazars(fileName)
			for i, item := range bazars {
				_, ok := item[userID]
				if ok == true {
					found = 1
					foundObj = item
					bazars = append(bazars[:i], bazars[i+1:]...)
				}
			}

			if found == 1 {
				totalBazarCost = foundObj[userID] + totalBazarCost
				foundObj[userID] = totalBazarCost
				bazars = append(bazars, foundObj)

			} else {
				data := map[string]int{
					userID: totalBazarCost,
				}
				bazars = append(bazars, data)

			}

			if len(bazars) == 0 {
				data := map[string]int{
					userID: totalBazarCost,
				}
				bazars = append(bazars, data)
			}

			jsonData, _ := json.MarshalIndent(bazars, "", "  ")

			err := os.WriteFile(fileName, jsonData, 0644)
			if err != nil {
				fmt.Println("Find error to write file")
				return
			} else {
				fmt.Println("Bazar data stored successfully")
			}

		}

	}
}
