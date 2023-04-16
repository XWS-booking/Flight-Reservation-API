package dtos

type JwtDto struct {
	Token string `json:"access_token"`
}

func NewJwtDto(token string) *JwtDto {
	return &JwtDto{Token: token}
}
