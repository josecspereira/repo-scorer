package scorer

import (
	"blip-fullstack.com/test/src/models"
)

// CalculateScores calculates the score, per repository, for a given list of commits.
func CalculateScores(commits []models.Commit) map[string]models.RepositoryScore {
	repoStats := make(map[string]*models.RepositoryScore)

	for _, commit := range commits {
		if repoStats[commit.Repo] == nil {
			repoStats[commit.Repo] = &models.RepositoryScore{Repo: commit.Repo, Contributors: make(map[string]struct{}), Score: 1}
		}

	}

	finalScores := make(map[string]models.RepositoryScore)
	for repo, stats := range repoStats {
		finalScores[repo] = models.RepositoryScore{
			Repo:         repo,
			Score:        stats.Score,
			Contributors: stats.Contributors,
		}
	}

	return finalScores
}
