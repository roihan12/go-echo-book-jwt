package controller

import (
	"echo-book/dto"
	"echo-book/entity"
	"echo-book/helper"
	"echo-book/service"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type BookController interface {
	All(ctx echo.Context) error
	FindByID(ctx echo.Context) error
	Insert(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type bookController struct {
	bookService service.BookService
	jwtService  service.JWTService
}

// NewBookController create a new instances of BoookController
func NewBookController(bookServ service.BookService, jwtServ service.JWTService) BookController {
	return &bookController{
		bookService: bookServ,
		jwtService:  jwtServ,
	}
}

func (c *bookController) All(ctx echo.Context) error {
	var books []entity.Book = c.bookService.All()
	res := helper.BuildResponse(true, "OK", books)

	return ctx.JSON(http.StatusOK, res)
}

func (c *bookController) FindByID(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		return ctx.JSON(http.StatusBadRequest, res)
	}

	var book entity.Book = c.bookService.FindByID(id)
	if (book == entity.Book{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		return ctx.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", book)
		return ctx.JSON(http.StatusOK, res)
	}
}

func (c *bookController) Insert(ctx echo.Context) error {
	var bookCreateDTO dto.BookCreateDTO
	err := ctx.Bind(&bookCreateDTO)

	if err := ctx.Validate(&bookCreateDTO); err != nil {
		return err
	}

	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		return ctx.JSON(http.StatusBadRequest, response)
	} else {

		autheader := ctx.Request().Header.Get("Authorization")
		jwtString := strings.Split(autheader, "Bearer ")[1]
		userID := c.getUserIDByToken(jwtString)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			bookCreateDTO.UserID = convertedUserID
		}
		result := c.bookService.Insert(bookCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		return ctx.JSON(http.StatusCreated, response)
	}

}

func (c *bookController) Update(ctx echo.Context) error {
	var bookUpdateDTO dto.BookUpdateDTO
	err := ctx.Bind(&bookUpdateDTO)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		return ctx.JSON(http.StatusBadRequest, res)

	}

	if err := ctx.Validate(&bookUpdateDTO); err != nil {
		return err
	}

	autheader := ctx.Request().Header.Get("Authorization")
	jwtString := strings.Split(autheader, "Bearer ")[1]
	token, errToken := c.jwtService.ValidateToken(jwtString)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.bookService.IsAllowedToEdit(userID, bookUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			bookUpdateDTO.UserID = id
		}
		result := c.bookService.Update(bookUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		return ctx.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		return ctx.JSON(http.StatusForbidden, response)
	}
}

func (c *bookController) Delete(ctx echo.Context) error {
	var book entity.Book
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		return ctx.JSON(http.StatusBadRequest, response)
	}
	book.ID = id
	autheader := ctx.Request().Header.Get("Authorization")
	jwtString := strings.Split(autheader, "Bearer ")[1]
	token, errToken := c.jwtService.ValidateToken(jwtString)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.bookService.IsAllowedToEdit(userID, book.ID) {
		c.bookService.Delete(book)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		return ctx.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		return ctx.JSON(http.StatusForbidden, response)
	}
}

func (c *bookController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}

	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id

}
