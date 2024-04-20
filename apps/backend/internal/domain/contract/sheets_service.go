package contract

import (
	"context"

	"github.com/alionapermes/pf2sheet/internal/domain/entity"
)

type SheetsService interface {
	AddSheet(
		ctx context.Context,
		playerID entity.PlayerID,
		sheet entity.Sheet,
	) (entity.Sheet, error)
}
