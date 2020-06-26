package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pastukhov-aleksandr/bookstore_utils-go/rest_errors"
	atDomain "github.com/pastukhov-aleksandr/captcha/domain/captcha"
	"github.com/pastukhov-aleksandr/captcha/services/captcha"
)

type CaptchaHandler interface {
	Validate(*gin.Context)
	Create(*gin.Context)
}

type captchaHandler struct {
	service captcha.Service
}

func NewCaptchaHandler(service captcha.Service) CaptchaHandler {
	return &captchaHandler{
		service: service,
	}
}

func (handler *captchaHandler) Validate(c *gin.Context) {
	var request atDomain.ValidateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	if err := handler.service.Validate(request); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (handler *captchaHandler) Create(c *gin.Context) {
	var request atDomain.CaptchaRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	err := handler.service.Create(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, "OK")
}
