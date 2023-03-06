package usecases

import (
	"errors"
	domain "pinpag/codingdojogolang/domain/entities"
	repositories "pinpag/codingdojogolang/domain/repository"
	"time"
)

type GetAllNewsAndPricesUseCase struct {
	Repository repositories.NewAndPricingRepositoriesInterface
}

func (g *GetAllNewsAndPricesUseCase) Execute(date string, code string, codeIn string) (*domain.NewsAndPrice, error) {
	dateConverted, dateErr := time.Parse("2006-01-02", date)

	if dateErr != nil {
		return nil, errors.New("data inválida")
	}

	newsAndPrice, err := g.Repository.GetAllNewsAndPricingFromDate(dateConverted, code, codeIn)

	if err != nil {
		return nil, err
	}

	return &newsAndPrice, nil

}
