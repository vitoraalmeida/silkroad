package main

import (
	"fmt"
	"github.com/vitoraalmeida/silkroad/entity"
	"github.com/vitoraalmeida/silkroad/usecase/category"
	"github.com/vitoraalmeida/silkroad/usecase/product"
	"github.com/vitoraalmeida/silkroad/usecase/sale"
	"github.com/vitoraalmeida/silkroad/usecase/saleitem"
)

type CartItem struct {
	ProductID uint
	Quantity  uint
	Subtotal  float64
}

type Cart []CartItem

func main() {
	ic := category.NewInmem()
	cs := category.NewService(ic)
	fmt.Println(cs.CreateCategory("livros"))
	fmt.Println(cs.CreateCategory("roupas"))
	categories, err := cs.SearchCategories("roupas")
	if err != nil {
		fmt.Println("Not found")
		return
	}
	for _, v := range categories {
		fmt.Printf("%v\n", v)
	}

	fmt.Println("\n-------------------------- All categories ---------------------------------\n")
	categories, err = cs.ListCategories()
	for _, v := range categories {
		fmt.Printf("%v\n", v)
	}

	// ------------------------------- Produtos -----------------------------------
	fmt.Println("\n-------------------------- Products ---------------------------------\n")
	ip := product.NewInmem()
	ps := product.NewService(ip)

	fmt.Println(ps.CreateProduct("O Capital Volume 1", 1, 60.00, 10, true))
	fmt.Println(ps.CreateProduct("O Capital Volume 2", 1, 70.00, 10, true))

	fmt.Println(ps.GetProduct(1))

	fmt.Println(ps.CreateProduct("Capital Inicial M", 2, 30.00, 5, true))

	fmt.Println("\n------- Search products containing 'capital' ---------------------------------\n")
	capitals, _ := ps.SearchProducts("Capital")
	for _, p := range capitals {
		fmt.Printf("%+v\n", p)
	}

	fmt.Println("\n-------------------------- All Products ---------------------------------\n")
	products, _ := ps.ListProducts()
	for _, p := range products {
		fmt.Printf("%+v\n", p)
	}

	capital, _ := ps.GetProduct(1)
	changedCapital := &entity.Product{
		ID:         1,
		Name:       "Das Kapital",
		CategoryID: 1,
		Price:      100.00,
		Stock:      5,
		Available:  true,
		CreatedAt:  capital.CreatedAt,
	}
	fmt.Printf("\n----- Updating Product\n %+v\n\n to\n\n %+v\n", capital, changedCapital)

	err = ps.UpdateProduct(changedCapital)
	if err != nil {
		fmt.Println("Not changed")
	}

	fmt.Println("\n-------------------------- All Products ---------------------------------\n")
	products, _ = ps.ListProducts()
	for _, p := range products {
		fmt.Printf("%+v\n", p)
	}

	fmt.Println("\n\nLista de vendas ------------------------------------\n")

	inmemSale := sale.NewInmem()
	ss := sale.NewService(inmemSale)
	saleId, err := ss.CreateSale(1, 500.00)
	if err != nil {
		fmt.Println("Could not create Sale")
		return
	}
	fmt.Println("SaleId: ", saleId)
	saleId, err = ss.CreateSale(2, 500.00)
	if err != nil {
		fmt.Println("Could not create Sale")
		return
	}
	fmt.Println("SaleId: ", saleId)

	sales, _ := ss.ListSales()
	for _, s := range sales {
		fmt.Printf("%+v\n", s)
	}

	fmt.Println("\n\nLista itens de venda ------------------------------------\n")
	fmt.Println("\n\nCarrinho -----------------------------------\n")

	ci1 := CartItem{1, 3, 300.00}
	ci2 := CartItem{2, 5, 350.00}
	ci3 := CartItem{3, 4, 400.00}
	cart := &Cart{ci1, ci2, ci3}
	fmt.Printf("Cart: %+v", cart)
	inmemSaleItem := saleitem.NewInmem()
	sis := saleitem.NewService(inmemSaleItem)

	totalAmount := 0.00
	for _, ci := range *cart {
		totalAmount += ci.Subtotal
	}
	fmt.Println("\nTotal price: ", totalAmount)

	saleId, err = ss.CreateSale(1, totalAmount)
	if err != nil {
		fmt.Println("Could not create Sale")
		return
	}
	fmt.Println("SaleId: ", saleId)

	for _, ci := range *cart {
		id, _ := sis.CreateSaleItem(saleId, ci.ProductID, ci.Quantity, ci.Subtotal)
		fmt.Println(id)
	}

	fmt.Println("SaleItems: \n")
	items, _ := sis.ListSaleItems()
	for _, i := range items {
		fmt.Printf("%+v\n", i)
	}

	fmt.Println("\nSales: \n")
	sales, _ = ss.ListSales()
	for _, s := range sales {
		fmt.Printf("%+v\n", s)
	}

}