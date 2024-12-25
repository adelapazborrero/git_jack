package util

import (
	"fmt"

	"github.com/adelapazborrero/git_jack/model"
)

func PrintRepository(repo model.Repository) {
	fmt.Println(repo)
}

func PrintRepositoryList(repos []model.Repository) {
	fmt.Println(repos)
}
