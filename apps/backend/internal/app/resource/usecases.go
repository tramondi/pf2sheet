package resource

import (
	"log/slog"

	"github.com/alionapermes/pf2sheet/internal/app/usecase"
)

type Usecases struct {
	GetAllAncestries usecase.GetAllAncestriesUsecase
	GetAllClasses    usecase.GetAllClassesUsecase
	Signin           usecase.SigninUsecase
	Signup           usecase.SignupUsecase
}

func initUsecases(
	logger *slog.Logger,
	repos Repositories,
	services Services,
) Usecases {
	getAllAncestriesUsecase := usecase.NewGetAllAncestriesUsecase(
		logger,
		repos.Knowledge,
	)
	getAllClassesUsecase := usecase.NewGetAllClassesUsecase(
		logger,
		repos.Knowledge,
	)
	signinUsecase := usecase.NewSigninUsecase(
		logger,
		services.Auth,
		repos.Sheets,
	)
	signupUsecase := usecase.NewSignupUsecase(
		logger,
		services.Auth,
		services.Players,
	)

	return Usecases{
		GetAllAncestries: getAllAncestriesUsecase,
		GetAllClasses:    getAllClassesUsecase,
		Signin:           signinUsecase,
		Signup:           signupUsecase,
	}
}
