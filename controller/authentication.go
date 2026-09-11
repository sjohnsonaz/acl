package controller

import (
	"context"
	"net/http"

	"acl/access"
	"acl/key"

	"github.com/cardboardrobots/baseerror"
	"github.com/cardboardrobots/listener"
	"github.com/cardboardrobots/token"
	"github.com/golang-jwt/jwt/v4"
)

func GetClaimsGRPC(ctx context.Context, key string) (AuthClaims, error) {
	var claims AuthClaims
	err := token.ParseClaims(listener.GetMetadataGRPC(ctx).Authorization(), key, &claims)
	return claims, err
}

func GetClaimsHTTP(r *http.Request, key string) (AuthClaims, error) {
	var claims AuthClaims
	err := token.ParseClaims(listener.GetMetadataHTTP(r).Authorization(), key, &claims)
	return claims, err
}

func CreateToken(claims AuthClaims, key string) (string, error) {
	return token.CreateToken(key, claims)
}

type AuthClaims struct {
	jwt.RegisteredClaims
	Scopes string                                         `json:"scopes"`
	Access map[key.OrganizationID]access.OrganizationRole `json:"access"`
}

func (ac AuthClaims) Valid() error {
	if err := ac.RegisteredClaims.Valid(); err != nil {
		return err
	}

	if ac.Subject == "" {
		return baseerror.ErrUnauthenticated
	}

	return nil
}
