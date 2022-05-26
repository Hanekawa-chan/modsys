package services

import (
	"awesomeProject/models"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (h *Handler) GetUserByCredentials(credentials models.LoginCredentials) (*models.User, error) {
	var user *models.User
	var err error
	user, err = h.db.GetUserByEmail(credentials.Email)
	if user == nil || !CheckPasswordHash(credentials.Password, user.Password) {
		return nil, errors.New("invalid password")
	}
	if err != nil {
		return nil, err
	}
	return user, err
}

func (h *Handler) GetUserByEmail(email string) (*models.User, error) {
	var user *models.User
	var err error
	user, err = h.db.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (h *Handler) SaveUser(user *models.User) error {
	lastUser, _ := h.GetUserByEmail(user.Email)
	if lastUser != nil {
		return errors.New("user already exists")
	}
	hash, err := HashPassword(user.Password)
	if err != nil {
		return err
	}
	user = models.NewUser(user.Name, user.Surname, user.Email, hash)
	err = h.db.SaveUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (h *Handler) GetUserByID(id uuid.UUID) (*models.User, error) {
	var user *models.User
	var err error
	user, err = h.db.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (h *Handler) GetAuthenticatedUserID(r *http.Request) (uuid.UUID, error) {
	var user *models.User
	var err error
	token, err := r.Cookie("jwt")
	if err != nil {
		return uuid.UUID{}, err
	}

	id, err := h.GetUserID(token.Value)
	if err != nil {
		return uuid.UUID{}, err
	}
	user, err = h.db.GetUserByID(id)
	if err != nil {
		return uuid.UUID{}, err
	}
	return user.GetID(), err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
