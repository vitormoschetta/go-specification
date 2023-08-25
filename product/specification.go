package product

import "time"

type ProductsCreatedTodaySpecification struct {
	today string
}

func NewProductsCreatedTodaySpecification() ProductsCreatedTodaySpecification {
	return ProductsCreatedTodaySpecification{today: time.Now().UTC().Format("2006-01-02")}
}

func (s ProductsCreatedTodaySpecification) IsSatisfiedBy(candidate any) bool {
	product := candidate.(Entity)
	return product.CreatedAt == s.today
}
