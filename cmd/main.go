package main

import (
	"products-service/internal/app"
	_ "products-service/internal/app/ent/runtime"
	"products-service/internal/config/env"
)

func main() {
	if err := env.ValidateEnv(); err != nil {
		panic("error validating env: " + err.Error())
	}
	port := env.PORT
	app.New(port)
}
