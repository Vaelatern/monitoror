package models

import (
	"time"

	"github.com/Vaelatern/monitoror/models"
)

type (
	Build struct {
		Number   string
		FullName string
		Author   *models.Author

		Building  bool
		Result    string
		StartedAt time.Time
		Duration  time.Duration
	}
)
