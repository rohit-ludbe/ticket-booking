package handlers

import (
	"context"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rohit-ludbe/ticket-booking-server-v1/models"
)

type EventHandler struct {
	repository models.EventRepository
}

func (h *EventHandler) GetAllEvents(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	events, err := h.repository.GetAllEvents(context)

	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"data":    events,
		"message": "success",
	})
}

func (h *EventHandler) GetOneEvent(ctx *fiber.Ctx) error {

	eventId, _ := strconv.Atoi(ctx.Params("eventId"))

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	event, err := h.repository.GetOneEvent(context, uint(eventId))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"data":    event,
		"message": "success",
	})
}

func (h *EventHandler) CreateOneEvent(ctx *fiber.Ctx) error {

	event := &models.Event{}

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	if err := ctx.BodyParser(event); err != nil {
		return ctx.Status(fiber.ErrUnprocessableEntity.Code).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	event, err := h.repository.CreateOneEvent(context, event)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"data":    event,
		"message": "success",
	})

}

func (h *EventHandler) UpdateOneEvent(ctx *fiber.Ctx) error {

	eventId, _ := strconv.Atoi(ctx.Params("eventId"))

	updateData := make(map[string]interface{})

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	if err := ctx.BodyParser(&updateData); err != nil {
		return ctx.Status(fiber.ErrUnprocessableEntity.Code).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	event, err := h.repository.UpdateOneEvent(context, uint(eventId), updateData)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"data":    event,
		"message": "success",
	})

}

func (h *EventHandler) DeleteOneEvent(ctx *fiber.Ctx) error {

	eventId, _ := strconv.Atoi(ctx.Params("eventId"))

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	err := h.repository.DeleteOneEvent(context, uint(eventId))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"data":    "",
		"message": "success",
	})

}

func NewEventHandler(router fiber.Router, repository models.EventRepository) {
	handler := &EventHandler{
		repository: repository,
	}

	router.Get("/", handler.GetAllEvents)

	router.Post("/", handler.CreateOneEvent)

	router.Get("/:eventId", handler.GetOneEvent)

	router.Put("/:eventId", handler.UpdateOneEvent)

	router.Delete("/:eventId", handler.DeleteOneEvent)

}
