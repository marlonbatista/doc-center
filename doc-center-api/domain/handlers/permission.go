package handlers

import (
	"doc-center-api/domain/models"
	"doc-center-api/infra/database"
)

func GetAllPermissions(permission *[]models.Permission) (err error) {
	if err = database.DB.Find(permission).Error; err != nil {
		return err
	}
	return nil
}

func GetPermissionById(permission *models.Permission, id string) (err error) {
	if err = database.DB.Where("id = ?", id).First(permission).Error; err != nil {
		return err
	}
	return nil
}

func CreatePermission(permission *models.Permission) (err error) {
	if err = database.DB.Create(&permission).Error; err != nil {
		return err
	}
	return nil
}
