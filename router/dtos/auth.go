package dtos

type RefreshToken struct {
	Refresh string `json:"token" query:"token"  valid:"required"`
}

type RefreshTokenResponse struct {
	Token   string `json:"token"`
}