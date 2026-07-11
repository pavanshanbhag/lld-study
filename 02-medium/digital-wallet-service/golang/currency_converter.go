package digitalwallet

import (
	"math/big"
	"sync"
)

type CurrencyConverter struct {
	exchangeRates map[Currency]*big.Float
	mu            sync.RWMutex
}

func NewCurrencyConverter() *CurrencyConverter {
	converter := &CurrencyConverter{
		exchangeRates: make(map[Currency]*big.Float),
	}
	converter.initializeRates()
	return converter
}

func (cc *CurrencyConverter) initializeRates() {
	cc.mu.Lock()
	defer cc.mu.Unlock()

	one := big.NewFloat(1)
	cc.exchangeRates[USD] = one
	cc.exchangeRates[EUR] = big.NewFloat(0.85)
	cc.exchangeRates[GBP] = big.NewFloat(0.72)
	cc.exchangeRates[JPY] = big.NewFloat(110.00)
}

func (cc *CurrencyConverter) Convert(amount *big.Float, sourceCurrency, targetCurrency Currency) *big.Float {
	cc.mu.RLock()
	defer cc.mu.RUnlock()

	sourceRate := cc.exchangeRates[sourceCurrency]
	targetRate := cc.exchangeRates[targetCurrency]

	result := new(big.Float).Mul(amount, sourceRate)
	return new(big.Float).Quo(result, targetRate)
}
