package product

import "go-specification/shared"

type IRepository interface {
	GetBySpecification(spec shared.Specification) ([]Entity, error)
}
