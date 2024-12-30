package parsers

import (
	"blip-fullstack.com/test/src/models"
	"encoding/csv"
	"io"
	"strconv"
)

// ParseCSV Parses the CSV of a given reader with the format timestamp,username,repository,files,additions,deletions.
func ParseCSV(reader io.Reader) ([]models.Commit, error) {
	csvReader := csv.NewReader(reader)
	rows, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	var commits []models.Commit
	for _, row := range rows[1:] {
		timestamp, _ := strconv.ParseInt(row[0], 10, 64)
		additions, _ := strconv.Atoi(row[4])
		deletions, _ := strconv.Atoi(row[5])
		files, _ := strconv.Atoi(row[3])
		commits = append(commits, models.Commit{
			Timestamp: timestamp,
			User:      row[1],
			Repo:      row[2],
			Files:     files,
			Additions: additions,
			Deletions: deletions,
		})
	}

	return commits, nil
}
