package modules

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
			// fmt.Println(itemInt)
			totalBazarCost += itemInt
		}
		// fmt.Print("Total bazar is : ", totalBazarCost)

		fileName := fmt.Sprintf("%s_bazar.json", time.Now().Month())
		_, errFile := CreateFile(fileName)

		if errFile != nil {
			fmt.Println("Find error to create the bazar file", errFile)
			return
		} else {
			data := map[string]int{
				userID: totalBazarCost,
			}

			jsonData, _ := json.MarshalIndent(data, "", "  ")

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
