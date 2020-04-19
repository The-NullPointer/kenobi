package main

import (
	"kenobi/cmd"
	"kenobi/config"
)

func main() {
	myConf := config.Config{}
	myConf["SecretKey"] = "123"
	cmd.Execute(&myConf)
}
