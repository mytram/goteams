package rest

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var ErrBadQuest = errors.New("Bad request")

type jsonError struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Details []string `json:"details,omitempty"`
}

func JSON(ctx *gin.Context, data interface{}, err error, errs ...error) {
	if err == nil {
		if data != nil {
			ctx.JSON(http.StatusOK, gin.H{"data": data})
		} else {
			ctx.Status(http.StatusOK)
		}

		return
	}

	var details []string

	if errs != nil {
		details = make([]string, len(errs))

		for i := range errs {
			details[i] = errs[i].Error()
		}
	}

	if err == ErrBadQuest {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": jsonError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Details: details,
		}})

		return
	}

	if err == gorm.ErrRecordNotFound {
		ctx.Status(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": jsonError{
		Code:    http.StatusUnprocessableEntity,
		Message: err.Error(),
		Details: details,
	}})
}
