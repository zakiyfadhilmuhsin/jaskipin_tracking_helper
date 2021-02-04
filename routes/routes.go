package routes

import (
	"jaskipin-front/controllers"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// Root
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to the Jungle!")
	})

	// Static File
	//e.Static("/", "./")

	// Destination Group
	// e.GET("/destinations_migrate", controllers.MigrateDestination)
	// e.POST("/destinations/:screenshot", controllers.CreateDestination)
	// e.GET("/destinations", controllers.FetchAllDestination)
	// e.GET("/destinations/:id", controllers.GetDestination)
	// e.PUT("/destinations/:id", controllers.UpdateDestination)
	// e.DELETE("/destinations/:id", controllers.DeleteDestination)

	// Tracking Group
	e.GET("/tracking/:id_order", controllers.GetTracking)
	e.GET("/manifest/:carrier/:trackingNumber", controllers.GetManifest)
	e.POST("/manifest", controllers.CreateManifest)

	return e
}
