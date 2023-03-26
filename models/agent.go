package models

import (
	"gorm.io/gorm"
)

type Agent struct {
	gorm.Model
	Name    string `json:"name"`
	Version string `json:"version"`
	UserID  uint   `json:"user_id"`
}

func GetAgents(user_id uint) []Agent {
	var agents []Agent

	DB.Where("user_id == ?", user_id).Find(&agents)

	return agents
}

func GetAgent(user_id uint, id int) Agent {
	var agent Agent

	DB.Where("id = ? AND user_id = ?", id, user_id).First(&agent)
	return agent
}
