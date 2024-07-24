package params

import (
	"github.com/Vaelatern/monitoror/internal/pkg/validator"
)

type (
	Validator interface {
		Validate() []validator.Error
	}

	Default struct{}
)

// Validate is empty by default. Override it if you want custom Validate
func (p *Default) Validate() []validator.Error {
	return nil
}
