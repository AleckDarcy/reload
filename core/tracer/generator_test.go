package tracer

import (
	"encoding/json"
	"fmt"
	"math"
	"testing"
)

func TestName(t *testing.T) {
	//// home
	//var names = []string{
	//	"GetSupportedCurrenciesRequest",
	//	"ListProductsRequest",
	//	"GetCartRequest",
	//	"CurrencyConversionRequest",
	//	"CurrencyConversionRequest",
	//	"CurrencyConversionRequest",
	//	"CurrencyConversionRequest",
	//	"CurrencyConversionRequest",
	//	"CurrencyConversionRequest",
	//	"CurrencyConversionRequest",
	//	"CurrencyConversionRequest",
	//	"CurrencyConversionRequest",
	//	"AdRequest",
	//}

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

	// checkout
	var names = []string{
		"PlaceOrderRequest",
		"GetCartRequest",
		"GetProductRequest",
		"CurrencyConversionRequest",
		"GetProductRequest",
		"CurrencyConversionRequest",
		"GetQuoteRequest",
		"CurrencyConversionRequest",
		"ChargeRequest",
		"ShipOrderRequest",
		"EmptyCartRequest",
		"SendOrderConfirmationRequest",
		"ListRecommendationsRequest",
		"GetProductRequest",
		"GetProductRequest",
		"GetProductRequest",
		"GetProductRequest",
		"GetProductRequest",
	}

	resultChan := make(chan *Faults, 1)

	go func() {
		c := NewGenerator(names)
		jsonBytes, _ := json.Marshal(c.faults)
		fmt.Println("TFIs:", string(jsonBytes))
		for i := 1; i <= len(names); i++ {
			c.SetLen(i)
			c.Generate(resultChan)
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
