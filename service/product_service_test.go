package service

import (
	"golang-learning-path/go-testify/models"
	"golang-learning-path/go-testify/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{Repository: productRepository}

func TestProductServiceGetOneProductNotFound(t *testing.T) {
	productRepository.Mock.On("FindById", "1").Return(nil)

	product, err := productService.GetOneProduct("1")

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), "error response has to be 'product not found'")
}

func TestProductServiceGetOneProductFound(t *testing.T) {
	product := models.Product{
		Title:       "Test",
		Description: "TestTestTest",
	}

	product.ID = 2

	productRepository.Mock.On("FindById", "2").Return(product)

	result, err := productService.GetOneProduct("2")

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, product.Title, result.Title, "result has to be 'Test'")
	assert.Equal(t, product.Description, result.Description, "result has to be 'TestTestTest'")
	assert.Equal(t, &product, result, "res has to be a product data with id '2'")
}

func TestProductServiceGetProductsNotFound(t *testing.T) {
	productRepository.Mock.On("GetProducts").Return(nil)

	products, err := productService.GetAllProduct()

	assert.Nil(t, products)
	assert.NotNil(t, err)
	assert.Equal(t, "product is empty", err.Error(), "error response has to be 'product is empty'")
}

func TestProductServiceGetProductsFound(t *testing.T) {
	products := []models.Product{
		{
			Title:       "Test 1",
			Description: "Test 1",
		},
		{
			Title:       "Test 2",
			Description: "Test 2",
		},
	}

	productRepository.Mock.On("GetProducts").Return(products)

	result, err := productService.GetAllProduct()

	assert.Nil(t, err)
	assert.NotNil(t, result)
}
