package infrastructure

import (
	"context"
	"time"

	coreAuthDomain "go.aoe.com/flamingo/core/auth/domain"
	"go.aoe.com/flamingo/core/order/domain"
)

type (
	// CustomerOrders is the CustomerOrders api service
	FakeCustomerOrders struct{}
)

// Get returns a CustomerOrders struct
func (co *FakeCustomerOrders) Get(ctx context.Context, authentication coreAuthDomain.Auth) ([]*domain.Order, error) {
	return []*domain.Order{
		{
			ID: "100",

			CreationTime: time.Now(),
			UpdateTime:   time.Now(),
			Status:       "Teststatus",
			OrderItems:   make([]domain.OrderItem, 0),
			Total:        123.45,
			CurrencyCode: "EUR",
		}, {
			ID: "100",

			CreationTime: time.Now(),
			UpdateTime:   time.Now(),
			Status:       "Teststatus",
			OrderItems:   make([]domain.OrderItem, 0),
			Total:        123.45,
			CurrencyCode: "EUR",
		}, {
			ID: "100",

			CreationTime: time.Now(),
			UpdateTime:   time.Now(),
			Status:       "Teststatus",
			OrderItems:   make([]domain.OrderItem, 0),
			Total:        123.45,
			CurrencyCode: "EUR",
		}, {
			ID: "100",

			CreationTime: time.Now(),
			UpdateTime:   time.Now(),
			Status:       "Teststatus",
			OrderItems:   make([]domain.OrderItem, 0),
			Total:        123.45,
			CurrencyCode: "EUR",
		},
	}, nil
}