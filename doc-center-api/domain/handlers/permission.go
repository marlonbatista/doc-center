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

func GetOwnerPermission(id string, permission *[]models.Permission) (err error) {
	if err = database.DB.Where("owner_id = ?", id).Find(&permission).Error; err != nil {
		return err
	}
	return nil
}

func GetGuestPermission(id string, permission *[]models.Permission) (err error) {
	if err = database.DB.Where("guest_id = ?", id).Find(&permission).Error; err != nil {
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

func UpdatePermission(permission *models.Permission) (err error) {
	if err = database.DB.Save(&permission).Error; err != nil {
		return err
	}
	return nil
}

func DeletePermission(id string, permission *models.Permission) (err error) {
	if err = database.DB.Delete(&permission, id).Error; err != nil {
		return err
	}
	return nil
}
