package usecase

import (
	"context"
	"log/slog"

	"github.com/alionapermes/pf2sheet/internal/domain/contract"
	"github.com/alionapermes/pf2sheet/internal/domain/entity"
)

type GetAllClassesUsecase struct {
	logger        *slog.Logger
	knowledgeRepo contract.KnowledgeRepo
}

func NewGetAllClassesUsecase(
	logger *slog.Logger,
	knowledgeRepo contract.KnowledgeRepo,
) GetAllClassesUsecase {
	return GetAllClassesUsecase{
		logger:        logger,
		knowledgeRepo: knowledgeRepo,
	}
}

func (self *GetAllClassesUsecase) Execute(
	ctx context.Context,
) ([]entity.Class, error) {
	return self.knowledgeRepo.GetAllClasses(ctx)
}
