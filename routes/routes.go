package routes

import (
	"balbibe/handler"
	"balbibe/middleware"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Setup(e *echo.Echo) {
	// t := &Template{
	// 	templates: template.Must(template.ParseGlob("postarticle/*.html")),
	// }
	// e.Renderer = t
	// Static For Marketplace
	e.Static("/img/market/", "./uploads/market")
	e.Static("/img/user/", "./uploads/userimg")
	e.Static("/img/ads/", "./uploads/ads")
	e.GET("/", handler.Index)
	e.GET("/create/article", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", nil)
	})
	// Auth

	e.POST("/register", handler.SignUp)
	e.POST("/login", handler.SignIn)
	// e.POST("/verify", handler.VerifyEmail)
	// e.POST("/forgot", handler.ForgotPassword)
	e.PUT("/edit/user/:id", handler.EditProfile, middleware.IsAuth)
	e.POST("/checkuser", handler.CheckUser)

	// Buku
	e.POST("/create/notification", handler.CreateNotification, middleware.IsAuth)

	// Course
	e.GET("/course", handler.GetCourseAll, middleware.IsAuth)
	e.GET("/course/:id", handler.GetCourseById, middleware.IsAuth)
	e.POST("/review/course", handler.AddReviewById, middleware.IsAuth)
	e.GET("/history-transaction/:id", handler.GetHistoryTransaction, middleware.IsAuth)
	e.PUT("/history-transaction/edit/:id/history/:idhistory", handler.EditHistoryTransaction, middleware.IsAuth)
	e.POST("/create/course", handler.CreateCourse, middleware.IsAuth)

	// Ads
	e.GET("/get-ads", handler.GetAdsAll, middleware.IsAuth)
	e.POST("/ads", handler.CreateAds)
	e.POST("/ads/:id", handler.CountInView, middleware.IsAuth)

	// ANAK
	e.POST("/create/anak-kesatu", handler.CreateAnakPertama, middleware.IsAuth)
	e.POST("/create/anak-kedua", handler.CreateAnakKedua, middleware.IsAuth)
	e.POST("/create/anak-ketiga", handler.CreateAnakKetiga, middleware.IsAuth)
	e.GET("/anak/:id", handler.GetAnakById, middleware.IsAuth)

	// Event
	e.GET("/event", handler.GetEventAll, middleware.IsAuth)

	// VideoPlace
	e.GET("/video-place", handler.GetVideoPlaceAll, middleware.IsAuth)

	// Article
	e.POST("/article", handler.CreateArticle)
	e.GET("/articles", handler.GetArticle)

	// Konsul
	e.POST("/create/konsul", handler.CreateKonsul, middleware.IsAuth)

	// Community
	e.GET("/community", handler.GetAllKomunitas)
	e.GET("/community/post/:id", handler.GetAllPostById)

}
