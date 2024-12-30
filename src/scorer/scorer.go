package scorer

import (
	"blip-fullstack.com/test/src/models"
	"math"
	"time"
)

// CalculateScores calculates the score, per repository, for a given list of commits.
func CalculateScores(commits []models.Commit) map[string]models.RepositoryScore {
	repoStats := make(map[string]*models.RepositoryScore)

	for _, commit := range commits {
		if repoStats[commit.Repo] == nil {
			repoStats[commit.Repo] = &models.RepositoryScore{Repo: commit.Repo, Contributors: make(map[string]struct{})}
		}

		ageFactor := calculateAgeFactor(commit)
		fileFactor := calculateFileFactor(commit)
		lineFactor := calculateLineFactor(commit)
		commitScore := ageFactor * fileFactor * lineFactor

		repo := repoStats[commit.Repo]
		repo.Score += commitScore
		if commit.User != "" {
			repo.Contributors[commit.User] = struct{}{}
		}
	}

	finalScores := make(map[string]models.RepositoryScore)
	for repo, stats := range repoStats {
		finalScores[repo] = models.RepositoryScore{
			Repo:         repo,
			Score:        stats.Score * calculateContributorFactor(stats.Contributors),
			Contributors: stats.Contributors,
		}
	}

	return finalScores
}

func calculateAgeFactor(commit models.Commit) float64 {
	decayRate := 0.00000001
	return math.Exp(float64(commit.Timestamp-time.Now().Unix()) * decayRate)
}

func calculateFileFactor(commit models.Commit) float64 {
	increaseFactor := 1.0
	return 1.0 + math.Log(increaseFactor+float64(commit.Files))
}

func calculateLineFactor(commit models.Commit) float64 {
	additionFactor := 1.0
	deletionFactor := 1.2
	return 1.0 +
		additionFactor*math.Log(1+float64(commit.Additions)) +
		deletionFactor*math.Log(1+float64(commit.Deletions))
}

func calculateContributorFactor(contributors map[string]struct{}) float64 {
	return 1.0 + math.Log(1+float64(len(contributors)))
}
