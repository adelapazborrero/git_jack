package service

import (
	"fmt"
	"net/http"

	"github.com/adelapazborrero/git_jack/model"
)

const (
	githubApiUrl           = "https://api.github.com"
	githubValidityEndpoint = "/user"
)

type GithubService struct {
	Token        string
	isTokenValid bool
}

func NewGithubService(token string) GitService {
	return &GithubService{
		Token: token,
	}
}

func (ghs *GithubService) GetTokenValidity() {
	req, err := http.NewRequest(
		http.MethodGet, githubApiUrl+githubValidityEndpoint,
		nil,
	)
	if err != nil {
		ghs.isTokenValid = false
		return
	}

	req.Header.Set("Authorization", "Bearer "+ghs.Token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Could not make request to validate token")
		ghs.isTokenValid = false
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("[-] Token is not valid")
		ghs.isTokenValid = false
		return
	}

	fmt.Println("[+] Token is valid")
	ghs.isTokenValid = true
}

func (ghs *GithubService) GetRepositories() ([]model.Repository, error) {
	models := []model.Repository{
		{
			Name: "Test github",
		},
	}

	return models, nil
}
