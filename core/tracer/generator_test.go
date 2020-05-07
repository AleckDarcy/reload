package tracer

import (
	"encoding/json"
	"fmt"
	"math"
	"testing"
)

func TestName(t *testing.T) {
	var names = []string{
		"GetSupportedCurrenciesRequest",
		"ListProductsRequest",
		"CurrencyConversionRequest",
		"CurrencyConversionRequest",
		"CurrencyConversionRequest",
		"CurrencyConversionRequest",
		"CurrencyConversionRequest",
		"CurrencyConversionRequest",
		"CurrencyConversionRequest",
		"CurrencyConversionRequest",
		"CurrencyConversionRequest",
		"AdRequest",
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
