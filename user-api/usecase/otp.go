package usecase

import (
	"context"
	"errors"
	"net/http"

	"github.com/pujidjayanto/choochoohub/user-api/apperror"
	"github.com/pujidjayanto/choochoohub/user-api/dto"
	"github.com/pujidjayanto/choochoohub/user-api/model"
	"github.com/pujidjayanto/choochoohub/user-api/pkg/otpcode"
	"github.com/pujidjayanto/choochoohub/user-api/pkg/stringhash"
	"github.com/pujidjayanto/choochoohub/user-api/repository"
	"gorm.io/gorm"
)

var MaxSignUpOtpSendAttempt = 3

type OtpUsecase interface {
	Create(c context.Context, req dto.OtpRequest) (*model.UserOtp, error)
	VerifyOtp(c context.Context, req dto.VerifyOtpRequest) error
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

func (u *otpUsecase) VerifyOtp(c context.Context, req dto.VerifyOtpRequest) error {
	otp, err := u.otpRepository.FindyByDestinationAndPurpose(c, req.Destination, req.Purpose)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperror.NewAppError(http.StatusNotFound, apperror.OtpNotFound, err)
		}
		return apperror.NewAppError(http.StatusInternalServerError, apperror.CodeInternalServerError, err)
	}

	otp.SendAttempts = otp.SendAttempts + 1
	if otp.SendAttempts > MaxSignUpOtpSendAttempt {
		otp.Status = model.UserOtpStatusMaxAttempted

		err := u.otpRepository.UpdateOtp(c, otp)
		if err != nil {
			return apperror.NewAppError(http.StatusInternalServerError, apperror.CodeInternalServerError, err)
		}

		return apperror.NewAppError(http.StatusUnprocessableEntity, apperror.OtpMaxAttempted, errors.New("otp is max attempeted"))
	}

	isOtpMatch := stringhash.Match(otp.OTPHash, req.OtpCode)
	if !isOtpMatch {
		return apperror.NewAppError(http.StatusUnprocessableEntity, apperror.OtpNotMatch, errors.New("otp not match"))
	}

	otp.Status = model.UserOtpStatusVerified
	err = u.otpRepository.UpdateOtp(c, otp)
	if err != nil {
		return apperror.NewAppError(http.StatusInternalServerError, apperror.CodeInternalServerError, err)
	}

	return nil
}
