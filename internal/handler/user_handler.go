package handler

import (
	"strconv"

	"github.com/dhofa/gofiber-clean-arch/internal/domain"
	"github.com/dhofa/gofiber-clean-arch/internal/entity"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	usecase domain.UserUsecase
}

func NewUserHandler(u domain.UserUsecase) *UserHandler {
	return &UserHandler{u}
}

func (h *UserHandler) Route(app fiber.Router) {
	app.Get("/", h.GetAll)
	app.Get("/:id", h.GetByID)
	app.Post("/", h.Create)
	app.Put("/:id", h.Update)
	app.Delete("/:id", h.Delete)
}

func (h *UserHandler) GetAll(c *fiber.Ctx) error {
	users, err := h.usecase.GetAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

func (h *UserHandler) GetByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user, err := h.usecase.GetByID(uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(user)
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	var user entity.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := h.usecase.Create(&user); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(user)
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var user entity.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	user.ID = uint(id)
	if err := h.usecase.Update(&user); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(user)
}

func (h *UserHandler) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := h.usecase.Delete(uint(id)); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(204)
}
