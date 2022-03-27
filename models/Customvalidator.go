package models

import "github.com/go-playground/validator"

type CustomValidator struct {
	validator validator.Validate
}
