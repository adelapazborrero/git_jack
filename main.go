package main

import (
	"flag"

	"github.com/adelapazborrero/git_jack/service"
	"github.com/adelapazborrero/git_jack/util"
)

func main() {
	token := flag.String("token", "value", "usage")
	flag.Parse()

	gitService := service.NewgitService(*token)

	repos := gitService.GetRepositories()
	util.PrintRepositoryList(repos)
}
