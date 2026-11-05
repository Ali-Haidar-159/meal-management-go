package service

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func AddMeal(userID string) {
	var todayMeal int
	fmt.Print("Enter the meal for today :")
	_, err := fmt.Scan(&todayMeal)

	if err != nil {
		fmt.Println("Find error to get todays meal input : ", err)
		return
	} else {
		var totalMeal int = 0

		fileName := fmt.Sprintf("%s_meal.json", time.Now().Month())
		_, errFile := CreateFile(fileName)

		if errFile != nil {
			fmt.Println("Find error to create the meal file", errFile)
			return
		} else {
			var found int = 0
			foundObj := make(map[string]int)

			// create new bazar data
			meals, _ := GetBazars(fileName)
			for i, item := range meals {
				_, ok := item[userID]
				if ok == true {
					found = 1
					foundObj = item
					meals = append(meals[:i], meals[i+1:]...)
				}
			}

			if found == 1 {
				totalMeal = foundObj[userID] + todayMeal
				foundObj[userID] = totalMeal
				meals = append(meals, foundObj)

			} else {
				data := map[string]int{
					"date": time.Now().Day(),
					userID: todayMeal,
				}
				meals = append(meals, data)

			}

			if len(meals) == 0 {
				data := map[string]int{
					userID: totalMeal,
				}
				meals = append(meals, data)
			}

			jsonData, _ := json.MarshalIndent(meals, "", "  ")

			err := os.WriteFile(fileName, jsonData, 0644)
			if err != nil {
				fmt.Println("Find error to write file")
				return
			} else {
				fmt.Println("Meals data stored successfully")
			}

		}

	}
}
