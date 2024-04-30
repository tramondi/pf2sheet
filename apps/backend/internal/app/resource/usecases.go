package resource

import (
	"log/slog"

	"github.com/alionapermes/pf2sheet/internal/app/usecase"
)

type Usecases struct {
	CreateSheet      *usecase.CreateSheetUsecase
	DeleteSheet      *usecase.DeleteSheetUsecase
	GetAllAncestries *usecase.GetAllAncestriesUsecase
	GetAllClasses    *usecase.GetAllClassesUsecase
	GetAllSheets     *usecase.GetAllSheetsUsecase
	Signin           *usecase.SigninUsecase
	Signup           *usecase.SignupUsecase
	Signout          *usecase.SignoutUsecase
	UpdateSheet      *usecase.UpdateSheetUsecase
	UpdatePlayer     *usecase.UpdatePlayerUsecase
	DeletePlayer     *usecase.DeletePlayerUsecase
	GetPlayer        *usecase.GetPlayerUsecase
}

func initUsecases(
	logger *slog.Logger,
	repos Repositories,
	services Services,
) Usecases {
	createSheetUsecase := usecase.NewCreateSheetsUsecase(
		logger, repos.Sheets, repos.Knowledge)
	deleteSheetUsecase := usecase.NewDeleteSheetsUsecase(
		logger, repos.Sheets)
	getAllAncestriesUsecase := usecase.NewGetAllAncestriesUsecase(
		logger, repos.Knowledge)
	getAllClassesUsecase := usecase.NewGetAllClassesUsecase(
		logger, repos.Knowledge)
	getAllSheetsUsecase := usecase.NewGetAllSheetsUsecase(
		logger, repos.Knowledge, repos.Sheets)
	signinUsecase := usecase.NewSigninUsecase(
		logger, services.Auth, repos.Sheets)
	signupUsecase := usecase.NewSignupUsecase(
		logger, services.Auth, services.Players)
	signoutUsecase := usecase.NewSignoutUsecase(logger, services.Auth)
	updatePlayerUsecase := usecase.NewUpdatePlayerUsecase(logger, repos.Players)
	deletePlayerUsecase := usecase.NewDeletePlayerUsecase(logger, repos.Players)
	updateSheetUsecase := usecase.NewUpdateSheetUsecase(logger, repos.Sheets)
	getPlayerUsecase := usecase.NewGetPlayerUsecase(logger, repos.Players)

	return Usecases{
		CreateSheet:      &createSheetUsecase,
		DeleteSheet:      &deleteSheetUsecase,
		GetAllAncestries: &getAllAncestriesUsecase,
		GetAllClasses:    &getAllClassesUsecase,
		GetAllSheets:     &getAllSheetsUsecase,
		Signin:           &signinUsecase,
		Signup:           &signupUsecase,
		Signout:          &signoutUsecase,
		UpdatePlayer:     &updatePlayerUsecase,
		DeletePlayer:     &deletePlayerUsecase,
		UpdateSheet:      &updateSheetUsecase,
		GetPlayer:        &getPlayerUsecase,
	}
}
