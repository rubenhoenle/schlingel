package main

import (
	"github.com/rubenhoenle/schlingel/api"
	"github.com/rubenhoenle/schlingel/persistence"
)

func main() {
	persistence := persistence.NewOrmPersistence()
	router := api.BuildRouter(persistence)
	router.Run(":8080")
}
