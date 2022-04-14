package routes

import (
	"balbibe/handler"
	"balbibe/middleware"

	"github.com/labstack/echo/v4"
)

func Setup(e *echo.Echo) {

	// Static For Marketplace
	e.Static("/img/market/", "./uploads/market")
	e.Static("/img/user/", "./userimg")
	e.Static("/img/ads/", "./uploads/ads")
	e.GET("/", handler.Index)

	// Auth

	e.POST("/register", handler.SignUp)
	e.POST("/login", handler.SignIn)
	// e.POST("/verify", handler.VerifyEmail)
	e.POST("/forgot", handler.ForgotPassword)
	e.PUT("/edit/user/:id", handler.EditProfile, middleware.IsAuth)

	// Buku

	// Notification
	e.POST("/create/notification", handler.CreateNotification, middleware.IsAuth)

	// Course
	e.GET("/course", handler.GetCourseAll, middleware.IsAuth)
	e.GET("/course/:id", handler.GetCourseById, middleware.IsAuth)
	e.GET("/history-transaction/:id", handler.GetHistoryTransaction, middleware.IsAuth)
	e.PUT("/history-transaction/edit/:id/history/:idhistory", handler.EditHistoryTransaction)
	e.POST("/create/course", handler.CreateCourse)

	// Ads
	e.GET("/get-ads", handler.GetAdsAll, middleware.IsAuth)
	e.POST("/ads", handler.CreateAds)
	e.POST("/ads/:id", handler.CountInView)

	// ANAK
	e.POST("/create/anak-kesatu", handler.CreateAnakPertama)
	e.POST("/create/anak-kedua", handler.CreateAnakKedua)
	e.POST("/create/anak-ketiga", handler.CreateAnakKetiga)
	e.GET("/anak/:id", handler.GetAnakById)

	// Event
	e.GET("/event", handler.GetEventAll)

	// VideoPlace
	e.GET("/video-place", handler.GetVideoPlaceAll)

	// Konsul
	e.POST("/create/konsul", handler.CreateKonsul, middleware.IsAuth)

}
