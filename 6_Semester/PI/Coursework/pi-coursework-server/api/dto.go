package api

type QeuryDTO struct {
	Query string `json:"query" binding:"required"`
	Token string `form:"token" binding:"required"`
}

type AuthDTO struct {
	Login    string `json:"login" binding:"required"`
	Password string `form:"password" binding:"required"`
}
