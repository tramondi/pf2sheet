package usecase

import (
	"context"

	"github.com/alionapermes/pf2sheet/internal/domain/contract"
	"github.com/alionapermes/pf2sheet/internal/domain/entity"
)

type GetAllAncestriesUsecase struct {
	logger        interface{}
	knowledgeRepo contract.KnowledgeRepo
}

func NewGetAllAncestriesUsecase(
	logger interface{},
	knowledgeRepo contract.KnowledgeRepo,
) GetAllAncestriesUsecase {
	return GetAllAncestriesUsecase{
		logger:        logger,
		knowledgeRepo: knowledgeRepo,
	}
}

func (self *GetAllAncestriesUsecase) Execute(
	ctx context.Context,
) ([]entity.Ancestry, error) {
	return self.knowledgeRepo.GetAllAncestries(ctx)
}
