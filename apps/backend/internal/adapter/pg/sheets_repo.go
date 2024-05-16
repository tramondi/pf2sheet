package pg

import (
	"context"
	"log/slog"

	"github.com/AlekSi/pointer"
	"github.com/doug-martin/goqu/v9"

	"github.com/alionapermes/pf2sheet/internal/domain/entity"
	add_sheet "github.com/alionapermes/pf2sheet/internal/infra/goqu-pg/add-sheet"
	del_sheet_by_id "github.com/alionapermes/pf2sheet/internal/infra/goqu-pg/del-sheet-by-id"
	get_sheet_by_id "github.com/alionapermes/pf2sheet/internal/infra/goqu-pg/get-sheet-by-id"
	get_sheets_by_player_id "github.com/alionapermes/pf2sheet/internal/infra/goqu-pg/get-sheets-by-player-id"
	upd_sheet "github.com/alionapermes/pf2sheet/internal/infra/goqu-pg/upd-sheet"
)

type SheetsRepo struct {
	logger *slog.Logger
	db     *goqu.Database
}

func NewSheetsRepo(logger *slog.Logger, db *goqu.Database) *SheetsRepo {
	return &SheetsRepo{
		logger: logger,
		db:     db,
	}
}

func (self *SheetsRepo) GetByPlayerID(
	ctx context.Context,
	playerID entity.PlayerID,
) ([]entity.Sheet, error) {
	input := get_sheets_by_player_id.Input{
		PlayerID: int64(playerID.Value()),
	}

	output, err := get_sheets_by_player_id.DB(self.db).Query(ctx, input)
	if err != nil {
		return nil, err
	}

	sheets := make([]entity.Sheet, 0, len(output.Items))

	for _, row := range output.Items {
		var ancestry *entity.Ancestry
		var class *entity.Class

		if row.Sheet.AncestryID != nil {
			ancestry = &entity.Ancestry{
				ID:    entity.AncestryID(*row.Ancestry.ID),
				Code:  *row.Ancestry.Code,
				Title: *row.Ancestry.Title,
			}
		}

		if row.Sheet.ClassID != nil {
			class = &entity.Class{
				ID:    entity.ClassID(*row.Class.ID),
				Code:  *row.Class.Code,
				Title: *row.Class.Title,
			}
		}

		sheet := entity.Sheet{
			ID:         entity.SheetID(row.Sheet.ID),
			PlayerID:   entity.PlayerID(row.Sheet.PlayerID),
			Ancestry:   ancestry,
			Class:      class,
			FullName:   row.Sheet.FullName,
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
	var ancestryID *int64 = nil
	var classID *int64 = nil

	if sheet.Ancestry != nil {
		ancestryID = pointer.ToInt64(int64(sheet.Ancestry.ID))
	}

	if sheet.Class != nil {
		classID = pointer.ToInt64(int64(sheet.Class.ID))
	}

	input := add_sheet.Input{
		Sheet: add_sheet.Sheet{
			PlayerID:   int64(playerID.Value()),
			AncestryID: ancestryID,
			ClassID:    classID,
			FullName:   sheet.FullName,
			Background: sheet.Background,
			Level:      sheet.Level,
			HpCurrent:  sheet.HpCurrent,
			HpMax:      sheet.HpMax,
		},
	}

	output, err := add_sheet.DB(self.db).Query(ctx, input)
	if err != nil {
		return 0, wrapQueryError(err, "failed to add sheet")
	}

	return entity.SheetID(output.SheetID), nil
}

func (self *SheetsRepo) DeleteByID(
	ctx context.Context,
	id entity.SheetID,
) error {
	input := del_sheet_by_id.Input{SheetID: int64(id.Value())}

	if err := del_sheet_by_id.DB(self.db).Query(ctx, input); err != nil {
		return err
	}

	return nil
}

func (self *SheetsRepo) Update(
	ctx context.Context,
	sheet entity.Sheet,
) error {
	var ancestryID, classID *int64

	if sheet.Ancestry != nil {
		ancestryID = pointer.ToInt64(int64(sheet.Ancestry.ID))
	}

	if sheet.Class != nil {
		classID = pointer.ToInt64(int64(sheet.Class.ID))
	}

	input := upd_sheet.Input{
		SheetID: int64(sheet.ID),
		Sheet: upd_sheet.Sheet{
			AncestryID: ancestryID,
			ClassID:    classID,
			FullName:   sheet.FullName,
			Background: sheet.Background,
			Level:      sheet.Level,
			HpCurrent:  sheet.HpCurrent,
			HpMax:      sheet.HpMax,
		},
	}

	if err := upd_sheet.DB(self.db).Query(ctx, input); err != nil {
		return wrapQueryError(err, "failed to delete sheet with id %d", sheet.ID)
	}

	return nil
}

func (self *SheetsRepo) GetByID(
	ctx context.Context,
	sheetID entity.SheetID,
) (entity.Sheet, error) {
	input := get_sheet_by_id.Input{SheetID: int64(sheetID)}

	output, err := get_sheet_by_id.DB(self.db).Query(ctx, input)
	if err != nil {
		return entity.Sheet{},
			wrapQueryError(err, "failed to get sheet by id %d", sheetID)
	}

	var ancestry *entity.Ancestry
	var class *entity.Class

	if output.Sheet.AncestryID != nil {
		ancestry = &entity.Ancestry{
			ID:    entity.AncestryID(*output.Ancestry.ID),
			Code:  *output.Ancestry.Code,
			Title: *output.Ancestry.Title,
		}
	}

	if output.Sheet.ClassID != nil {
		class = &entity.Class{
			ID:    entity.ClassID(*output.Class.ID),
			Code:  *output.Class.Code,
			Title: *output.Class.Title,
		}
	}

	sheet := entity.Sheet{
		ID:         entity.SheetID(output.Sheet.ID),
		PlayerID:   entity.PlayerID(output.Sheet.PlayerID),
		Level:      output.Sheet.Level,
		FullName:   output.Sheet.FullName,
		Ancestry:   ancestry,
		Class:      class,
		Background: output.Sheet.Background,
		HpCurrent:  output.Sheet.HpCurrent,
		HpMax:      output.Sheet.HpMax,
	}

	return sheet, nil
}
