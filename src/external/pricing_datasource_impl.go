package external

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	domain "pinpag/codingdojogolang/src/domain/entities"
	"time"
)

type PricingDataSourceImpl struct {
	HttpClient http.Client
}

func (p *PricingDataSourceImpl) FetchPricingFromDate(date time.Time, code string, codein string) (domain.Price, error) {
	time.Sleep(3 * time.Second)

	req, err := http.NewRequest("GET", fmt.Sprintf("https://economia.awesomeapi.com.br/json/daily/%s-%s/", code, codein), nil)
	if err != nil {
		fmt.Print(err)
	}

	q := req.URL.Query()
	q.Add("start_date", date.Format("20060102"))
	q.Add("end_date", date.Format("20060102"))
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())

	response, responseErr := p.HttpClient.Do(req)

	if responseErr != nil {
		fmt.Println("ERROR:", err)
	}

	defer response.Body.Close()
	// Extrai conteúdo para salvar no banco de dados
	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var priceResponse []domain.Price
	json.Unmarshal([]byte(content), &priceResponse)

	return priceResponse[0], nil
}
