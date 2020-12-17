package checkout

import (
	"errors"
	"fmt"

	"github.com/vitoraalmeida/silkroad/usecase/customer"
	"github.com/vitoraalmeida/silkroad/usecase/delivery"
	"github.com/vitoraalmeida/silkroad/usecase/product"
	"github.com/vitoraalmeida/silkroad/usecase/sale"
	"github.com/vitoraalmeida/silkroad/usecase/saleitem"
)

type Service struct {
	customerService customer.UseCase
	deliveryService delivery.UseCase
	productService  product.UseCase
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
var ErrInvalidProduct = errors.New("Product doesn't exist")

func NewService(s sale.UseCase, si saleitem.UseCase, cs customer.UseCase, ps product.UseCase, ds delivery.UseCase) *Service {
	return &Service{
		customerService: cs,
		productService:  ps,
		saleService:     s,
		saleItemService: si,
		deliveryService: ds,
	}
}

func (s *Service) Checkout(c *Cart, customerID uint) error {
	_, err := s.customerService.GetCustomer(customerID)
	if err != nil {
		return fmt.Errorf("Customer not found")
	}

	// check cart
	if len(*c) == 0 {
		return ErrCartEmpty
	}

	for _, ci := range *c {
		p, err := s.productService.GetProduct(ci.ProductID)
		if err != nil {
			return fmt.Errorf("Item Cart: %#v\nProduct %d not found", ci, ci.ProductID)
		}
		if p.Stock < ci.Quantity {
			return fmt.Errorf("Item Cart: %#v\nProduct %d out of stock", ci, ci.ProductID)
		}
	}

	// calculate cart's total amount
	totalAmount := 0.00
	for _, ci := range *c {
		totalAmount += ci.Subtotal
	}

	// create sale
	saleID, _ := s.saleService.CreateSale(customerID, totalAmount)
	_, err = s.deliveryService.CreateDelivery(saleID, customerID, "Av 2 n 2", "aguardando envio")
	if err != nil {
		return err
	}

	// register sale's items
	for _, ci := range *c {
		err := s.productService.DecrementProductStock(ci.ProductID, ci.Quantity)
		if err != nil {
			return fmt.Errorf("Item Cart: %#v\nProduct %d out of stock", ci, ci.ProductID)
		}
		_, err = s.saleItemService.CreateSaleItem(saleID, ci.ProductID, ci.Quantity, ci.Subtotal)
		if err != nil {
			return err
		}
	}

	return nil
}
