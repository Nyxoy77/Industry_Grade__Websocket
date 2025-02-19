package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("./.env"); err != nil {
		fmt.Printf("error occured loading the env %v", err)
		return
	}
	// fmt.Println(os.Getenv("SECRET_KEY"))
}
