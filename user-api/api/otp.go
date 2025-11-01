package api

import (
	"github.com/labstack/echo/v4"
	"github.com/pujidjayanto/choochoohub/user-api/dto"
	"github.com/pujidjayanto/choochoohub/user-api/pkg/delivery"
	"github.com/pujidjayanto/choochoohub/user-api/usecase"
)

type OtpApi interface {
	Verify(c echo.Context) error
}

type otpApi struct {
	otpUsecase usecase.OtpUsecase
}

func NewOtpApi(otpUsecase usecase.OtpUsecase) OtpApi {
	return &otpApi{otpUsecase}
}

func (otpApi *otpApi) Verify(c echo.Context) error {
	var req dto.VerifyOtpRequest
	if err := BindAndValidate(c, &req); err != nil {
		return delivery.Failed(c, err)
	}

	err := otpApi.otpUsecase.VerifyOtp(c.Request().Context(), req)
	if err != nil {
		return delivery.Failed(c, err)
	}

	return delivery.SuccessNoContent(c)
}
