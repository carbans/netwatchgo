package types

type LoginType struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,min=3,max=50"`
}

type LoginResponseType struct {
	Token string `json:"token"`
}
