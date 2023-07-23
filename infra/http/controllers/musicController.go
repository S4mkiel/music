package controller

import (
	"github.com/S4mkiel/music/domain/entity"
	"github.com/S4mkiel/music/domain/service"
	"github.com/S4mkiel/music/infra/http/dto"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type MusicController struct {
	logger      *zap.SugaredLogger
	UserService *service.UserService
}

func NewMusicController(logger *zap.SugaredLogger, userService *service.UserService) *MusicController {
	return &MusicController{logger: logger, UserService: userService}
}

func (c MusicController) RegisterRoutes(app fiber.Router) {
	music := app.Group("/music")
	music.Post("/create", c.CreateUser)
}

// Create a new user
//
// @Summary Create a new user
// @Description Create a new user in the database
// @Tags music
// @Accept json
// @Produce json
// @Param email body string true "Email"
// @Param name body string true "Name"
// @Param date body string true "Date"
// @Param city body string true "City"
// @Param contact body string true "Contact"
// @Success 201 {object} dto.BaseSuccess "Success"
// @Failure 400 {object} dto.BaseFailure "Invalid request"
// @Router /music/create [post]
func (c MusicController) CreateUser(ctx *fiber.Ctx) error {
	response := dto.BaseSuccess{
		Message: "",
		Data:    nil,
	}
	failureResponse := dto.BaseFailure{
		Message: "",
		Error:   "",
	}

	type Payload struct {
		Email   string `json:"email"`
		Name    string `json:"name"`
		Date    string `json:"birth_date"`
		City    string `json:"city"`
		Contact string `json:"contact"`
	}

	var request Payload
	if err := ctx.BodyParser(&request); err != nil {
		failureResponse.Error = err.Error()
		failureResponse.Message = "Invalid request"
		return ctx.Status(fiber.StatusBadRequest).JSON(failureResponse)
	}

	u := entity.User{
		Email:   request.Email,
		Name:    request.Name,
		Date:    request.Date,
		City:    request.City,
		Contact: request.Contact,
	}

	user, err := c.UserService.Create(ctx.Context(), u)

	if err != nil {
		failureResponse.Error = err.Error()
		failureResponse.Message = "Invalid request"
		return ctx.Status(fiber.StatusBadRequest).JSON(failureResponse)
	}

	response.Data = user

	return ctx.Status(fiber.StatusCreated).JSON(response)
}
