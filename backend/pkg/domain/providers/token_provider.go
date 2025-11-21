package providers

import "github.com/golang-jwt/jwt/v5"

type TokenProvider interface {
	Generate(claims interface{}) (string, error)
	Validate(token string) (interface{}, error)
}

type Claims struct {
    UserID string `json:"uid"`
    OrganizationID string `json:"org_id"`
    Plan string `json:"plan"`
    jwt.RegisteredClaims
}
