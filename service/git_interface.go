package service

import (
	"fmt"
	"os"
	"strings"

	"github.com/adelapazborrero/git_jack/model"
)

type GitService interface {
	GetTokenValidity()
	GetRepositories() []model.Repository
}

func NewgitService(token string) GitService {
	isGitlab := strings.HasPrefix(token, "glpat_")
	isGithub := strings.HasPrefix(token, "ghp_")

	if !isGitlab && !isGithub {
		fmt.Println("Not a valid token format")
		os.Exit(0)
	}

	var gitService GitService

	if isGitlab {
		gitService = NewGitlabService(token)
	}
	if isGithub {
		gitService = NewGithubService(token)
	}

	return gitService

}
