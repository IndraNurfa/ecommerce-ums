package main

import (
	"ecommerce-ums/cmd"
	"ecommerce-ums/helpers"
)

func main() {

	// load config
	helpers.SetupConfig()

	// load log
	helpers.SetupLogger()

	// load db
	// helpers.SetupMySQL()

	// load redis
	// helpers.SetupRedis()

	// load kafka
	// cmd.ServeKafka()

	// run http
	cmd.ServeHTTP()
}
