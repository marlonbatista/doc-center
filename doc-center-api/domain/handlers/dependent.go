package handlers

import (
	"doc-center-api/domain/models"
	"doc-center-api/infra/database"
)

func GetAllDependents(dependent *[]models.Dependent) (err error) {
	if err = database.DB.Joins("JOIN users ON users.id = dependents.responsable_id").
		Find(dependent).Error; err != nil {
		return err
	}
	return nil
}

func GetDependentById(dependent *models.Dependent, id string) (err error) {
	if err = database.DB.Where("id = ?", id).First(dependent).Error; err != nil {
		return err
	}
	return nil
}
