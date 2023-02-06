package services

import (
	"lms/database"
	"lms/models"
)

func GetConfigByKey(config_key string) string {
	var config_master models.ConfigMaster
	database.DB.Table("config_master").Select("`key`,value, value_type").
		Where("key", config_key).First(&config_master)
	return config_master.Value
}
