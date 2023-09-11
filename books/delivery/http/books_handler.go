package http

import (
	"net/http"
	"strconv"

	"github.com/LayssonENS/go-simple-gin/domain"
	"github.com/gin-gonic/gin"
)

type BooksHandler struct {
	bUseCase domain.BooksUseCase
}

func NewBooksHandler(routerGroup *gin.Engine, us domain.BooksUseCase) {
	handler := &BooksHandler{
		bUseCase: us,
	}

	routerGroup.GET("/v1/books/:booksId", handler.GetByID)
	routerGroup.GET("/v1/books/all", handler.GetAllBooks)
	routerGroup.POST("/v1/books", handler.CreateBooks)
}

// GetByID godoc
// @Summary Get Books by ID
// @Description get Books by ID
// @Tags Books
// @Accept  json
// @Produce  json
// @Param booksId path int true "Books ID"
// @Success 200 {object} domain.Books
// @Failure 400 {object} domain.ErrorResponse
// @Router /v1/books/{id} [get]
func (h *BooksHandler) GetByID(c *gin.Context) {
	idParam, err := strconv.Atoi(c.Param("booksId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{ErrorMessage: err.Error()})
		return
	}

	booksId := int64(idParam)

	response, err := h.bUseCase.GetByID(booksId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetAllBooks godoc
// @Summary Get all Bookss
// @Description get all Bookss
// @Tags Books
// @Accept  json
// @Produce  json
// @Success 200 {array} domain.Books
// @Failure 400 {object} domain.ErrorResponse
// @Router /v1/books/all [get]
func (h *BooksHandler) GetAllBooks(c *gin.Context) {
	response, err := h.bUseCase.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{ErrorMessage: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// CreateBooks godoc
// @Summary Create a new Books
// @Description create new Books
// @Tags Books
// @Accept  json
// @Produce  json
// @Param books body domain.BooksRequest true "Create Books"
// @Success 201 {object} string
// @Failure 400 {object} domain.ErrorResponse
// @Failure 422 {object} domain.ErrorResponse
// @Router /v1/books [post]
func (h *BooksHandler) CreateBooks(c *gin.Context) {
	var books domain.BooksRequest
	if err := c.ShouldBindJSON(&books); err != nil {
		c.JSON(http.StatusUnprocessableEntity, domain.ErrorResponse{ErrorMessage: err.Error()})
		return
	}

	err := h.bUseCase.CreateBooks(books)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{ErrorMessage: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"msg": "Created"})
}
