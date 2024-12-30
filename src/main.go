package main

import (
	"blip-fullstack.com/test/src/models"
	"blip-fullstack.com/test/src/parsers"
	"blip-fullstack.com/test/src/scorer"
	"fmt"
	"os"
	"sort"
)

func main() {
	file, errOpenFile := os.Open("commits.csv")
	if errOpenFile != nil {
		fmt.Println("Error opening file:", errOpenFile)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	commits, errParseCSV := parsers.ParseCSV(file)
	if errParseCSV != nil {
		fmt.Println("Error reading CSV:", errParseCSV)
		return
	}
	repoScores := scorer.CalculateScores(commits)

	printTopRepositories(repoScores, 10)
}

func printTopRepositories(scores map[string]models.RepositoryScore, topRepos int) {
	repoList := make([]models.RepositoryScore, 0, len(scores))
	for _, score := range scores {
		repoList = append(repoList, score)
	}

	sort.SliceStable(repoList, func(i, j int) bool {
		return repoList[i].Score > repoList[j].Score
	})

	fmt.Println("Top Repositories:")
	for i, repo := range repoList {
		if i >= topRepos {
			break
		}
		fmt.Printf("%d. %s - Score: %.f, Unique Contributors: %d\n", i+1, repo.Repo, repo.Score, len(repo.Contributors))
	}
}
