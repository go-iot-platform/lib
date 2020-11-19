package utl

import (
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-iot-platform/lib/errors"
	"github.com/go-iot-platform/lib/rbac/model"
)

// ServiceName Seperate Char
const ServiceName = "rbac"

// Struct to parse claim
type jwtUser struct {
	ID             int    `json:"id,omitempty"`
	CustomerID     int    `json:"c,omitempty"`
	CustomerNumber string `json:"c_n,omitempty"`
	Username       string `json:"u,omitempty"`
	UserUUID       string `json:"u_uid,omitempty"`
	Role           int8   `json:"r,omitempty"`
	jwt.StandardClaims
}

// ParseToken parses token from Authorization header
func ParseToken(key string, algo string, token string) (*model.AuthUser, error) {

	parts := strings.SplitN(token, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return nil, errors.BadRequest(ServiceName, "rbac:utl:ParseToken:invalid:%s", token)
	}
	j, err := jwt.ParseWithClaims(parts[1], &jwtUser{}, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod(algo) != token.Method {
			return nil, errors.Forbidden(ServiceName, "rbac:utl:ParseToken:forbidden:%s", parts[1])
		}
		return []byte(key), nil
	})
	if err != nil || !j.Valid {
		return nil, err
	}
	if claims, ok := j.Claims.(*jwtUser); ok {
		return &model.AuthUser{
			ID:             claims.ID,
			Username:       claims.Username,
			CustomerNumber: claims.CustomerNumber,
			CustomerID:     claims.CustomerID,
			UserUuid:       claims.UserUUID,
			Role:           model.AccessRole(claims.Role),
		}, nil
	}
	//check cache
	return nil, nil
}
