package collections_market_pricings

type MarketPricing struct {
	Key           string  `json:"key,omitempty"`
	BasicAmount   int64   `json:"basicAmount,omitempty"`
	PremiumAmount float64 `json:"premiumAmount,omitempty"`
}
