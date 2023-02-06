package services

import (
	"lms/database"
	"lms/models"
)

func GetProductName(product_type string, id int64) string {
	var result models.ProductType
	database.DB.Table(product_type+"_master").
		Select("id, name").
		Where("id", id).
		First(&result)
	return result.Name
}
