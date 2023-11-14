package product

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProductBySpecification(t *testing.T) {
	// When | Arrange
	repository := NewInMemoryRepository()
	repository.Seed()
	handler := NewProductHandler(repository)

	spec := NewGetProductsCreatedTodaySpecification()

	// Given | Act
	items, err := handler.GetBySpecification(spec)

	// Then | Assert
	assert.Nil(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, 1, len(items.([]Entity)))
}

func TestGetProductBySpecificationWithAnd(t *testing.T) {
	// When | Arrange
	repository := NewInMemoryRepository()
	repository.Seed()
	handler := NewProductHandler(repository)

	spec := GetProductsWithPriceBetweenSpecification{
		min: 100,
		max: 200,
	}

	// Given | Act
	items, err := handler.GetBySpecification(spec)

	// Then | Assert
	assert.Nil(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, 2, len(items.([]Entity)))
}
