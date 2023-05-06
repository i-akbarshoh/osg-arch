package entity

import "github.com/i-akbarshoh/osg-arch/internal/pkg/jwt"

type RegisterResponse struct {
	ID     string
	Tokens jwt.Tokens
}
