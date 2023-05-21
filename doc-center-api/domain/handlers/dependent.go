package handlers

import (
	"doc-center-api/domain/models"
	"doc-center-api/infra/database"
)

func GetAllDependents(dependent *[]models.Dependent) (err error) {
	if err = database.DB.Joins("Responsable").
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

func CreateDependent(dependent *models.Dependent) (err error) {
	if err = database.DB.Create(&dependent).Error; err != nil {
		return err
	}
	return nil
}

func DeleteDependent(dependentId string) (err error) {
	if err = database.DB.Delete(models.Dependent{}, dependentId).Error; err != nil {
		return err
	}
	return nil
}

func UpdateDependent(dependent *models.Dependent) (err error) {
	if err = database.DB.Updates(&dependent).Error; err != nil {
		return err
	}
	return nil
}
