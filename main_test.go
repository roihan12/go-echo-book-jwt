package main_test

import (
	"echo-book/config"
	"echo-book/controller"
	"echo-book/dto"
	"echo-book/entity"
	"echo-book/helper"
	"echo-book/middleware"
	"echo-book/repository"
	"echo-book/service"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/steinfletcher/apitest"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	dbTest             *gorm.DB                  = config.SetupDatabaseConnectionTest()
	userRepositoryTest repository.UserRepository = repository.NewUserRepository(dbTest)
	bookRepositoryTest repository.BookRepository = repository.NewBookRepository(dbTest)
	jwtServiceTest     service.JWTService        = service.NewJWTService()
	bookServiceTest    service.BookService       = service.NewBookService(bookRepositoryTest)
	authServiceTest    service.AuthService       = service.NewAuthService(userRepositoryTest)
	userServiceTest    service.UserService       = service.NewUserService(userRepositoryTest)
	authControllerTest controller.AuthController = controller.NewAuthController(authServiceTest, jwtServiceTest)
	userControllerTest controller.UserController = controller.NewUserController(userServiceTest, jwtServiceTest)
	bookControllerTest controller.BookController = controller.NewBookController(bookServiceTest, jwtServiceTest)
)

func SeedUser() entity.User {
	password, _ := bcrypt.GenerateFromPassword([]byte("123123"), bcrypt.DefaultCost)

	var user entity.User = entity.User{
		Name:     "Test",
		Email:    "testing@mail.com",
		Password: string(password),
	}

	if err := dbTest.Create(&user).Error; err != nil {
		panic(err)
	}

	var createdUser entity.User

	dbTest.Last(&createdUser)

	createdUser.Password = "123123"

	return createdUser
}

func SeedBook() entity.Book {

	var note entity.Book = entity.Book{
		Title:       "test",
		Description: "test",
	}

	if err := dbTest.Create(&note).Error; err != nil {
		panic(err)
	}

	var createdNote entity.Book

	dbTest.Last(&createdNote)

	return createdNote
}

func CleanSeeders(dbTest *gorm.DB) {
	dbTest.Exec("SET FOREIGN_KEY_CHECKS = 0")

	// categoryResult := dbTest.Exec("DELETE FROM categories")
	itemResult := dbTest.Exec("DELETE FROM books")
	userResult := dbTest.Exec("DELETE FROM users")

	var isFailed bool = itemResult.Error != nil || userResult.Error != nil
	if isFailed {
		panic(errors.New("error when cleaning up seeders"))
	}

	log.Println("Seeders are cleaned up successfully")
}

func newApp() *echo.Echo {

	defer config.CloseDatabaseConnection(dbTest)
	e := echo.New()

	// e.Validator = &helper.CustomValidator{Validator: validator.New()}
	e.Validator = &helper.CustomValidator{Validator: validator.New()}
	authRoute := e.Group("api/auth")

	authRoute.POST("/login", authControllerTest.Login)

	authRoute.POST("/register", authControllerTest.Register)

	userRoute := e.Group("api/users", middleware.AuthorizeJWT)

	userRoute.GET("/profile", userControllerTest.Profile)
	userRoute.PUT("/profile", userControllerTest.Update)

	bookRoute := e.Group("api/books", middleware.AuthorizeJWT)
	bookRoute.GET("", bookControllerTest.All)
	bookRoute.POST("", bookControllerTest.Insert)
	bookRoute.GET("/:id", bookControllerTest.FindByID)
	bookRoute.PUT("/:id", bookControllerTest.Update)
	bookRoute.DELETE("/:id", bookControllerTest.Delete)
	e.Logger.Fatal(e.Start(":1323"))

	return e
}

func cleanup(res *http.Response, req *http.Request, apiTest *apitest.APITest) {
	if http.StatusOK == res.StatusCode || http.StatusCreated == res.StatusCode {
		// configDB := _dbDriver.ConfigDB{
		// 	DB_USERNAME: util.GetConfig("DB_USERNAME"),
		// 	DB_PASSWORD: util.GetConfig("DB_PASSWORD"),
		// 	DB_HOST:     util.GetConfig("DB_HOST"),
		// 	DB_PORT:     util.GetConfig("DB_PORT"),
		// 	DB_NAME:     util.GetConfig("DB_TEST_NAME"),
		// }

		// db := configDB.InitDB()

		CleanSeeders(dbTest)
	}
}

func getJWTToken(t *testing.T) string {
	// configDB := _dbDriver.ConfigDB{
	// 	DB_USERNAME: util.GetConfig("DB_USERNAME"),
	// 	DB_PASSWORD: util.GetConfig("DB_PASSWORD"),
	// 	DB_HOST:     util.GetConfig("DB_HOST"),
	// 	DB_PORT:     util.GetConfig("DB_PORT"),
	// 	DB_NAME:     util.GetConfig("DB_TEST_NAME"),
	// }

	// db := configDB.InitDB()

	user := SeedUser()

	var userRequest *dto.LoginDTO = &dto.LoginDTO{
		Email:    user.Email,
		Password: user.Password,
	}

	var resp *http.Response = apitest.New().
		Handler(newApp()).
		Post("/api/auth/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusOK).
		End().Response

	var response map[string]string = map[string]string{}

	json.NewDecoder(resp.Body).Decode(&response)

	var token string = response["token"]

	var JWT_TOKEN = "Bearer " + token

	return JWT_TOKEN
}

func getUser() entity.User {
	// configDB := _dbDriver.ConfigDB{
	// 	DB_USERNAME: util.GetConfig("DB_USERNAME"),
	// 	DB_PASSWORD: util.GetConfig("DB_PASSWORD"),
	// 	DB_HOST:     util.GetConfig("DB_HOST"),
	// 	DB_PORT:     util.GetConfig("DB_PORT"),
	// 	DB_NAME:     util.GetConfig("DB_TEST_NAME"),
	// }

	// db := configDB.InitDB()

	user := SeedUser()

	return user
}

func getNote() entity.Book {
	// configDB := _dbDriver.ConfigDB{
	// 	DB_USERNAME: util.GetConfig("DB_USERNAME"),
	// 	DB_PASSWORD: util.GetConfig("DB_PASSWORD"),
	// 	DB_HOST:     util.GetConfig("DB_HOST"),
	// 	DB_PORT:     util.GetConfig("DB_PORT"),
	// 	DB_NAME:     util.GetConfig("DB_TEST_NAME"),
	// }

	// db := configDB.InitDB()

	book := SeedBook()

	return book
}

func TestRegister_Success(t *testing.T) {
	var userRequest *dto.RegisterDTO = &dto.RegisterDTO{
		Name:     "test",
		Email:    "test@mail.com",
		Password: "123123",
	}

	apitest.New().
		Observe(cleanup).
		Handler(newApp()).
		Post("/api/auth/register").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusCreated).
		End()
}

func TestRegister_ValidationFailed(t *testing.T) {
	var userRequest *dto.RegisterDTO = &dto.RegisterDTO{
		Name:     "",
		Email:    "",
		Password: "",
	}

	apitest.New().
		Handler(newApp()).
		Post("/api/v1/users/register").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestLogin_Success(t *testing.T) {
	user := getUser()

	var userRequest *dto.LoginDTO = &dto.LoginDTO{
		Email:    user.Email,
		Password: user.Password,
	}

	apitest.New().
		Handler(newApp()).
		Post("/api/auth/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestLogin_ValidationFailed(t *testing.T) {
	var userRequest *dto.LoginDTO = &dto.LoginDTO{
		Email:    "",
		Password: "",
	}

	apitest.New().
		Handler(newApp()).
		Post("/api/auth/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestLogin_Failed(t *testing.T) {
	var userRequest *dto.LoginDTO = &dto.LoginDTO{
		Email:    "notfound@mail.com",
		Password: "123123",
	}

	apitest.New().
		Handler(newApp()).
		Post("/api/auth/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusUnauthorized).
		End()
}

// func TestGetNotes_Success(t *testing.T) {
// 	var token string = getJWTToken(t)

// 	apitest.New().
// 		Observe(cleanup).
// 		Handler(newApp()).
// 		Get("/api/v1/notes").
// 		Header("Authorization", token).
// 		Expect(t).
// 		Status(http.StatusOK).
// 		End()
// }

// func TestGetNote_Success(t *testing.T) {
// 	var note notes.Note = getNote()

// 	noteID := strconv.Itoa(int(note.ID))

// 	var token string = getJWTToken(t)

// 	apitest.New().
// 		Observe(cleanup).
// 		Handler(newApp()).
// 		Get("/api/v1/notes/"+noteID).
// 		Header("Authorization", token).
// 		Expect(t).
// 		Status(http.StatusOK).
// 		End()
// }

// func TestGetNote_NotFound(t *testing.T) {
// 	var token string = getJWTToken(t)

// 	apitest.New().
// 		Handler(newApp()).
// 		Get("/api/v1/notes/0").
// 		Header("Authorization", token).
// 		Expect(t).
// 		Status(http.StatusNotFound).
// 		End()
// }

// func TestCreateNote_Success(t *testing.T) {
// 	category := getCategory()

// 	var noteRequest *notes.Note = &notes.Note{
// 		Title:      "test",
// 		Content:    "test",
// 		CategoryID: category.ID,
// 	}

// 	var token string = getJWTToken(t)

// 	apitest.New().
// 		Observe(cleanup).
// 		Handler(newApp()).
// 		Post("/api/v1/notes").
// 		Header("Authorization", token).
// 		JSON(noteRequest).
// 		Expect(t).
// 		Status(http.StatusCreated).
// 		End()
// }

// func TestCreateNote_ValidationFailed(t *testing.T) {
// 	var noteRequest *notes.Note = &notes.Note{}

// 	var token string = getJWTToken(t)

// 	apitest.New().
// 		Handler(newApp()).
// 		Post("/api/v1/notes").
// 		Header("Authorization", token).
// 		JSON(noteRequest).
// 		Expect(t).
// 		Status(http.StatusBadRequest).
// 		End()
// }

// func TestUpdateNote_Success(t *testing.T) {
// 	var note notes.Note = getNote()

// 	category := getCategory()

// 	var noteRequest *notes.Note = &notes.Note{
// 		Title:      "test",
// 		Content:    "test",
// 		CategoryID: category.ID,
// 	}

// 	noteID := strconv.Itoa(int(note.ID))

// 	var token string = getJWTToken(t)

// 	apitest.New().
// 		Observe(cleanup).
// 		Handler(newApp()).
// 		Put("/api/v1/notes/"+noteID).
// 		Header("Authorization", token).
// 		JSON(noteRequest).
// 		Expect(t).
// 		Status(http.StatusOK).
// 		End()
// }

// func TestUpdateNote_ValidationFailed(t *testing.T) {
// 	var note notes.Note = getNote()

// 	var noteRequest *notes.Note = &notes.Note{}

// 	noteID := strconv.Itoa(int(note.ID))

// 	var token string = getJWTToken(t)

// 	apitest.New().
// 		Handler(newApp()).
// 		Put("/api/v1/notes/"+noteID).
// 		Header("Authorization", token).
// 		JSON(noteRequest).
// 		Expect(t).
// 		Status(http.StatusBadRequest).
// 		End()
// }

// func TestDeleteNote_Success(t *testing.T) {
// 	var note notes.Note = getNote()

// 	var token string = getJWTToken(t)

// 	noteID := strconv.Itoa(int(note.ID))

// 	apitest.New().
// 		Observe(cleanup).
// 		Handler(newApp()).
// 		Delete("/api/v1/notes/"+noteID).
// 		Header("Authorization", token).
// 		Expect(t).
// 		Status(http.StatusOK).
// 		End()
// }

// func TestDeleteNote_Failed(t *testing.T) {
// 	var token string = getJWTToken(t)

// 	apitest.New().
// 		Handler(newApp()).
// 		Observe(cleanup).
// 		Delete("/api/v1/notes/-1").
// 		Header("Authorization", token).
// 		Expect(t).
// 		Status(http.StatusNotFound).
// 		End()
// }

// func TestRestoreNote_Success(t *testing.T) {
// 	var note notes.Note = getNote()

// 	var token string = getJWTToken(t)

// 	noteID := strconv.Itoa(int(note.ID))

// 	apitest.New().
// 		Observe(cleanup).
// 		Handler(newApp()).
// 		Post("/api/v1/notes/"+noteID).
// 		Header("Authorization", token).
// 		Expect(t).
// 		Status(http.StatusOK).
// 		End()
// }

// func TestRestoreNote_Failed(t *testing.T) {
// 	var token string = getJWTToken(t)

// 	apitest.New().
// 		Handler(newApp()).
// 		Observe(cleanup).
// 		Post("/api/v1/notes/-1").
// 		Header("Authorization", token).
// 		Expect(t).
// 		Status(http.StatusNotFound).
// 		End()
// }

// func TestForceDeleteNote_Success(t *testing.T) {
// 	var note notes.Note = getNote()

// 	var token string = getJWTToken(t)

// 	noteID := strconv.Itoa(int(note.ID))

// 	apitest.New().
// 		Observe(cleanup).
// 		Handler(newApp()).
// 		Delete("/api/v1/notes/force/"+noteID).
// 		Header("Authorization", token).
// 		Expect(t).
// 		Status(http.StatusOK).
// 		End()
// }

// func TestForceDeleteNote_Failed(t *testing.T) {
// 	var token string = getJWTToken(t)

// 	apitest.New().
// 		Handler(newApp()).
// 		Observe(cleanup).
// 		Post("/api/v1/notes/force/-1").
// 		Header("Authorization", token).
// 		Expect(t).
// 		Status(http.StatusNotFound).
// 		End()
// }

// func TestLogout_Success(t *testing.T) {
// 	var token string = getJWTToken(t)

// 	apitest.New().
// 		Handler(newApp()).
// 		Observe(cleanup).
// 		Post("/api/v1/users/logout").
// 		Header("Authorization", token).
// 		Expect(t).
// 		Status(http.StatusOK).
// 		End()
// }

// func TestLogout_Failed(t *testing.T) {
// 	apitest.New().
// 		Handler(newApp()).
// 		Observe(cleanup).
// 		Post("/api/v1/users/logout").
// 		Expect(t).
// 		Status(http.StatusBadRequest).
// 		End()
// }
