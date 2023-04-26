package types

type CreateAgentType struct {
	Name    string `json:"name" binding:"required"`
	Version string `json:"version"`
}

type AgentResponseType struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	UserID  uint   `json:"user_id"`
}
