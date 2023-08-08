package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/anilaydinn/tdd-example/internal/mock"
	"github.com/anilaydinn/tdd-example/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_CreateStudent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	requestBody := model.CreateStudentRequest{
		Name:    "Anil",
		Surname: "Aydin",
		Age:     25,
		Gender:  "Male",
	}

	t.Run("it should return 201 when student created successfully", func(t *testing.T) {
		mockStudentActions := mock.NewMockStudentActions(ctrl)

		mockStudentActions.EXPECT().CreateStudent(requestBody).Return(model.CreateStudentResponse{
			ID:      1,
			Name:    requestBody.Name,
			Surname: requestBody.Surname,
			Age:     requestBody.Age,
			Gender:  requestBody.Gender,
		}, nil)

		body, err := json.Marshal(requestBody)
		assert.Nil(t, err)

		req := httptest.NewRequest(http.MethodPost, "/students", bytes.NewReader(body))
		req.Header.Add("Content-Type", "application/json")

		app := fiber.New()

		handler := NewHandler(mockStudentActions)
		handler.RegisterRoutes(app)

		res, err := app.Test(req, 30000)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})

	t.Run("it should return 400 when request body is invalid", func(t *testing.T) {
		mockStudentActions := mock.NewMockStudentActions(ctrl)

		body, err := json.Marshal("invalid")
		assert.Nil(t, err)

		req := httptest.NewRequest(http.MethodPost, "/students", bytes.NewReader(body))
		req.Header.Add("Content-Type", "application/json")

		app := fiber.New()

		handler := NewHandler(mockStudentActions)
		handler.RegisterRoutes(app)

		res, err := app.Test(req, 30000)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("it should return 500 when student creation failed", func(t *testing.T) {
		mockStudentActions := mock.NewMockStudentActions(ctrl)

		mockStudentActions.EXPECT().CreateStudent(requestBody).Return(model.CreateStudentResponse{}, assert.AnError)

		body, err := json.Marshal(requestBody)
		assert.Nil(t, err)

		req := httptest.NewRequest(http.MethodPost, "/students", bytes.NewReader(body))
		req.Header.Add("Content-Type", "application/json")

		app := fiber.New()

		handler := NewHandler(mockStudentActions)
		handler.RegisterRoutes(app)

		res, err := app.Test(req, 30000)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	})
}
