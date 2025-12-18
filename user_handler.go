package handler

import (
	"strconv"
	"time"

	"go-user-api/db/sqlc"
	"go-user-api/internal/models"
	"go-user-api/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type UserHandler struct {
	Queries  *sqlc.Queries
	Validate *validator.Validate
	Logger   *zap.Logger
}

// POST /users
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req models.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	if err := h.Validate.Struct(req); err != nil {
		return fiber.ErrBadRequest
	}

	dob, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		return fiber.ErrBadRequest
	}

	user, err := h.Queries.CreateUser(c.Context(), sqlc.CreateUserParams{
		Name: req.Name,
		Dob:  dob,
	})
	if err != nil {
		h.Logger.Error("failed to create user", zap.Error(err))
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

// GET /users/:id
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	user, err := h.Queries.GetUserByID(c.Context(), int32(id))
	if err != nil {
		return fiber.ErrNotFound
	}

	age := service.CalculateAge(user.Dob)

	return c.JSON(fiber.Map{
		"id":   user.ID,
		"name": user.Name,
		"dob":  user.Dob.Format("2006-01-02"),
		"age":  age,
	})
}

// GET /users
func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	users, err := h.Queries.ListUsers(c.Context())
	if err != nil {
		return fiber.ErrInternalServerError
	}

	var response []fiber.Map
	for _, u := range users {
		response = append(response, fiber.Map{
			"id":   u.ID,
			"name": u.Name,
			"dob":  u.Dob.Format("2006-01-02"),
			"age":  service.CalculateAge(u.Dob),
		})
	}

	return c.JSON(response)
}

// PUT /users/:id
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	var req models.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	dob, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		return fiber.ErrBadRequest
	}

	user, err := h.Queries.UpdateUser(c.Context(), sqlc.UpdateUserParams{
		ID:   int32(id),
		Name: req.Name,
		Dob:  dob,
	})
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(user)
}

// DELETE /users/:id
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	err = h.Queries.DeleteUser(c.Context(), int32(id))
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.SendStatus(fiber.StatusNoContent)
}
