package service

import (
	"fmt"
	"net/http"

	"github.com/adelapazborrero/git_jack/model"
)

const (
	gitlabApiUrl           = "https://gitlab.com/api/v4"
	gitlabValidityEndpoint = "/personal_access_tokens/self"
)

type GitlabService struct {
	Token        string
	isTokenValid bool
}

func NewGitlabService(token string) GitService {
	return &GitlabService{
		Token: token,
	}
}

func (gls *GitlabService) GetTokenValidity() {
	req, err := http.NewRequest(
		http.MethodGet, gitlabApiUrl+gitlabValidityEndpoint,
		nil,
	)
	if err != nil {
		gls.isTokenValid = false
	}

	req.Header.Set("Authorization", "Bearer "+gls.Token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Could not make request to validate token")
		gls.isTokenValid = false
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("[-] Token is not valid")
		gls.isTokenValid = false
		return
	}

	fmt.Println("[+] Token is valid")
	gls.isTokenValid = true
}

func (gls *GitlabService) GetRepositories() []model.Repository {
	return []model.Repository{
		{
			Name: "Test gitlab",
		},
	}
}
