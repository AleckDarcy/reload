package tracer

import (
	"encoding/json"
	"fmt"
	"math"
	"testing"
)

func TestName(t *testing.T) {
	getSupportedCurrenciesRequest := &NameMeta{
		Service: "currencyservice",
		Name:    "GetSupportedCurrenciesRequest",
	}
	currencyConversionRequest := &NameMeta{
		Service: "currencyservice",
		Name:    "CurrencyConversionRequest",
	}

	// home
	var names = []*NameMeta{
		getSupportedCurrenciesRequest,
		{
			Service: "productcatalogservice",
			Name:    "ListProductsRequest",
		},
		{
			Service: "cartservice",
			Name:    "GetCartRequest",
		},
		currencyConversionRequest,
		currencyConversionRequest,
		currencyConversionRequest,
		currencyConversionRequest,
		currencyConversionRequest,
		currencyConversionRequest,
		currencyConversionRequest,
		currencyConversionRequest,
		currencyConversionRequest,
		{
			Service: "adservice",
			Name:    "AdRequest",
		},
	}

	//// product
	//var names = []string{
	//	"GetProductRequest",
	//	"GetSupportedCurrenciesRequest",
	//	"GetCartRequest",
	//	"CurrencyConversionRequest",
	//	"ListRecommendationsRequest",
	//	"GetProductRequest",
	//	"GetProductRequest",
	//	"GetProductRequest",
	//	"GetProductRequest",
	//	"GetProductRequest",
	//	"AdRequest",
	//}

	//// checkout
	//var names = []string{
	//	"PlaceOrderRequest",
	//	"GetCartRequest",
	//	"GetProductRequest",
	//	"CurrencyConversionRequest",
	//	"GetProductRequest",
	//	"CurrencyConversionRequest",
	//	"GetQuoteRequest",
	//	"CurrencyConversionRequest",
	//	"ChargeRequest",
	//	"ShipOrderRequest",
	//	"EmptyCartRequest",
	//	"SendOrderConfirmationRequest",
	//	"ListRecommendationsRequest",
	//	"GetProductRequest",
	//	"GetProductRequest",
	//	"GetProductRequest",
	//	"GetProductRequest",
	//	"GetProductRequest",
	//}

	resultChan := make(chan *Faults, 1)

	go func() {
		c := NewGenerator(names)
		jsonBytes, _ := json.Marshal(c.faults)
		fmt.Println("TFIs:", string(jsonBytes))
		for i := 1; i <= len(names); i++ {
			c.SetLen(i)
			c.GenerateCombinartorial(resultChan)
		}
	}()

	signal := make(chan struct{}, 1)
	go func(signal chan struct{}) {
		count := 0
		expect := int(math.Pow(2, float64(len(names))) - 1)
		for {
			select {
			case result := <-resultChan:
				jsonBytes, _ := json.Marshal(result)
				fmt.Println(string(jsonBytes))
				_ = jsonBytes
				count++
				if count == expect {
					signal <- struct{}{}

					return
				}
			}
		}
	}(signal)
	<-signal
}
