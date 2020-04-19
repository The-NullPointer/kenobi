package main

import (
	"kenobi/cmd"
	"kenobi/config"
)

func main() {
	myConf := config.Config{
		"SecretKey": "123",
	}

	cmd.Execute(myConf)
}
