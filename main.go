package main

import (
	"flag"

	"github.com/adelapazborrero/git_jack/menu"
	"github.com/adelapazborrero/git_jack/service"
)

func main() {
	token := flag.String("token", "value", "usage")
	flag.Parse()

	gitService := service.NewGitService(*token)
	gitService.GetTokenValidity()

	menu := menu.Build(gitService)
	menu.Show()
}
