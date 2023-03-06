package repositories

import (
	"fmt"
	"pinpag/codingdojogolang/data/datasources"
	domain "pinpag/codingdojogolang/domain/entities"
	"time"
)

type NewAndPricingRepositories struct {
	NewsDataSource    datasources.NewsDataSourceInterface
	PricingDataSource datasources.PricingDataSourceInterface
}

func (n *NewAndPricingRepositories) GetAllNewsAndPricingFromDate(date time.Time, code string, codein string) (domain.NewsAndPrice, error) {
	newsResponse, newsError := n.NewsDataSource.FetchLastestNewsFromDate(date, code, codein)

	if newsError != nil {
		fmt.Println("um erro aconteceu ao buscar notificas sobre", code, codein)
	}

	if len(newsResponse) == 0 {
		fmt.Println("Não temos noticias para", code, codein)
	}

	pricingResponse, pricingError := n.PricingDataSource.FetchPricingFromDate(date, code, codein)

	if pricingError != nil {
		fmt.Println("um erro aconteceu ao buscar preços para", code, codein)
	}

	fmt.Print(pricingResponse)
	fmt.Print(newsResponse)

	newsAndPrice := domain.NewsAndPrice{Price: pricingResponse, News: newsResponse}

	return newsAndPrice, nil
}
