package logic

import (
	"go.mod/dao"
	"go.mod/models"
)

func AccountCheck(user *models.Admin) bool {
	var admin []models.Admin
	dao.DB.Find(&admin)

	for _, uid := range admin {
		if uid.ID == user.ID {
			if uid.Password == user.Password {
				return true
			}
		}
	}

	return false
}
