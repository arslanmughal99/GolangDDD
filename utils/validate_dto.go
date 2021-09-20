package utils

import (
	"strings"

	"github.com/asaskevich/govalidator"
)

func ValidateDto(dto interface{}) *[]string {
	valid, err := govalidator.ValidateStruct(dto)

	if !valid {
		res := strings.Split(err.Error(), ";")
		return &res
	}

	return nil
}
