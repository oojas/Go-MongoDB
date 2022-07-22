package main

import (
	"GoMongo/Routers"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	Routers.Start()
}
