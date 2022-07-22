package main

import (
	"GoMongo/Routers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		return
	}
	Routers.Start()
}
