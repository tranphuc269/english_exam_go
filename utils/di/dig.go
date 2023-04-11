package di

import (
	v1 "english_exam_go/application/api/v1"
	"english_exam_go/application/routers"
	"english_exam_go/domain/services"
	"english_exam_go/infrastructure/data/repositories"
	"english_exam_go/infrastructure/data/repositories/persistence"
	"english_exam_go/utils/app_logger"
	"go.uber.org/dig"
)

var errors []error
var errCount = 0
var errMsg = ""

func RegisterDIFunction() *dig.Container {
	c := dig.New()

	registerTransaction(c)
	registerRepository(c)
	registerService(c)
	registerController(c)
	registerRouter(c)

	for _, e := range errors {
		if e != nil {
			errMsg = errMsg + e.Error() + ", "
			errCount = errCount + 1
		}
	}
	if errCount > 0 {
		app_logger.Logger.Panic("Can not register DI function. error:" + errMsg)
	}

	return c
}

func registerTransaction(c *dig.Container) {
	errors = append(errors, c.Provide(repositories.TransactionImpl))
}

func registerRepository(c *dig.Container) {
	errors = append(errors, c.Provide(persistence.CreateBookRepository))
	errors = append(errors, c.Provide(persistence.CreateAuthRepository))
	errors = append(errors, c.Provide(persistence.CreateExamRepository))
	errors = append(errors, c.Provide(persistence.CreateExamInviteRepository))
}

func registerService(c *dig.Container) {
	errors = append(errors, c.Provide(services.CreateBookService))
	errors = append(errors, c.Provide(services.CreateAuthService))
	errors = append(errors, c.Provide(services.CreateExamService))
	errors = append(errors, c.Provide(services.CreateExamInviteService))
}

func registerController(c *dig.Container) {
	errors = append(errors, c.Provide(v1.CreateBookController))
	errors = append(errors, c.Provide(v1.CreateAuthController))
	errors = append(errors, c.Provide(v1.CreateExamController))
	errors = append(errors, c.Provide(v1.CreateExamInviteController))
}

func registerRouter(c *dig.Container) {

	errors = append(errors, c.Provide(routers.CreateBookRouter, dig.Name("book")))
	errors = append(errors, c.Provide(routers.CreateAuthRouter, dig.Name("auth")))
	errors = append(errors, c.Provide(routers.CreateExamRouter, dig.Name("exam")))
	errors = append(errors, c.Provide(routers.CreateExamInviteRouter, dig.Name("exam_invite")))

	// RouterService
	errors = append(errors, c.Provide(routers.NewService))
}
