package pg

import (
	"context"
	"errors"
	"log/slog"

	"github.com/alionapermes/pf2sheet/internal/domain/entity"
	"github.com/alionapermes/pf2sheet/internal/infra/sqlc-pg/dao"
)

type SheetsRepo struct {
	logger  *slog.Logger
	querier dao.Querier
}

func NewSheetsRepo(logger *slog.Logger, querier dao.Querier) *SheetsRepo {
	return &SheetsRepo{
		logger:  logger,
		querier: querier,
	}
}

func (self *SheetsRepo) GetByPlayerID(
	ctx context.Context,
	playerID entity.PlayerID,
) ([]entity.Sheet, error) {
	rows, err := self.querier.GetPlayerSheets(ctx, int32(playerID.Value()))
	if err != nil {
		return nil,
			wrapQueryError(err, "sheets with player_id %d not found", playerID.Value())
	}

	sheets := make([]entity.Sheet, 0, len(rows))
	for _, row := range rows {
		var ancestry *entity.Ancestry
		var class *entity.Class

		if row.Ancestry != (dao.Ancestry{}) {
			ancestry = &entity.Ancestry{
				ID:    entity.AncestryID(row.Ancestry.ID),
				Code:  row.Ancestry.Code,
				Title: row.Ancestry.Title,
			}
		}

		if row.Class != (dao.Class{}) {
			class = &entity.Class{
				ID:    entity.ClassID(row.Class.ID),
				Code:  row.Class.Code,
				Title: row.Class.Title,
			}
		}

		sheet := entity.Sheet{
			ID:         entity.SheetID(row.Sheet.ID),
			PlayerID:   entity.PlayerID(row.Player.ID),
			Ancestry:   ancestry,
			Class:      class,
			FullName:   row.Sheet.Fullname,
			Background: row.Sheet.Background,
			Level:      row.Sheet.Level,
			HpCurrent:  row.Sheet.HpCurrent,
			HpMax:      row.Sheet.HpMax,
		}

		sheets = append(sheets, sheet)
	}

	return sheets, nil
}

func (self *SheetsRepo) Add(
	ctx context.Context,
	playerID entity.PlayerID,
	sheet entity.Sheet,
) (entity.SheetID, error) {
	return 0, errors.New("not implemented yet")
}

func (self *SheetsRepo) DeleteByID(
	ctx context.Context,
	id entity.SheetID,
) error {
	return errors.New("not implemented yet")
}
