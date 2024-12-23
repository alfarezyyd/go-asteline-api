package mapper

import (
	"fmt"
	"github.com/go-viper/mapstructure/v2"
	"github.com/golang-jwt/jwt/v5"
	"go-asteline-api/model"
	"go-asteline-api/user/dto"
	"time"
)

func MapUserRegisterDtoIntoUserModel(userRegisterDto dto.UserRegisterDto) (*model.User, error) {
	// Konversi struct dto ke map
	var userMap map[string]interface{}
	err := mapstructure.Decode(userRegisterDto, &userMap)
	if err != nil {
		return nil, err
	}

	// Custom handling untuk field yang membutuhkan konversi (misalnya BirthDate)
	birthDate, err := time.Parse("2006-01-02", userRegisterDto.BirthDate)
	if err != nil {
		return nil, err
	}

	// Tambahkan konversi manual
	userMap["BirthDate"] = birthDate

	// Decode kembali ke struct User
	var user model.User
	fmt.Println(userMap)
	err = mapstructure.Decode(userMap, &user)
	if err != nil {
		return nil, err
	}
	fmt.Println(user)
	return &user, nil
}

func MapJwtClaimIntoUserClaim(jwtClaim jwt.MapClaims) (*dto.UserClaims, error) {
	var userClaim dto.UserClaims
	err := mapstructure.Decode(jwtClaim, &userClaim)
	if err != nil {
		return nil, err
	}
	return &userClaim, nil
}
