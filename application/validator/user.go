package validator

import (
	"errors"
	"github.com/go-playground/validator/v10"
	. "urils/application/model"
	. "urils/application/utils"
)

/**
登陆验证器
*/

func LoginValidator(user *User) error {
	validate, trans := GenValidate()
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return errors.New(err.Translate(trans))
		}
	}
	return nil
}
