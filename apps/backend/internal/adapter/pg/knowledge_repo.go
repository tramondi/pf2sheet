package pg

import (
	"context"
	"log/slog"

	"github.com/alionapermes/pf2sheet/internal/domain/entity"
	"github.com/alionapermes/pf2sheet/internal/infra/sqlc-pg/dao"
)

type KnowledgeRepo struct {
	logger  *slog.Logger
	querier dao.Querier
}

func NewKnowledgeRepo(
	logger *slog.Logger,
	querier dao.Querier,
) *KnowledgeRepo {
	return &KnowledgeRepo{
		logger:  logger,
		querier: querier,
	}
}

func (self *KnowledgeRepo) GetAllAncestries(
	ctx context.Context,
) ([]entity.Ancestry, error) {
	rows, err := self.querier.GetAllAncestries(ctx)
	if err != nil {
		return nil, wrapQueryError(err, "failed to get ancestries")
	}

	ancestries := make([]entity.Ancestry, 0, len(rows))
	for _, row := range rows {
		ancestry := entity.Ancestry{
			ID:    entity.AncestryID(row.Ancestry.ID),
			Code:  row.Ancestry.Code,
			Title: row.Ancestry.Title,
		}

		ancestries = append(ancestries, ancestry)
	}

	return ancestries, nil
}

func (self *KnowledgeRepo) GetAllClasses(
	ctx context.Context,
) ([]entity.Class, error) {
	rows, err := self.querier.GetAllClasses(ctx)
	if err != nil {
		return nil, wrapQueryError(err, "failed to get classes")
	}

	classes := make([]entity.Class, 0, len(rows))
	for _, row := range rows {
		class := entity.Class{
			ID:    entity.ClassID(row.Class.ID),
			Code:  row.Class.Code,
			Title: row.Class.Title,
		}

		classes = append(classes, class)
	}

	return classes, nil
}
