package models

import "github.com/Vaelatern/monitoror/models"

type (
	Commit struct {
		SHA    string
		Author models.Author
	}
)
