package bootstrap

import (
	"context"

	"github.com/pujidjayanto/choochoohub/user-api/dto"
	"github.com/pujidjayanto/choochoohub/user-api/pkg/eventbus"
	"github.com/pujidjayanto/choochoohub/user-api/usecase"
)

func RegisterOtpSubscriber(bus eventbus.EventBus, otpUsecase usecase.OtpUsecase) {
	bus.Subscribe("user.created", func(payload any) {
		data, ok := payload.(dto.OtpRequest)
		if !ok {
			return
		}

		_ = otpUsecase.Create(context.Background(), dto.OtpRequest{
			UserId:      data.UserId,
			Channel:     data.Channel,
			Purpose:     data.Purpose,
			Destination: data.Destination,
			ExpiredAt:   data.ExpiredAt,
		})
	})
}
