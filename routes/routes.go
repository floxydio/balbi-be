package routes

import (
	"balbibe/handler"
	"balbibe/middleware"

	"github.com/labstack/echo/v4"
)

func Setup(e *echo.Echo) {

	// Static For Marketplace
	e.Static("/img/market/", "./upload")
	e.Static("/img/user/", "./userimg")
	e.GET("/", handler.Index)

	// Auth
	e.POST("/register", handler.SignUp)
	e.POST("/login", handler.SignIn)
	e.POST("/verify", handler.VerifyEmail)
	e.PUT("/edit/user/:id", handler.EditProfile, middleware.IsAuth)

	// Buku

	// Course
	e.GET("/course", handler.GetCourseAll, middleware.IsAuth)
	e.GET("/course/:id", handler.GetCourseById, middleware.IsAuth)
	e.GET("/history-transaction/:id", handler.GetHistoryTransaction, middleware.IsAuth)
	e.PUT("/history-transaction/edit/:id/history/:idhistory", handler.EditHistoryTransaction)
	e.POST("/create/course", handler.CreateCourse)

	// Ads
	e.GET("/ads", handler.GetAdsAll, middleware.IsAuth)
	e.POST("/ads/:id", handler.CountInView)

	// Event
	e.GET("/event", handler.GetEventAll)

	// Konsul
	e.POST("/create/konsul", handler.CreateKonsul, middleware.IsAuth)

}
