package main

import (
	"fon/persistence"
	"fon/router"
)

func main() {
	persistence.Start()
	router.Start()
}
