package service

import (
	"errors"
	"golang-learning-path/go-testify/models"
	"golang-learning-path/go-testify/repository"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (service ProductService) GetOneProduct(id string) (*models.Product, error) {
	product := service.Repository.FindById(id)
	if product == nil {
		return nil, errors.New("product not found")
	}

	return product, nil
}

func (service ProductService) GetAllProduct() (*[]models.Product, error) {
	products := service.Repository.GetProducts()
	if products == nil {
		return nil, errors.New("product is empty")
	}

	return products, nil
}
