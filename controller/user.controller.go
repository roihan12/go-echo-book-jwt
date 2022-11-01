package controller

import (
	"echo-book/dto"
	"echo-book/helper"
	"echo-book/service"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type UserController interface {
	Update(ctx echo.Context) error
	Profile(ctx echo.Context) error
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}

}

func (c *userController) Update(ctx echo.Context) error {
	var userUpdateDTO dto.UserUpdateDTO
	err := ctx.Bind(&userUpdateDTO)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		return ctx.JSON(http.StatusBadRequest, response)
	}

	if err := ctx.Validate(&userUpdateDTO); err != nil {
		return err
	}

	autheader := ctx.Request().Header.Get("Authorization")
	jwtString := strings.Split(autheader, "Bearer ")[1]
	token, errToken := c.jwtService.ValidateToken(jwtString)
	if errToken != nil {
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	userUpdateDTO.ID = id
	u := c.userService.Update(userUpdateDTO)
	res := helper.BuildResponse(true, "OK!", u)

	return ctx.JSON(http.StatusOK, res)
}

func (c *userController) Profile(ctx echo.Context) error {

	autheader := ctx.Request().Header.Get("Authorization")
	jwtString := strings.Split(autheader, "Bearer ")[1]

	token, errToken := c.jwtService.ValidateToken(jwtString)
	if errToken != nil {
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user := c.userService.Profile(id)
	res := helper.BuildResponse(true, "OK", user)
	return ctx.JSON(http.StatusOK, res)
}
