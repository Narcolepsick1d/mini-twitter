package validator

import (
	"Narcolepsick1d/mini-twitter/internal/models"
	"github.com/asaskevich/govalidator"
)

func ValidateUserSignUp(user models.UserSignUp) bool {
	if len(user.Name) == 0 || len(user.Email) == 0 || len(user.Password) == 0 {
		return false
	}
	return govalidator.IsEmail(user.Email) && (len(user.Name) < 20)
}
func ValidateUserSignIn(user models.User) bool {
	if len(user.Email) == 0 || len(user.Password) == 0 {
		return false
	}
	return govalidator.IsEmail(user.Email)
}
