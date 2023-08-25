package product

import (
	"go-specification/shared"
	"log"
)

type ProductUseCase struct {
	Repository IRepository
}

func NewProductUseCase(repository IRepository) *ProductUseCase {
	return &ProductUseCase{
		Repository: repository,
	}
}

func (c *ProductUseCase) GetBySpecification(spec shared.Specification) (output Output) {
	entities, err := c.Repository.GetBySpecification(spec)
	if err != nil {
		log.Println(err)
		output.SetError(DomainCodeInternalError, err)
		return output
	}

	output.SetOk(entities)
	return output
}
