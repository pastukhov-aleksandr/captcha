package db

import (
	"errors"
	"time"

	"github.com/gocql/gocql"
	"github.com/pastukhov-aleksandr/bookstore_utils-go/rest_errors"
	"github.com/pastukhov-aleksandr/captcha/clients/cassandra"
	"github.com/pastukhov-aleksandr/captcha/domain/captcha"
)

const (
	queryGetCaptcha        = "SELECT pin FROM pincode WHERE id=?;"
	queryCreateAccessToken = "INSERT INTO pincode(id, pin, client_id, time) VALUES (?, ?, ?, ?) USING TTL 180;"
	queryUpdateExpires     = "UPDATE refresh_tokens SET expires=? WHERE refresh_tokens=?;"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*captcha.Captcha, rest_errors.RestErr)
	Create(captcha.Captcha) rest_errors.RestErr
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*captcha.Captcha, rest_errors.RestErr) {
	var result captcha.Captcha
	if err := cassandra.GetSession().Query(queryGetCaptcha, id).Scan(
		&result.Pin,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, rest_errors.NewNotFoundError("no pin found with given email")
		}
		return nil, rest_errors.NewInternalServerError("error when trying to get current email", errors.New("database error"))
	}
	return &result, nil
}

func (r *dbRepository) Create(cp captcha.Captcha) rest_errors.RestErr {
	if err := cassandra.GetSession().Query(queryCreateAccessToken,
		cp.ID,
		cp.Pin,
		cp.ClientID,
		time.Now(),
	).Exec(); err != nil {
		return rest_errors.NewInternalServerError("error when trying to save refresh token in database", err)
	}
	return nil
}
