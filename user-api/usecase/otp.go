package usecase

import (
	"context"
	"net/http"

	"github.com/pujidjayanto/choochoohub/user-api/apperror"
	"github.com/pujidjayanto/choochoohub/user-api/dto"
	"github.com/pujidjayanto/choochoohub/user-api/model"
	"github.com/pujidjayanto/choochoohub/user-api/pkg/otpcode"
	"github.com/pujidjayanto/choochoohub/user-api/pkg/stringhash"
	"github.com/pujidjayanto/choochoohub/user-api/repository"
)

type OtpUsecase interface {
	Create(c context.Context, req dto.OtpRequest) (*model.UserOtp, error)
}

type otpUsecase struct {
	otpRepository repository.UserOtpRepository
}

func NewOtpUsecase(otpRepository repository.UserOtpRepository) OtpUsecase {
	return &otpUsecase{otpRepository}
}

func (u *otpUsecase) Create(c context.Context, req dto.OtpRequest) (*model.UserOtp, error) {
	otpCode, err := otpcode.Generate()
	if err != nil {
		return nil, apperror.NewAppError(http.StatusInternalServerError, apperror.CodeInternalServerError, err)
	}

	otpHashed, err := stringhash.Hash(otpCode)
	if err != nil {
		return nil, apperror.NewAppError(http.StatusInternalServerError, apperror.CodeInternalServerError, err)
	}

	otp := model.UserOtp{
		UserID:      req.UserId,
		Channel:     model.UserOtpChannel(req.Channel),
		Purpose:     model.UserOtpPurpose(req.Purpose),
		Destination: req.Destination,
		ExpiresAt:   req.ExpiredAt,
		OTPHash:     otpHashed,
	}

	createdOtp, err := u.otpRepository.Create(c, &otp)
	if err != nil {
		return nil, apperror.NewAppError(http.StatusInternalServerError, apperror.CodeInternalServerError, err)
	}

	createdOtp.OTPCode = otpCode

	return createdOtp, nil
}
