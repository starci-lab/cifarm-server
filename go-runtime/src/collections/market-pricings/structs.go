package collections_market_pricings

type MarketPricing struct {
	Key           string `json:"key"`
	BasicAmount   int64  `json:"basicAmount"`
	PremiumAmount int64  `json:"premiumAmount"`
}
