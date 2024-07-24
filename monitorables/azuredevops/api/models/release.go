package models

import (
	"time"

	"github.com/Vaelatern/monitoror/models"
)

type (
	Release struct {
		ReleaseNumber  string
		DefinitionName string
		Author         *models.Author

		Status string

		FinishedAt *time.Time
		StartedAt  *time.Time
		QueuedAt   *time.Time
	}
)
