package service

import (
	"math/rand"
	"strconv"
	"time"
)

func CreateId() (id string) {
	currentTime := time.Now().Nanosecond()
	randomSalt := rand.Intn(9999)

	id = strconv.Itoa(currentTime + randomSalt)

	return

}
