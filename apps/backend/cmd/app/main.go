package main

import (
	"github.com/alionapermes/pf2sheet/internal/app/server"
)

func main() {
	s := server.Server{}
	s.Start()

	// e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	// e.Logger.Fatal(e.Start(":1323"))
}
