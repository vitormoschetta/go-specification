package product

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProductBySpecification(t *testing.T) {
	// When | Arrange
	spec := NewProductsCreatedTodaySpecification()

	repository := NewInMemoryRepository()
	repository.Seed()
	useCase := NewProductUseCase(repository)

	// Given | Act
	output := useCase.GetBySpecification(spec)

	// Then | Assert
	assert.NotNil(t, output)
	assert.Equal(t, 0, len(output.GetErrors()))
	assert.Equal(t, 1, len(output.GetData().([]Entity)))
}
