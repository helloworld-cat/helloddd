package product

type Entity struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	PriceCents uint   `json:"price_cents"`
}
