package pg

import (
	"context"
	"log/slog"

	"github.com/doug-martin/goqu/v9"

	"github.com/alionapermes/pf2sheet/internal/domain/entity"
	get_all_ancestries "github.com/alionapermes/pf2sheet/internal/infra/goqu-pg/get-all-ancestries"
	get_all_classes "github.com/alionapermes/pf2sheet/internal/infra/goqu-pg/get-all-classes"
)

type KnowledgeRepo struct {
	logger *slog.Logger
	db     *goqu.Database
}

func NewKnowledgeRepo(
	logger *slog.Logger,
	db *goqu.Database,
) *KnowledgeRepo {
	return &KnowledgeRepo{
		logger: logger,
		db:     db,
	}
}

func (self *KnowledgeRepo) GetAllAncestries(
	ctx context.Context,
) ([]entity.Ancestry, error) {
	output, err := get_all_ancestries.DB(self.db).Query(ctx)
	if err != nil {
		return nil, wrapQueryError(err, "failed to get ancestries")
	}

	ancestries := make([]entity.Ancestry, 0, len(output.Items))

	for _, item := range output.Items {
		ancestry := entity.Ancestry{
			ID:    entity.AncestryID(item.Ancestry.ID),
			Code:  item.Ancestry.Code,
			Title: item.Ancestry.Title,
		}

		ancestries = append(ancestries, ancestry)
	}

	return ancestries, nil
}

func (self *KnowledgeRepo) GetAllClasses(
	ctx context.Context,
) ([]entity.Class, error) {
	output, err := get_all_classes.DB(self.db).Query(ctx)
	if err != nil {
		return nil, wrapQueryError(err, "failed to get ancestries")
	}

	classes := make([]entity.Class, 0, len(output.Items))

	for _, item := range output.Items {
		class := entity.Class{
			ID:    entity.ClassID(item.Class.ID),
			Code:  item.Class.Code,
			Title: item.Class.Title,
		}

		classes = append(classes, class)
	}

	return classes, nil
}
