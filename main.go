package main

import (
    "go-test/database"
    "go-test/utils"
    "go-test/routes"
)

func main() {
    utils.LoadEnv()	
    database.Connect()
    router := routes.SetupRoutes()

    router.Run("localhost:3333")
}
