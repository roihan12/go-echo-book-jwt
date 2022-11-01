package controller

import (
	"echo-book/dto"
	"echo-book/entity"
	"echo-book/helper"
	"echo-book/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AuthController interface {
	Login(ctx echo.Context) error
	Register(ctx echo.Context) error
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

// NewAuthController adalah membuat instance baru dari AuthController
func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx echo.Context) error {
	var loginDto dto.LoginDTO
	err := ctx.Bind(&loginDto)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		return ctx.JSON(http.StatusBadRequest, response)
	}

	if err := ctx.Validate(&loginDto); err != nil {
		return err
	}

	authResult := c.authService.VerifyCredential(loginDto.Email, loginDto.Password)
	if v, ok := authResult.(entity.User); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generatedToken
		response := helper.BuildResponse(true, "Ok!", v)
		return ctx.JSON(http.StatusOK, response)

	}

	response := helper.BuildErrorResponse("Please check again your email or password", "invalid credential", helper.EmptyObj{})
	return ctx.JSON(http.StatusUnauthorized, response)

}

func (c *authController) Register(ctx echo.Context) error {
	var registerDto dto.RegisterDTO
	err := ctx.Bind(&registerDto)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		return ctx.JSON(http.StatusBadRequest, response)
	}

	if err := ctx.Validate(&registerDto); err != nil {
		return err
	}

	if !c.authService.IsDuplicateEmail(registerDto.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "duplicate email", helper.EmptyObj{})
		return ctx.JSON(http.StatusBadRequest, response)
	} else {
		createdUser := c.authService.CreateUser(registerDto)
		token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
		createdUser.Token = token
		response := helper.BuildResponse(true, "Ok!", createdUser)
		return ctx.JSON(http.StatusCreated, response)

	}
}
