package entity

import (
	"errors"
)

type OrderRepositoryInterface interface {
	Save(order *Order) error
}

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

// => * Ponteiro pega o valor direto da memoria
func NewOrder(id string, price float64, tax float64) (*Order, error) {
	// & => Referencia de Order pois estamos mexendo com ponteiro
	order := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}
	error := order.IsValid()
	if error != nil {
		return nil, error
	}

	return order, nil
}

func (o *Order) CalculateFinalPrice() error {
	err := o.IsValid()
	if err != nil {
		return err
	}
	o.FinalPrice = o.Price + o.Tax
	return nil
}

func (o Order) IsValid() error {
	if o.ID == "" {
		return errors.New("invalid id")
	}
	if o.Price == 0 {
		return errors.New("invalid price")

	}
	if o.Tax == 0 {
		return errors.New("invalid tax")

	}
	return nil
}
