package models

import (
	coreModels "github.com/Vaelatern/monitoror/models"
)

type PullRequest struct {
	ID     int
	Title  string
	Author coreModels.Author

	SourceOwner      string
	SourceRepository string
	SourceBranch     string
	CommitSHA        string
}
