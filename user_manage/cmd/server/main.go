package main

import (
	"user-manage-backend/internal/app"
	"user-manage-backend/internal/configs"
)

func main() {
	config := configs.NewConfig()

	application := app.NewApplication(config)
	if err := application.Run(); err != nil {
		panic(err)
	}

}
