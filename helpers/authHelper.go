package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func MatchUserTypeToUserid(ctx *gin.Context, userID string) (err error) {
	userType := ctx.GetString("user_type")
	uid := ctx.GetString("uid")
	err = nil

	if userType == "USER" && uid != userID {

		err = errors.New("Unauthorized to access this resource")
		return err
	}

	err = CheckUserType(ctx, userType)

	return err
}
