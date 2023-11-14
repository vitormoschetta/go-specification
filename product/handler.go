package product

import (
	"go-specification/shared"
	"log"
)

type ProductHandler struct {
	Repository IRepository
}

func NewProductHandler(repository IRepository) *ProductHandler {
	return &ProductHandler{
		Repository: repository,
	}
}

func (c *ProductHandler) GetBySpecification(spec shared.Specification) (any, error) {
	entities, err := c.Repository.GetBySpecification(spec)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return entities, nil
}
