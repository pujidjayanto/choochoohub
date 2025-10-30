// https://echo.labstack.com/docs/request#validate-data
package validator

import "github.com/go-playground/validator/v10"

type CustomValidator struct {
	validator *validator.Validate
}

func New() *CustomValidator {
	v := validator.New(validator.WithRequiredStructEnabled())
	return &CustomValidator{validator: v}
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}
