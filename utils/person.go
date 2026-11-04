package utils

type Person struct {
	Id         string `json:"_id"`
	UserName   string `json:"username"`
	Password   string `json:"password"`
	Department string `json:"department"`
}
