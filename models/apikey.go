package models

import (
	"github.com/carbans/netdebug/types"
	"github.com/carbans/netdebug/utils/token"
	"gorm.io/gorm"
)

type ApiKey struct {
	gorm.Model
	Key    string `gorm:"unique_index;not null"`
	Secret string `gorm:"not null"`
	UserId uint   `gorm:"not null"`
}

func CreateApiKey(key types.CreateApiKeyInput, user_id uint) (ApiKey, error) {

	secret, e := token.GenerateToken(user_id)

	if e != nil {
		return ApiKey{}, e
	}

	apikey := ApiKey{Key: key.Key, Secret: secret, UserId: user_id}
	var err = DB.Create(&apikey).Error
	if err != nil {
		return ApiKey{}, err
	}
	return apikey, nil
}

func GetApiKeyById(id uint) (ApiKey, error) {
	var key ApiKey
	var err = DB.First(&key, id).Error
	if err != nil {
		return key, err
	}
	return key, nil
}

func GetApiKeyByUserId(userId uint) (ApiKey, error) {
	var key ApiKey
	var err = DB.Where("user_id = ?", userId).First(&key).Error
	if err != nil {
		return key, err
	}
	return key, nil
}
