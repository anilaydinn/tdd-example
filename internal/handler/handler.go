package handler

import (
	"github.com/anilaydinn/tdd-example/internal/model"
	"github.com/gofiber/fiber/v2"
)

type StudentActions interface {
	CreateStudent(request model.CreateStudentRequest) (model.CreateStudentResponse, error)
}

type Handler struct {
	studentActions StudentActions
}

func NewHandler(studentActions StudentActions) *Handler {
	return &Handler{
		studentActions: studentActions,
	}
}

func (h *Handler) RegisterRoutes(app *fiber.App) {
	app.Post("/students", h.CreateStudent)
}

func (h *Handler) CreateStudent(c *fiber.Ctx) error {
	var requestBody model.CreateStudentRequest
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response, err := h.studentActions.CreateStudent(requestBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
