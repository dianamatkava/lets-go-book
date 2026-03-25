package main

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)


func main() {
	godotenv.Load()
	hostName := os.Getenv("HOST_NANE")
	fmt.Println(hostName)
}