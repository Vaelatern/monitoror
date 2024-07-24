package models

import (
	coreModels "github.com/Vaelatern/monitoror/models"
)

type MergeRequest struct {
	ID     int
	Title  string
	Author coreModels.Author

	SourceProjectID int
	SourceBranch    string
	CommitSHA       string
}
