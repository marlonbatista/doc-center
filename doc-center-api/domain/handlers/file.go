package handlers

import (
	"doc-center-api/domain/models"
	"doc-center-api/infra/database"
)

func GetAllFiles(userId string) (file []models.File, err error) {
	var files []models.File
	if err = database.DB.Where("user_id = ?", userId).Find(&files).Error; err != nil {
		return files, err
	}
	return files, nil
}

func GetFileByID(file *models.File, id string) (err error) {
	if err = database.DB.Where("id = ?", id).First(file).Error; err != nil {
		return err
	}
	return nil
}

func CreateFile(file *models.File) (err error) {
	if err = database.DB.Create(&file).Error; err != nil {
		return err
	}
	return nil
}

func UpdateFile(file *models.File) (err error) {
	if err = database.DB.Save(&file).Error; err != nil {
		return err
	}
	return nil
}

func DeleteFile(file *models.File, id string) (err error) {
	if err = database.DB.Delete(&file, id).Error; err != nil {
		return err
	}
	return nil
}
