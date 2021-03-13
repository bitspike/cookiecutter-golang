package main

import (
	"{{cookiecutter.app_name}}/cmd"
	"{{cookiecutter.app_name}}/config"
	"{{cookiecutter.app_name}}/log"
)

func main() {
	config.Initialize()
	log.Initialize()
	cmd.NewManager().Execute()
}
