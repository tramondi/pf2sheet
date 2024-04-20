package contract

import (
	"context"

	"github.com/alionapermes/pf2sheet/internal/domain/entity"
)

type SheetsRepo interface {
	GetPlayerSheets(
		ctx context.Context,
		playerID entity.PlayerID,
	) ([]entity.Sheet, error)

	AddSheet(
		ctx context.Context,
		playerID entity.PlayerID,
		sheet entity.Sheet,
	) (entity.SheetID, error)

	DeleteSheet(ctx context.Context, id entity.SheetID) error
}
