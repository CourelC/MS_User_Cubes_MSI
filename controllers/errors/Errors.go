package errors

import (
	"net/http"

	models "MS_User_Cubes_MSI/errors"

	"github.com/gin-gonic/gin"
)

type MessagesResponse struct {
	Message string `json:"message"`
}

func Handler(c *gin.Context) {

	c.Next()

	errs := c.Errors

	if len(errs) > 0 {
		err, ok := errs[0].Err.(*models.AppError)
		if ok {
			resp := MessagesResponse{Message: err.Error()}
			switch err.Type {
			case models.NotFound:
				c.JSON(http.StatusNotFound, resp)
				return
			case models.ValidationError:
				c.JSON(http.StatusBadRequest, resp)
				return
			case models.ResourceAlreadyExists:
				c.JSON(http.StatusConflict, resp)
				return
			case models.NotAuthenticated:
				c.JSON(http.StatusUnauthorized, resp)
				return
			case models.NotAuthorized:
				c.JSON(http.StatusForbidden, resp)
				return
			case models.RepositoryError:
				c.JSON(http.StatusInternalServerError, MessagesResponse{Message: "We are working to improve the flow of this request."})
				return
			default:
				c.JSON(http.StatusInternalServerError, MessagesResponse{Message: "We are working to improve the flow of this request."})
				return
			}
		}

		return
	}
}
