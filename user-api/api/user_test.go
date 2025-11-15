package api_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/matryer/is"

	"github.com/pujidjayanto/choochoohub/user-api/api"
	"github.com/pujidjayanto/choochoohub/user-api/apperror"
	"github.com/pujidjayanto/choochoohub/user-api/dto"
	"github.com/pujidjayanto/choochoohub/user-api/mocks"
	"github.com/pujidjayanto/choochoohub/user-api/pkg/validator"
)

func TestUserApi_SignUp(t *testing.T) {
	is := is.New(t)

	t.Run("success", func(t *testing.T) {
		mockUC := &mocks.UserUsecaseMock{
			SignupFunc: func(_ context.Context, req dto.SignupRequest) error {
				return nil
			},
		}

		apiHandler := api.NewUserApi(mockUC)

		e := echo.New()
		e.Validator = validator.New()

		body, _ := json.Marshal(dto.SignupRequest{
			Email:    "test@example.com",
			Password: "secure123",
		})
		req := httptest.NewRequest(http.MethodPost, "/v1/signup", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		err := apiHandler.SignUp(c)
		is.NoErr(err)
		is.Equal(rec.Code, http.StatusNoContent)
	})

	t.Run("bad request", func(t *testing.T) {
		mockUC := &mocks.UserUsecaseMock{}

		apiHandler := api.NewUserApi(mockUC)
		e := echo.New()
		e.Validator = validator.New()

		// invalid JSON
		req := httptest.NewRequest(http.MethodPost, "/v1/signup", bytes.NewReader([]byte(`invalid`)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		err := apiHandler.SignUp(c)
		is.NoErr(err)
		is.Equal(rec.Code, http.StatusBadRequest)
	})

	t.Run("duplicate email", func(t *testing.T) {
		mockUC := &mocks.UserUsecaseMock{
			SignupFunc: func(_ context.Context, req dto.SignupRequest) error {
				return apperror.NewAppError(http.StatusUnprocessableEntity, apperror.CodeValidationFailed,
					errors.New("email already exists"))
			},
		}

		apiHandler := api.NewUserApi(mockUC)
		e := echo.New()
		e.Validator = validator.New()

		body, _ := json.Marshal(dto.SignupRequest{
			Email:    "exists@example.com",
			Password: "secure123",
		})
		req := httptest.NewRequest(http.MethodPost, "/v1/signup", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		err := apiHandler.SignUp(c)
		is.NoErr(err)
		is.Equal(rec.Code, http.StatusUnprocessableEntity)
	})
}
