package main

import (
	"github.com/rubenhoenle/schlingel/api"
)

func main() {
	router := api.BuildRouter()
	router.Run(":8080")
}
