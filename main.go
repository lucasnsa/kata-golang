package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"pinpag/codingdojogolang/src/data/datasources"
	"pinpag/codingdojogolang/src/data/repositories"
	domain "pinpag/codingdojogolang/src/domain/repository"
	"pinpag/codingdojogolang/src/external"
	"pinpag/codingdojogolang/src/usecases"
	"time"

	"github.com/gin-gonic/gin"
)

// variaveis
var (
	mustBeHttpClient  *http.Client
	newsDataSource    datasources.NewsDataSourceInterface
	pricingDataSource datasources.PricingDataSourceInterface
	repo              domain.NewAndPricingRepositoriesInterface
	coRepo            domain.NewAndPricingRepositoriesInterface
)

func main() {
	fmt.Println("Bem vindo ao Golang!!")

	mustBeHttpClient = &http.Client{}

	newsDataSource = &external.NewsDataSourceImpl{HttpClient: *mustBeHttpClient}
	pricingDataSource = &external.PricingDataSourceImpl{HttpClient: *mustBeHttpClient}

	repo = &repositories.NewAndPricingRepositories{NewsDataSource: newsDataSource, PricingDataSource: pricingDataSource}
	coRepo = &repositories.ParallelNewAndPricingRepositories{NewsDataSource: newsDataSource, PricingDataSource: pricingDataSource}

	router := gin.Default()

	apiV1 := router.Group("v1")

	apiV1.GET("/randomInt", getRandomInt)

	apiV1.GET("/mysecret", getNewsAndPrice)

	apiV2 := router.Group("v2")

	apiV2.GET("/mysecret", getNewsAndPriceParallel)

	router.Run("localhost:8080")
}

func getRandomInt(c *gin.Context) {
	c.IndentedJSON(http.StatusBadRequest, rand.Int())
}

func getNewsAndPrice(c *gin.Context) {
	date := c.DefaultQuery("date", time.Now().Format("2006-01-02"))
	code := c.Query("code")
	codein := c.Query("codein")
	useCase := usecases.GetAllNewsAndPricesUseCase{Repository: repo}

	resp, err := useCase.Execute(date, code, codein)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func getNewsAndPriceParallel(c *gin.Context) {
	date := c.DefaultQuery("date", time.Now().Format("2006-01-02"))
	code := c.Query("code")
	codein := c.Query("codein")
	useCase := usecases.GetAllNewsAndPricesUseCase{Repository: coRepo}

	resp, err := useCase.Execute(date, code, codein)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
