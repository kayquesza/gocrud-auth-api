package service

import (
	"github.com/kayquesza/gocrud-auth-api/src/configuration/rest_err"
	"github.com/kayquesza/gocrud-auth-api/src/model"
)

func (*userDomainInterface) FindUser(string) (
	*model.UserDomainInterface,
	*rest_err.RestErr) {
	return nil, nil
}
