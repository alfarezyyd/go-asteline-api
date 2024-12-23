package dto

type UserClaims struct {
	Email string  `mapstructure:"email"`
	Exp   float64 `mapstructure:"exp"`
}
