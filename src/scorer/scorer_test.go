package scorer_test

import (
	"blip-fullstack.com/test/src/models"
	"blip-fullstack.com/test/src/scorer"
	"math"
	"testing"
	"time"
)

func TestCalculateScores(t *testing.T) {
	now := time.Now().Unix()
	commits := []models.Commit{
		{Timestamp: now - 86400, User: "user1", Repo: "repo1", Files: 3, Additions: 10, Deletions: 5},
		{Timestamp: now - 172800, User: "user2", Repo: "repo1", Files: 1, Additions: 5, Deletions: 10},
		{Timestamp: now - 86400, User: "user1", Repo: "repo2", Files: 5, Additions: 20, Deletions: 0},
		{Timestamp: now - 259200, User: "user3", Repo: "repo2", Files: 2, Additions: 10, Deletions: 5},
	}

	scores := scorer.CalculateScores(commits)

	expected := map[string]models.RepositoryScore{
		"repo1": {
			Repo:         "repo1",
			Score:        47.87,
			Contributors: map[string]struct{}{"user1": {}, "user2": {}},
		},
		"repo2": {
			Repo:         "repo2",
			Score:        48.05,
			Contributors: map[string]struct{}{"user1": {}, "user3": {}},
		},
	}
	for repo, expectedScore := range expected {
		actualScore, exists := scores[repo]
		if !exists {
			t.Errorf("Expected repository %s to be present", repo)
			continue
		}

		if math.Abs(actualScore.Score-expectedScore.Score) > 0.1 {
			t.Errorf("For repo %s: expected score %.2f, got %.2f", repo, expectedScore.Score, actualScore.Score)
		}

		if len(actualScore.Contributors) != len(expectedScore.Contributors) {
			t.Errorf("For repo %s: expected %d contributors, got %d", repo, len(expectedScore.Contributors), len(actualScore.Contributors))
		}
	}
}

func TestPrivilegeDeletionOverAddition(t *testing.T) {
	now := time.Now().Unix()
	commits := []models.Commit{
		{Timestamp: now - 86400, User: "user1", Repo: "repo1", Files: 3, Additions: 5, Deletions: 10},
		{Timestamp: now - 86400, User: "user1", Repo: "repo2", Files: 3, Additions: 10, Deletions: 5},
	}

	scores := scorer.CalculateScores(commits)

	if scores["repo1"].Score <= scores["repo2"].Score {
		t.Errorf("Expected repo1's score (%.2f) to be higher than repo2's score (%.2f)", scores["repo1"].Score, scores["repo2"].Score)
	}
}

func TestPrivilegeNewerCommits(t *testing.T) {
	now := time.Now().Unix()
	commits := []models.Commit{
		{Timestamp: now - 42400, User: "user1", Repo: "repo1", Files: 3, Additions: 5, Deletions: 10},
		{Timestamp: now - 86400, User: "user1", Repo: "repo2", Files: 3, Additions: 5, Deletions: 10},
		{Timestamp: now - 1286400, User: "user1", Repo: "repo3", Files: 3, Additions: 5, Deletions: 10},
	}

	scores := scorer.CalculateScores(commits)

	if scores["repo1"].Score <= scores["repo2"].Score {
		t.Errorf("Expected repo1's score (%.2f) to be higher than repo2's score (%.2f)", scores["repo1"].Score, scores["repo2"].Score)
	}

	if scores["repo2"].Score <= scores["repo3"].Score {
		t.Errorf("Expected repo2's score (%.2f) to be higher than repo3's score (%.2f)", scores["repo2"].Score, scores["repo3"].Score)
	}
}

func TestPrivilegeUniqueUsers(t *testing.T) {
	now := time.Now().Unix()
	commits := []models.Commit{
		{Timestamp: now - 86400, User: "user1", Repo: "repo1", Files: 3, Additions: 10, Deletions: 5},
		{Timestamp: now - 172800, User: "user2", Repo: "repo1", Files: 1, Additions: 5, Deletions: 10},
		{Timestamp: now - 86400, User: "user1", Repo: "repo2", Files: 3, Additions: 10, Deletions: 5},
		{Timestamp: now - 172800, User: "user1", Repo: "repo2", Files: 1, Additions: 5, Deletions: 10},
	}

	scores := scorer.CalculateScores(commits)

	if scores["repo1"].Score <= scores["repo2"].Score {
		t.Errorf("Expected repo1's score (%.2f) to be higher than repo2's score (%.2f)", scores["repo1"].Score, scores["repo2"].Score)
	}
}

func TestPrivilegeMoreFilesChanged(t *testing.T) {
	now := time.Now().Unix()
	commits := []models.Commit{
		{Timestamp: now - 86400, User: "user1", Repo: "repo1", Files: 5, Additions: 5, Deletions: 10},
		{Timestamp: now - 86400, User: "user1", Repo: "repo2", Files: 3, Additions: 5, Deletions: 10},
	}

	scores := scorer.CalculateScores(commits)

	if scores["repo1"].Score <= scores["repo2"].Score {
		t.Errorf("Expected repo1's score (%.2f) to be higher than repo2's score (%.2f)", scores["repo1"].Score, scores["repo2"].Score)
	}
}
