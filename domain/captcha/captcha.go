package captcha

import (
	"math/rand"
	"strings"
	"time"

	"github.com/pastukhov-aleksandr/bookstore_utils-go/rest_errors"
)

const (
	expirationTime             = 15
	refreshExpirationTime      = 60 * 24 * 7
	grantTypePassword          = "password"
	grandTypeClientCredentials = "client_credentials"
)

type CaptchaRequest struct {
	ID       string `json:"id"`
	ClientID int64  `json:"client_id"`
}

type ValidateRequest struct {
	Email string `json:"email"`
	Pin   string `json:"Pin"`
}

func (cr *CaptchaRequest) Validate() rest_errors.RestErr {
	cr.ID = strings.TrimSpace(cr.ID)

	if cr.ID == "" {
		return rest_errors.NewBadRequestError("invalid mail")
	}

	return nil
}

type Captcha struct {
	ID       string `json:"id"`
	ClientID int64  `json:"client_id,omitempty"`
	Pin      string `json:"pin"`
}

func (cp *Captcha) Validate() rest_errors.RestErr {

	return nil
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// Generate a random string of A-Z chars with len = l
func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(48, 57))
	}
	return string(bytes)
}

func GetNewCaptcha(Id string, clientID int64) Captcha {
	rand.Seed(time.Now().UnixNano())
	return Captcha{
		ID:       Id,
		ClientID: clientID,
		Pin:      randomString(4),
	}
}

// func (at Captcha) IsExpired() bool {
// 	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
// }

func (at *Captcha) Generate(access_sicret string, refresh_sicret string) rest_errors.RestErr {

	return nil
}
