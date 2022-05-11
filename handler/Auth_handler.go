package handler

import (
	"balbibe/models"
	"balbibe/repository"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// const CONFIG_SMTP_HOST = "smtp.gmail.com"
// const CONFIG_SMTP_PORT = 587
// const CONFIG_SENDER_NAME = "BalbiID - Success"
// const CONFIG_AUTH_EMAIL = "diooktar@gmail.com"
// const CONFIG_AUTH_PASSWORD = "kqzqetxuotcrfysv"

func SignUp(c echo.Context) error {
	passwordInput := []byte(c.FormValue("password"))

	hashedPassword, _ := bcrypt.GenerateFromPassword(passwordInput, 8)
	userModel := models.User{
		Nama:      c.FormValue("nama"),
		Asal:      c.FormValue("asal"),
		Email:     c.FormValue("email"),
		Username:  c.FormValue("username"),
		Password:  string(hashedPassword),
		Photo:     c.FormValue("photo"),
		Jawaban:   c.FormValue("jawaban"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		CreatedAt: time.Now().Format("02-Jan-2006"),
	}
	if errRespon := c.Validate(userModel); errRespon != nil {
		return c.JSON(http.StatusBadRequest, errRespon)
	}
	// m := gomail.NewMessage()
	// m.SetHeader("From", "halo@balbiid.com")
	// m.SetHeader("To", userModel.Email)
	// m.SetHeader("Subject", "Hello!")
	// token := jwt.New(jwt.SigningMethodHS256)
	// t, _ := token.SignedString([]byte("tokenregist"))

	// m.SetBody("text/html", "<h1 style='text-align:center'>Terima kasih telah melakukan registrasi</h1><br><br>"+"<p>Token anda adalah :"+t+"</p>")

	// // Send the email to Bob
	// d := gomail.NewPlainDialer(CONFIG_SMTP_HOST, 587, CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD)
	// if err := d.DialAndSend(m); err != nil {
	// 	panic(err)
	// }
	return repository.CreateUser(userModel, c)

}

func SignIn(c echo.Context) error {
	user := models.User{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}
	return repository.LoginUser(user.Email, user, c)
}

func EditProfile(c echo.Context) error {
	idUser, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user := models.User{
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	if errRespon := c.Bind(user); errRespon != nil {
		return c.JSON(http.StatusBadRequest, errRespon)
	}
	return repository.EditUserById(idUser, user, c)
}

func CheckUser(c echo.Context) error {
	jwtOk := c.FormValue("jwt")

	token, _ := jwt.Parse(jwtOk, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	return c.JSON(http.StatusOK, echo.Map{
		"data":  token.Claims,
		"token": token.Valid,
	})
}

func VerifyEmail(c echo.Context) error {
	user := c.FormValue("email")
	token := c.FormValue("token")

	return repository.VerifyEmail(user, token, c)

}

func ForgotPassword(c echo.Context) error {

	user := models.User{}
	if errRespon := c.Bind(&user); errRespon != nil {
		return c.JSON(http.StatusBadRequest, errRespon)
	}
	return repository.ForgotPassword(user, c)
}

// func sendMail(to []string, cc []string, subject, message string) error {
// 	body := "From: " + CONFIG_SENDER_NAME + "\n" +
// 		"To: " + strings.Join(to, ",") + "\n" +
// 		"Cc: " + strings.Join(cc, ",") + "\n" +
// 		"Subject: " + subject + "\n\n" +
// 		message

// 	auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)
// 	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

// 	err := smtp.SendMail(smtpAddr, auth, CONFIG_AUTH_EMAIL, append(to, cc...), []byte(body))
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
