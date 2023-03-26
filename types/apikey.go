package types

type CreateApiKeyInput struct {
	Key string `json:"key" binding:"required"`
}
