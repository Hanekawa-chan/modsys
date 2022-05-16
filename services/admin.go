package services

import (
	"github.com/rs/zerolog/log"
	"strconv"
)

func (h *Handler) SetRole(email string, role int16) error {
	log.Info().Msg(email + " " + strconv.Itoa(int(role)))
	user, err := h.GetUserByEmail(email)
	if err != nil {
		return err
	}
	err = h.db.SetRole(user, role)
	return err
}
