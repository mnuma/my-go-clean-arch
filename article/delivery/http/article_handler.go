package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/mnuma/my-go-clean-arch/article/usecase"
	"github.com/Sirupsen/logrus"
	"github.com/mnuma/my-go-clean-arch/article"
)

type HttpArticleHandler struct {
	AUsecase usecase.ArticleUsecase
}

func getStatusCode(err error) int {
	if err != nil {
		logrus.Error(err)
	}
	switch err {
	case article.INTERNAL_SERVER_ERROR:

		return http.StatusInternalServerError
	case article.NOT_FOUND_ERROR:
		return http.StatusNotFound
	case article.CONFLIT_ERROR:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func (a *HttpArticleHandler) GetByID(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	id := int64(idP)

	art, err := a.AUsecase.GetByID(id)

	if err != nil {
		return c.JSON(getStatusCode(err), err.Error())
	}
	return c.JSON(http.StatusOK, art)
}

func NewArticleHttpHandler(e *echo.Echo, us usecase.ArticleUsecase) {
	handler := &HttpArticleHandler{
		AUsecase: us,
	}
	e.GET("/article/:id", handler.GetByID)
}