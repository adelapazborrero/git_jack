package menuoption

import (
	"fmt"

	"github.com/adelapazborrero/git_jack/service"
)

func GetRepositoryList(gitService service.GitService) {
	repos, err := gitService.GetRepositories()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(repos)
}
