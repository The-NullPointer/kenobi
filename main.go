package main

import (
	"kenobi/cmd"
	"kenobi/config"
)

func main() {
	myConf := config.Config{
		"SecretKey":       "123",
		"DatabaseUri":     "host=localhost port=5432 user=admin dbname=zipi password=adminpassword sslmode=disable",
		"DatabaseDialect": "postgres",
		"AllowedOrigins":  []string{"*"},
	}

	cmd.Execute(myConf)

}
