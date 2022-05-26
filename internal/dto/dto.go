package dto

type LinkPairCreate struct {
	OriginLink string `json:"originLink" binding:"required,url"`
}
