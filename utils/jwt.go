package utils

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model/system/request"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{
		SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey),
	}
}

func (j *JWT) ClearClaims(baseClaims request.BaseClaims) request.CustomClaims {

}
