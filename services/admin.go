package services

func (h *Handler) SetRole(email string, roleId int16) error {
	user, err := h.GetUserByEmail(email)
	if err != nil {
		return err
	}
	user.RoleId = roleId
	user.Role = *h.db.GetRoleById(roleId)
	err = h.db.SetRole(user)
	return err
}
