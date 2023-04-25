package handlers

import (
	"doc-center-api/domain/models"
	"doc-center-api/infra/database"
)

func GetAllFiles(arquivo *[]models.File) (err error) {
	if err = database.DB.Find(arquivo).Error; err != nil {
		return err
	}
	return nil
}

func GetFileByID(file *models.File, id string) (err error) {
	if err = database.DB.Where("id = ?", id).First(file).Error; err != nil {
		return err
	}
	return nil
}

//	func CreateFile(file *models.File) (err error) {
//		if err = database.DB.Create(file).Error; err != nil {
//			return err
//		}
//		return nil
//	}
//
//func UpdateFile(file *models.File, id string) (err error) {
//	fmt.Print(file)
//	database.DB.Save(file)
//	return nil
//}
//
//func DeleteFile(file *models.File, id string) (err error) {
//	database.DB.Where("id = ?", id).Delete(file)
//	return nil
//}