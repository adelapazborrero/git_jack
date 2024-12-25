package main

import (
	"flag"

	"github.com/adelapazborrero/git_jack/service"
)

func main() {
	token := flag.String("token", "value", "usage")
	flag.Parse()

	gitService := service.NewgitService(*token)
	gitService.GetTokenValidity()
}
