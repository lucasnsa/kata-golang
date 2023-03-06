package repositories

import (
	"errors"
	"fmt"
	"pinpag/codingdojogolang/src/data/datasources"
	domain "pinpag/codingdojogolang/src/domain/entities"
	"time"
)

type ParallelNewAndPricingRepositories struct {
	NewsDataSource    datasources.NewsDataSourceInterface
	PricingDataSource datasources.PricingDataSourceInterface
}

func (n *ParallelNewAndPricingRepositories) GetAllNewsAndPricingFromDate(date time.Time, code string, codein string) (domain.NewsAndPrice, error) {
	chanErr := make(chan error, 1)
	chanNews := make(chan []domain.News)
	chanPrice := make(chan domain.Price)

	go n.coGetNews(date, code, codein, chanNews, chanErr)
	go n.coGetPricing(date, code, codein, chanPrice, chanErr)

	if err := <-chanErr; err != nil {
		return domain.NewsAndPrice{}, err
	}

	newsResponse := <-chanNews
	priceResponse := <-chanPrice
	newsAndPrice := domain.NewsAndPrice{Price: priceResponse, News: newsResponse}

	return newsAndPrice, nil
}

func (n *ParallelNewAndPricingRepositories) coGetNews(date time.Time, code string, codein string, news chan []domain.News, err chan error) {
	newsResponse, newsError := n.NewsDataSource.FetchLastestNewsFromDate(date, code, codein)

	if newsError != nil {
		fmt.Println("um erro aconteceu ao buscar notificas sobre", code, codein)
		err <- newsError
	}

	if len(newsResponse) == 0 {
		fmt.Println("Não temos noticias para", code, codein)
		err <- errors.New("não temos noticias")
	}

	err <- nil
	news <- newsResponse
	close(news)
}
func (n *ParallelNewAndPricingRepositories) coGetPricing(date time.Time, code string, codein string, pricing chan domain.Price, err chan error) {
	pricingResponse, pricingError := n.PricingDataSource.FetchPricingFromDate(date, code, codein)

	if pricingError != nil {
		fmt.Println("um erro aconteceu ao buscar preços para", code, codein)
		err <- errors.New("não podemos buscar os erros")
	}

	err <- nil
	pricing <- pricingResponse
	close(pricing)
}
