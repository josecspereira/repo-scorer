package parsers_test

import (
	"blip-fullstack.com/test/src/models"
	"blip-fullstack.com/test/src/parsers"
	"strings"
	"testing"
)

func TestParseCSV(t *testing.T) {
	sampleCSV := `timestamp,username,repository,files,additions,deletions
1610969774,user0,repo2,5,153,0
1610963057,user0,repo2,2,16,12
1614333792,user1,repo3,1,1,1
`
	reader := strings.NewReader(sampleCSV)

	commits, err := parsers.ParseCSV(reader)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := []models.Commit{
		{Timestamp: 1610969774, User: "user0", Repo: "repo2", Files: 5, Additions: 153, Deletions: 0},
		{Timestamp: 1610963057, User: "user0", Repo: "repo2", Files: 2, Additions: 16, Deletions: 12},
		{Timestamp: 1614333792, User: "user1", Repo: "repo3", Files: 1, Additions: 1, Deletions: 1},
	}
	if len(commits) != len(expected) {
		t.Fatalf("expected %d commits, got %d", len(expected), len(commits))
	}
	for i, commit := range commits {
		if commit != expected[i] {
			t.Errorf("commit %d: expected %+v, got %+v", i, expected[i], commit)
		}
	}
}
