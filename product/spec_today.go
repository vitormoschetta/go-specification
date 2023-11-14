package product

import "time"

type GetProductsCreatedTodaySpecification struct {
	today string
}

func NewGetProductsCreatedTodaySpecification() GetProductsCreatedTodaySpecification {
	return GetProductsCreatedTodaySpecification{today: time.Now().UTC().Format("2006-01-02")}
}

func (s GetProductsCreatedTodaySpecification) IsSatisfiedBy(candidate any) bool {
	product := candidate.(Entity)
	return product.CreatedAt == s.today
}
