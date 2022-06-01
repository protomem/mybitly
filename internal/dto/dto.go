package dto

type LinkPairCreate struct {
	OriginLink string `json:"originLink" binding:"required,url"`
}

type UserCreate struct {
	Nickname string `json:"nickname" binding:"required"`
	Password string `json:"password" binding:"required,min=6,max=18"`
	Email    string `json:"email"    binding:"required,email"`
}
