package models

import (
	"time"

	"github.com/Vaelatern/monitoror/models"
)

type (
	Build struct {
		ID     uint
		Branch string
		Author models.Author

		State      string // see https://github.com/shuheiktgw/go-travis/blob/master/builds.go#L116
		StartedAt  time.Time
		FinishedAt time.Time
		Duration   time.Duration
	}
)
