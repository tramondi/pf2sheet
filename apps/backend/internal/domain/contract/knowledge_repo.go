package contract

import (
	"context"

	"github.com/alionapermes/pf2sheet/internal/domain/entity"
)

type KnowledgeRepo interface {
	GetAllAncestries(ctx context.Context) ([]entity.Ancestry, error)

	GetAllClasses(ctx context.Context) ([]entity.Class, error)
}
