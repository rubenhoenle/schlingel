package main

import (
	"github.com/rubenhoenle/schlingel/api"
	"github.com/jgero/schlingel/persistence/inmemory"
)

func main() {
	persistence := inmemory.NewInMemoryPersistence()
	router := api.BuildRouter(persistence)
	router.Run(":8080")
}
