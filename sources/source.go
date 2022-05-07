package sources

import (
	"github.com/flywave/morph/models"
)

type Source interface {
	Migrations() (migrations []*models.Migration)
}
