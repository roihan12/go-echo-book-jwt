package controller

import (
	"bytes"
	"echo-book/config"
	"echo-book/dto"
	"echo-book/entity"
	"echo-book/repository"
	"echo-book/service"
	"encoding/json"

	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	db                 *gorm.DB                  = config.SetupDatabaseConnectionTest()
	bookRepository     repository.BookRepository = repository.NewBookRepository(db)
	jwtService         service.JWTService        = service.NewJWTService()
	bookService        service.BookService       = service.NewBookService(bookRepository)
	bookTestController BookController            = NewBookController(bookService, jwtService)
)

func SeedUser() entity.User {
	password, _ := bcrypt.GenerateFromPassword([]byte("123123"), bcrypt.DefaultCost)

	var user entity.User = entity.User{
		Name:     "Test",
		Email:    "testing@mail.com",
		Password: string(password),
	}

	if err := db.Create(&user).Error; err != nil {
		panic(err)
	}

	var createdUser entity.User

	db.Last(&createdUser)

	createdUser.Password = "123123"

	return createdUser
}

func SeedBook() entity.Book {

	var note entity.Book = entity.Book{
		Title:       "test",
		Description: "test",
	}

	if err := db.Create(&note).Error; err != nil {
		panic(err)
	}

	var createdNote entity.Book

	db.Last(&createdNote)

	return createdNote
}

func InitEcho() *echo.Echo {
	config.SetupDatabaseConnectionTest()

	e := echo.New()

	return e
}

func TestGetAllBooks_Success(t *testing.T) {
	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         bool
		expectedBodyStartsWith string
	}{{
		name:                   "success",
		path:                   "/api/books",
		expectedStatus:         true,
		expectedBodyStartsWith: "{\"status\":",
	},
	}

	e := InitEcho()

	req := httptest.NewRequest(http.MethodGet, "/api/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, bookTestController.All(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}
}

func TestCreateBook_Success(t *testing.T) {
	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         bool
		expectedBodyStartsWith string
	}{{
		name:                   "success",
		path:                   "/api/books",
		expectedStatus:         true,
		expectedBodyStartsWith: "{\"status\":",
	},
	}

	e := InitEcho()

	// category := database.SeedCategory()

	bookInput := dto.BookCreateDTO{
		Title:       "test",
		Description: "test",
	}

	jsonBody, _ := json.Marshal(&bookInput)
	bodyReader := bytes.NewReader(jsonBody)

	req := httptest.NewRequest(http.MethodPost, "/api/books", bodyReader)
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/json")

	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, bookTestController.Insert(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}

}

// func TestGetNoteByID_Success(t *testing.T) {
// 	var testCases = []struct {
// 		name                   string
// 		path                   string
// 		expectedStatus         int
// 		expectedBodyStartsWith string
// 	}{{
// 		name:                   "success",
// 		path:                   "/api/v1/notes",
// 		expectedStatus:         http.StatusOK,
// 		expectedBodyStartsWith: "{\"status\":",
// 	},
// 	}

// 	e := InitEcho()

// 	note := database.SeedNote()
// 	noteID := strconv.Itoa(int(note.ID))

// 	req := httptest.NewRequest(http.MethodGet, "/api/v1/notes", nil)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	for _, testCase := range testCases {
// 		c.SetPath(testCase.path)
// 		c.SetParamNames("id")
// 		c.SetParamValues(noteID)

// 		if assert.NoError(t, GetByID(c)) {
// 			assert.Equal(t, http.StatusOK, rec.Code)
// 			body := rec.Body.String()

// 			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
// 		}
// 	}

// }

// func TestUpdateNote_Success(t *testing.T) {
// 	var testCases = []struct {
// 		name                   string
// 		path                   string
// 		expectedStatus         int
// 		expectedBodyStartsWith string
// 	}{{
// 		name:                   "success",
// 		path:                   "/api/v1/notes",
// 		expectedStatus:         http.StatusOK,
// 		expectedBodyStartsWith: "{\"status\":",
// 	},
// 	}

// 	e := InitEcho()

// 	note := database.SeedNote()

// 	noteInput := model.NoteInput{
// 		Title:      "test",
// 		Content:    "test",
// 		CategoryID: note.CategoryID,
// 	}

// 	jsonBody, _ := json.Marshal(&noteInput)
// 	bodyReader := bytes.NewReader(jsonBody)

// 	noteID := strconv.Itoa(int(note.ID))

// 	req := httptest.NewRequest(http.MethodPut, "/api/v1/notes", bodyReader)
// 	rec := httptest.NewRecorder()

// 	req.Header.Add("Content-Type", "application/json")

// 	c := e.NewContext(req, rec)

// 	for _, testCase := range testCases {
// 		c.SetPath(testCase.path)
// 		c.SetParamNames("id")
// 		c.SetParamValues(noteID)

// 		if assert.NoError(t, Update(c)) {
// 			assert.Equal(t, http.StatusOK, rec.Code)
// 			body := rec.Body.String()

// 			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
// 		}
// 	}

// }

// func TestDeleteNote_Success(t *testing.T) {
// 	var testCases = []struct {
// 		name                   string
// 		path                   string
// 		expectedStatus         int
// 		expectedBodyStartsWith string
// 	}{{
// 		name:                   "success",
// 		path:                   "/api/v1/notes",
// 		expectedStatus:         http.StatusOK,
// 		expectedBodyStartsWith: "{\"status\":",
// 	},
// 	}

// 	e := InitEcho()

// 	note := database.SeedNote()
// 	noteID := strconv.Itoa(int(note.ID))

// 	req := httptest.NewRequest(http.MethodDelete, "/api/v1/notes", nil)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	for _, testCase := range testCases {
// 		c.SetPath(testCase.path)
// 		c.SetParamNames("id")
// 		c.SetParamValues(noteID)

// 		if assert.NoError(t, Delete(c)) {
// 			assert.Equal(t, http.StatusOK, rec.Code)
// 			body := rec.Body.String()

// 			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
// 		}
// 	}

// }
