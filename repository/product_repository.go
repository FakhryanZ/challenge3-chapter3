package repository

import "golang-learning-path/go-testify/models"

type ProductRepository interface {
	FindById(id string) *models.Product
	GetProducts() *[]models.Product
}
