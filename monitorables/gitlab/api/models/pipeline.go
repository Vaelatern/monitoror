package models

import (
	"time"

	coreModels "github.com/Vaelatern/monitoror/models"
)

type Pipeline struct {
	ID         int
	Branch     string
	Author     coreModels.Author
	Status     string
	StartedAt  *time.Time
	FinishedAt *time.Time
}
