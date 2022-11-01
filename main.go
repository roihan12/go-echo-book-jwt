package main

import (
	"echo-book/config"
	"echo-book/controller"
	"echo-book/helper"
	"echo-book/middleware"
	"echo-book/repository"
	"echo-book/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	bookRepository repository.BookRepository = repository.NewBookRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	bookService    service.BookService       = service.NewBookService(bookRepository)
	authService    service.AuthService       = service.NewAuthService(userRepository)
	userService    service.UserService       = service.NewUserService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
	bookController controller.BookController = controller.NewBookController(bookService, jwtService)
)

func main() {

	defer config.CloseDatabaseConnection(db)
	e := echo.New()

	// e.Validator = &helper.CustomValidator{Validator: validator.New()}
	e.Validator = &helper.CustomValidator{Validator: validator.New()}
	authRoute := e.Group("api/auth")

	authRoute.POST("/login", authController.Login)

	authRoute.POST("/register", authController.Register)

	userRoute := e.Group("api/users", middleware.AuthorizeJWT)

	userRoute.GET("/profile", userController.Profile)
	userRoute.PUT("/profile", userController.Update)

	bookRoute := e.Group("api/books", middleware.AuthorizeJWT)
	bookRoute.GET("", bookController.All)
	bookRoute.POST("", bookController.Insert)
	bookRoute.GET("/:id", bookController.FindByID)
	bookRoute.PUT("/:id", bookController.Update)
	bookRoute.DELETE("/:id", bookController.Delete)
	e.Logger.Fatal(e.Start(":1323"))
}
