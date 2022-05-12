package services

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type Generator struct {
	secretKey string
}

const UserID string = "user_id"

var (
	ErrInvalidToken = errors.New("token not valid")
	ErrNotMapClaims = errors.New("parsedToken.Claims not jwt.MapClaims")
	ErrIdNotFound   = errors.New("id not found")
)

func New(secretKey string) *Generator {
	return &Generator{secretKey: secretKey}
}

func (g *Generator) Generate(claims map[string]interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(claims))

	return token.SignedString([]byte(g.secretKey))
}

func (g *Generator) ParseToken(token string) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(g.secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		return nil, ErrInvalidToken
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrNotMapClaims
	}

	return claims, nil
}

func (h *Handler) GetUserID(jwt string) (uuid.UUID, error) {
	claims, err := h.generator.ParseToken(jwt)
	if err != nil {
		return uuid.UUID{}, err
	}

	id, ok := claims[UserID]
	if !ok {
		return uuid.UUID{}, ErrIdNotFound
	}

	userID, err := uuid.Parse(id.(string))
	if err != nil {
		return uuid.UUID{}, err
	}

	return userID, nil
}

func (h *Handler) GenerateUserIDJwt(userId uuid.UUID) (string, error) {
	token, err := h.generator.Generate(map[string]interface{}{UserID: userId})
	if err != nil {
		return token, err
	}
	return token, err
}
