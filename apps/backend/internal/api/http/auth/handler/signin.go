package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/alionapermes/pf2sheet/internal/app/resource"
)

func Signin(container resource.Container) echo.HandlerFunc {
	type credentials struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	signin := container.Usecases.Signin

	return func(ctx echo.Context) error {
		var creds credentials
		if err := ctx.Bind(&creds); err != nil {
			return ctx.NoContent(http.StatusBadRequest)
		}

		session, err := signin.Execute(
			context.Background(),
			creds.Login,
			creds.Password,
		)
		if err != nil {
			return ctx.NoContent(http.StatusUnauthorized)
		}

		// sheetDTOs := make([]dto.Sheet, 0, len(sheets))
		// for _, sheet := range sheets {
		// 	var ancestryDTO *dto.Ancestry
		// 	var classDTO *dto.Class
		//
		// 	if ancestry := sheet.Ancestry; ancestry != nil {
		// 		ancestryDTO = &dto.Ancestry{
		// 			Code:  ancestry.Code,
		// 			Title: ancestry.Title,
		// 		}
		// 	}
		//
		// 	if class := sheet.Class; class != nil {
		// 		classDTO = &dto.Class{
		// 			Code:  class.Code,
		// 			Title: class.Title,
		// 		}
		// 	}
		//
		// 	sheetDTOs = append(sheetDTOs, dto.Sheet{
		// 		ID:         sheet.ID.Value(),
		// 		Ancestry:   ancestryDTO,
		// 		Class:      classDTO,
		// 		Background: sheet.Background,
		// 		FullName:   sheet.FullName,
		// 		Level:      sheet.Level,
		// 		HpCurrent:  sheet.HpCurrent,
		// 		HpMax:      sheet.HpMax,
		// 	})
		// }

		ctx.SetCookie(&http.Cookie{
			Name:  "session_token",
			Value: session.Token,
		})

		return ctx.NoContent(http.StatusOK)
	}
}
