package exception

import (
	"english_exam_go/utils/app_logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
)

func Handle(err error, c *gin.Context) {
	ae, ok := errors.Cause(err).(applicationError)
	if ok {
		handleAppError(ae, c, err)
		return
	}

	se, ok := errors.Cause(err).(systemError)
	if ok {
		handleSysError(se, c, err)
		return
	}

}

func handleAppError(ae applicationError, c *gin.Context, err error) {
	app_logger.Logger.Info(fmt.Sprintf("application error occured: %s", err.Error()), zap.String("errorCode", ae.Code()))

	c.JSON(ae.HTTPStatus(),
		gin.H{"code": ae.Code(), "message": ae.Message(), "status": ae.HTTPStatus()})
	return
}

const systemErrMsg string = "Đã xảy ra lỗi hệ thống không mong muốn. Xin vui lòng liên hệ với quản trị hệ thống của bạn\n"

func handleSysError(se systemError, c *gin.Context, err error) {

	if se.IsInternal() {
		app_logger.Logger.Warn(fmt.Sprintf("system error occurred: %s", err.Error()), zap.String("errorCode", se.Code()))
	} else {
		app_logger.Logger.Warn(fmt.Sprintf("unexpected error occurred: %s", err.Error()), zap.String("errorCode", se.Code()))
	}

	c.JSON(http.StatusInternalServerError,
		gin.H{"code": se.Code(), "message": systemErrMsg, "status": http.StatusInternalServerError})
	return
}
