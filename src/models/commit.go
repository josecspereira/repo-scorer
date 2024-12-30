package models

type Commit struct {
	Timestamp int64
	User      string
	Repo      string
	Files     int
	Additions int
	Deletions int
}

// RepositoryScore Structure that holds the score and UNIQUE contributors for a given repo.
type RepositoryScore struct {
	Repo         string
	Score        float64
	Contributors map[string]struct{}
}
