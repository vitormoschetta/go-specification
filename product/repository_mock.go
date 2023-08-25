package product

import (
	"go-specification/shared"
	"time"
)

type InMemoryRepository struct {
	storage map[string]Entity
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		storage: make(map[string]Entity),
	}
}

func (r *InMemoryRepository) GetBySpecification(spec shared.Specification) ([]Entity, error) {
	var result []Entity
	for _, entity := range r.storage {
		if spec.IsSatisfiedBy(entity) {
			result = append(result, entity)
		}
	}
	return result, nil
}

func (r *InMemoryRepository) Seed() {
	r.storage = map[string]Entity{
		"1": {
			ID:        "1",
			Name:      "Product 1",
			Price:     100,
			CreatedAt: time.Now().UTC().AddDate(0, 0, -1).Format("2006-01-02"),
		},
		"2": {
			ID:        "2",
			Name:      "Product 2",
			Price:     200,
			CreatedAt: time.Now().UTC().Format("2006-01-02"),
		},
	}
}
