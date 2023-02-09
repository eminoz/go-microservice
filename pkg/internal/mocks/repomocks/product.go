// Code generated by mockery v2.18.0. DO NOT EDIT.

package repomocks

import (
	model "github.com/eminoz/go-api/pkg/model"
	mock "github.com/stretchr/testify/mock"
)

// ProductRepository is an autogenerated mock type for the ProductRepository type
type ProductRepository struct {
	mock.Mock
}

// CreateProduct provides a mock function with given fields: product
func (_m *ProductRepository) CreateProduct(product model.Product) model.ProductDto {
	ret := _m.Called(product)

	var r0 model.ProductDto
	if rf, ok := ret.Get(0).(func(model.Product) model.ProductDto); ok {
		r0 = rf(product)
	} else {
		r0 = ret.Get(0).(model.ProductDto)
	}

	return r0
}

// GetProductByID provides a mock function with given fields: id
func (_m *ProductRepository) GetProductByID(id string) {
	_m.Called(id)
}

type mockConstructorTestingTNewProductRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewProductRepository creates a new instance of ProductRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProductRepository(t mockConstructorTestingTNewProductRepository) *ProductRepository {
	mock := &ProductRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
