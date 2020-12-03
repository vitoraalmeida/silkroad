package main

import (
	"fmt"
	"github.com/vitoraalmeida/silkroad/usecase/category"
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

	fmt.Printf("\n-------------------------- All categories ---------------------------------\n")
	categories, err = cs.ListCategories()
	for _, v := range categories {
		fmt.Printf("%v\n", v)
	}

}
