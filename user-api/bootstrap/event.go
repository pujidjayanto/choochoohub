package bootstrap

import (
	"context"
	"encoding/json"

	"github.com/pujidjayanto/choochoohub/user-api/dto"
	"github.com/pujidjayanto/choochoohub/user-api/pkg"
	"github.com/pujidjayanto/choochoohub/user-api/usecase"
)

func RegisterOtpSubscriber(shared *pkg.Dependency, otpUsecase usecase.OtpUsecase) {
	shared.EventBus.Subscribe("user.created", func(payload any) {
		data, ok := payload.(dto.OtpRequest)
		if !ok {
			return
		}

		otp, err := otpUsecase.Create(context.Background(), dto.OtpRequest{
			UserId:      data.UserId,
			Channel:     data.Channel,
			Purpose:     data.Purpose,
			Destination: data.Destination,
			ExpiredAt:   data.ExpiredAt,
		})
		if err != nil {
			shared.Logger.WithField("err", err).Info("failed to create otp")
			return
		}

		eventData, _ := json.Marshal(map[string]any{
			"destination": otp.Destination,
			"channel":     otp.Channel,
			"purpose":     otp.Purpose,
			"expiresAt":   otp.ExpiresAt,
			"otpId":       otp.ID,
			"code":        otp.OTPCode,
		})

		err = shared.KafkaProducer.SendMessage(context.Background(), "otp.created", 0, eventData)
		if err != nil {
			shared.Logger.WithField("err", err).Info("failed to send kafka message")
		}
	})
}
