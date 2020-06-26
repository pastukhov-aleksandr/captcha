package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pastukhov-aleksandr/captcha/controllers/ping"
	"github.com/pastukhov-aleksandr/captcha/http"
	"github.com/pastukhov-aleksandr/captcha/repositore/db"
	"github.com/pastukhov-aleksandr/captcha/services/captcha"
)

var router = gin.Default()

func StartApplication() {
	atHandler := http.NewCaptchaHandler(
		captcha.NewService(db.NewRepository()))

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	router.Use(cors.New(config))

	router.GET("/captcha/ping", ping.Ping)
	router.POST("/captcha/mail", atHandler.Create)
	router.POST("/captcha/validate", atHandler.Validate)

	router.Run(":8082")
}
