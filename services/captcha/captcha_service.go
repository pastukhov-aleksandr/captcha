package captcha

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
	"github.com/pastukhov-aleksandr/bookstore_utils-go/rest_errors"
	"github.com/pastukhov-aleksandr/captcha/domain/captcha"
	"github.com/pastukhov-aleksandr/captcha/repositore/db"
)

const (
	ACCESS_SECRET  = "ACCESS_SECRET"
	REFRESH_SECRET = "REFRESH_SECRET"
)

type Service interface {
	Validate(captcha.ValidateRequest) rest_errors.RestErr
	Create(captcha.CaptchaRequest) rest_errors.RestErr
}

type service struct {
	dbRepo db.DbRepository
}

func NewService(dbRepo db.DbRepository) Service {
	return &service{
		dbRepo: dbRepo,
	}
}

func (s *service) Validate(request captcha.ValidateRequest) rest_errors.RestErr {
	cp, err := s.dbRepo.GetById(request.Email)
	if err != nil {
		return err
	}

	if cp.Pin != request.Pin {
		return rest_errors.NewBadRequestError("invalid pin code")
	}

	return nil
}

func (s *service) Create(request captcha.CaptchaRequest) rest_errors.RestErr {
	if err := request.Validate(); err != nil {
		return err
	}

	pin := captcha.GetNewCaptcha(request.ID, request.ClientID)

	if err := s.dbRepo.Create(pin); err != nil {
		return err
	}

	// send to email
	e := email.NewEmail()
	e.From = "Validate robot <pactx@yandex.ru>"
	e.To = []string{request.ID}
	e.Subject = "Pin validate"
	e.HTML = []byte(fmt.Sprintf("pin: <h1>%s</h1>", pin.Pin))
	if err := e.Send("smtp.yandex.ru:587", smtp.PlainAuth("", "pactx@yandex.ru", "mlvgfmvoesibkfhc", "smtp.yandex.ru")); err != nil {
		return rest_errors.NewInternalServerError("error smtp server", err)
	}

	return nil
}
