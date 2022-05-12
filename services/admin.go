package services

func (h *Handler) SetRole(email string, role int16) error {
	user, err := h.GetUserByEmail(email)
	if err != nil {
		return err
	}
	err = h.db.SetRole(user, role)
	return err
}
