package models

import "gorm.io/gorm"

type ApiKey struct {
	gorm.Model
	Key    string `gorm:"unique_index;not null"`
	Secret string `gorm:"not null"`
	UserId uint   `gorm:"not null"`
}

func CreateApiKey(key ApiKey) (ApiKey, error) {
	var err = DB.Create(&key).Error
	if err != nil {
		return key, err
	}
	return key, nil
}

func GetApiKeyById(id uint) (ApiKey, error) {
	var key ApiKey
	var err = DB.First(&key, id).Error
	if err != nil {
		return key, err
	}
	return key, nil
}
