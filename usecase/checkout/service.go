package checkout

import (
	"errors"
	"fmt"
	"github.com/vitoraalmeida/silkroad/usecase/sale"
	"github.com/vitoraalmeida/silkroad/usecase/saleitem"
)

type Service struct {
	saleService     sale.UseCase
	saleItemService saleitem.UseCase
}

type CartItem struct {
	ProductID uint
	Quantity  uint
	Subtotal  float64
}

type Cart []CartItem

var ErrInvalidCustomer = errors.New("Customer doesn't exist")
var ErrCartEmpty = errors.New("Cart cannot be empty")

func NewService(s sale.UseCase, si saleitem.UseCase) *Service {
	return &Service{
		saleService:     s,
		saleItemService: si,
	}
}

func (s *Service) Checkout(c *Cart, customerID uint) error {
	if len(*c) == 0 {
		return ErrCartEmpty
	}
	if customerID <= 0 {
		return ErrInvalidCustomer
	}

	totalAmount := 0.00
	for _, ci := range *c {
		totalAmount += ci.Subtotal
	}

	saleID, _ := s.saleService.CreateSale(customerID, totalAmount)

	fmt.Println("\nSale ID: ", saleID)

	for _, ci := range *c {
		id, _ := s.saleItemService.CreateSaleItem(saleID, ci.ProductID, ci.Quantity, ci.Subtotal)
		fmt.Println(id)
	}

	return nil
}
