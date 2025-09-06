package main

import (
	"fon/genai"
	"fon/persistence"
	"fon/router"
)

func main() {
	genai.Start()
	persistence.Start()
	router.Start()
}
