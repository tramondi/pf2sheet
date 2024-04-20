package pg

import (
	"context"
	"strings"

	"github.com/alionapermes/pf2sheet/internal/domain/entity"
	"github.com/alionapermes/pf2sheet/internal/infra/sqlc-pg/dao"
)

type KnowledgeRepo struct {
	logger  interface{}
	querier dao.Querier
}

func NewKnowledgeRepo(logger interface{}, querier dao.Querier) *KnowledgeRepo {
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
		return nil, wrapQueryError(err, "ancestries not found")
	}

	ancestries := make([]entity.Ancestry, 0, len(rows))
	for _, row := range rows {
		ancestry := entity.Ancestry{
			ID:    entity.AncestryID(row.Ancestry.ID),
			Code:  strings.TrimSpace(row.Ancestry.Code),
			Title: strings.TrimSpace(row.Ancestry.Title),
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
		return nil, wrapQueryError(err, "classes not found")
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
