package product

type GetProductsWithPriceBetweenSpecification struct {
	min, max float64
}

func (s GetProductsWithPriceBetweenSpecification) IsSatisfiedBy(candidate any) bool {
	product := candidate.(Entity)
	return product.Price >= s.min && product.Price <= s.max
}
