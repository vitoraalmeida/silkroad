package main

import (
	"fmt"
	"github.com/vitoraalmeida/silkroad/entity"
	"github.com/vitoraalmeida/silkroad/usecase/category"
	"github.com/vitoraalmeida/silkroad/usecase/product"
)

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

}
