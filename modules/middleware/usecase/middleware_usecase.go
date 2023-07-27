package usecase

import "github.com/pheethy/go-fiber-tutorial/modules/middleware/repository"

type ImiddlewareUsecase interface {

}

type middlewareUsecase struct {
	middleRepo repository.ImiddlewareRepository
}

func NewMiddlewareUsecase(middleRepo repository.ImiddlewareRepository) ImiddlewareUsecase {
	return middlewareUsecase{middleRepo: middleRepo}
}