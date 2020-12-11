package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"

	_ "github.com/vitoraalmeida/silkroad/entity"
	"github.com/vitoraalmeida/silkroad/infra/repository"
	"github.com/vitoraalmeida/silkroad/usecase/category"
	_ "github.com/vitoraalmeida/silkroad/usecase/checkout"
	_ "github.com/vitoraalmeida/silkroad/usecase/customer"
	_ "github.com/vitoraalmeida/silkroad/usecase/product"
	_ "github.com/vitoraalmeida/silkroad/usecase/sale"
	_ "github.com/vitoraalmeida/silkroad/usecase/saleitem"
)

type CartItem struct {
	ProductID uint
	Quantity  uint
	Subtotal  float64
}

type Cart []CartItem

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env: %v", err)
	} else {
		log.Println("Getting the env values...")
	}

	dsn := fmt.Sprintf(
		`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"), os.Getenv("DB_NAME"),
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	// category
	//ic := category.NewInmem()
	cr := repository.NewCategoryPQSL(db)
	cs := category.NewService(cr)
	// product
	//ip := product.NewInmem()
	//ps := product.NewService(ip)
	//// sale
	//inmemSale := sale.NewInmem()
	//ss := sale.NewService(inmemSale)
	//// customer
	//inmemCustomer := customer.NewInmem()
	//css := customer.NewService(inmemCustomer)
	//// saleitem
	//inmemSaleItem := saleitem.NewInmem()
	//sis := saleitem.NewService(inmemSaleItem)
	//// checkout
	//chs := checkout.NewService(ss, sis, css, ps)

	fmt.Println(cs.CreateCategory("livros"))
	fmt.Println(cs.CreateCategory("roupas"))
	fmt.Println(cs.CreateCategory("livros capa dura"))
	fmt.Println(cs.CreateCategory("Tipo Livros"))
	category, err := cs.GetCategory(1)
	if err != nil {
		fmt.Println("Not found")
		return
	}
	fmt.Printf("%v\n", category)

	//fmt.Println("\n-------------------------- All categories ---------------------------------\n")
	categories, err := cs.ListCategories()
	if err != nil {
		fmt.Errorf("%v", err)
		return
	}
	for _, v := range categories {
		fmt.Printf("%v\n", v)
	}

	c1, err := cs.GetCategory(1)
	fmt.Printf("Atualizando: %+v\n", c1)
	c1.Name = "Livro"
	fmt.Printf("Para: %+v\n", c1)
	err = cs.UpdateCategory(c1)
	if err != nil {
		fmt.Println(err)
	}
	c1, err = cs.GetCategory(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", c1)

	err = cs.DeleteCategory(1)
	if err != nil {
		fmt.Println(err)
	}

	categories, err = cs.ListCategories()
	if err != nil {
		fmt.Errorf("%v", err)
		return
	}
	fmt.Println("depois de deletar")
	for _, v := range categories {
		fmt.Printf("%v\n", v)
	}
	fmt.Println("todas que contenham 'ivros'")
	categories, err = cs.SearchCategories("livros")
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range categories {
		fmt.Printf("%v\n", v)
	}

	//// ------------------------------- Produtos -----------------------------------
	//fmt.Println("\n-------------------------- Products ---------------------------------\n")

	//fmt.Println(ps.CreateProduct("O Capital Volume 1", 1, 60.00, 10, true))
	//fmt.Println(ps.CreateProduct("O Capital Volume 2", 1, 70.00, 10, true))

	//fmt.Println(ps.GetProduct(1))

	//fmt.Println(ps.CreateProduct("Capital Inicial M", 2, 30.00, 10, true))

	//fmt.Println("\n------- Search products containing 'capital' ---------------------------------\n")
	//capitals, _ := ps.SearchProducts("Capital")
	//for _, p := range capitals {
	//	fmt.Printf("%+v\n", p)
	//}

	//fmt.Println("\n-------------------------- All Products ---------------------------------\n")
	//products, _ := ps.ListProducts()
	//for _, p := range products {
	//	fmt.Printf("%+v\n", p)
	//}

	//capital, _ := ps.GetProduct(1)
	//changedCapital := &entity.Product{
	//	ID:         1,
	//	Name:       "Das Kapital",
	//	CategoryID: 1,
	//	Price:      100.00,
	//	Stock:      5,
	//	Available:  true,
	//	CreatedAt:  capital.CreatedAt,
	//}
	//fmt.Printf("\n----- Updating Product\n %+v\n\n to\n\n %+v\n", capital, changedCapital)

	//err = ps.UpdateProduct(changedCapital)
	//if err != nil {
	//	fmt.Println("Not changed")
	//}

	//fmt.Println("\n-------------------------- All Products ---------------------------------\n")
	//products, _ = ps.ListProducts()
	//for _, p := range products {
	//	fmt.Printf("%+v\n", p)
	//}

	//fmt.Println("\n\n----------------------- Lista de vendas ------------------------------------\n")

	//sales, _ := ss.ListSales()
	//for _, s := range sales {
	//	fmt.Printf("%+v\n", s)
	//}

	//fmt.Println("\n\n----------------------- Carrinho 1 -----------------------------------\n")

	//ci1 := checkout.CartItem{1, 3, 300.00}
	//ci2 := checkout.CartItem{2, 5, 350.00}
	//ci3 := checkout.CartItem{3, 4, 400.00}
	//cart := &checkout.Cart{ci1, ci2, ci3}
	//fmt.Printf("Cart: %+v", cart)

	//fmt.Println("\n\nCustomer 1 ------------------------------------\n")

	//customerID, _ := css.CreateCustomer("Vitor Almeida", "vitor@mail.com", "01234567890", "123456")
	//customer1, err := css.GetCustomer(customerID)
	//fmt.Printf("customer: %+v", customer1)

	//fmt.Println("\n\n----------------------- Checkout 1 ----------------------------------------------")
	//chs.Checkout(cart, 1)

	//fmt.Println("--------------------------- Sale Items ----------------------------------- \n")
	//items, _ := sis.ListSaleItems()
	//for _, i := range items {
	//	fmt.Printf("%+v\n", i)
	//}

	//fmt.Println("\n------------------------- All Sales ------------------------------------ \n")
	//sales, _ = ss.ListSales()
	//for _, s := range sales {
	//	fmt.Printf("%+v\n", s)
	//}

	//fmt.Println("\n---------------------- Sale1's items ----------------- \n")
	//sale3Items, _ := sis.SearchItemsBySaleID(1)
	//for _, si := range sale3Items {
	//	fmt.Printf("%+v\n", si)
	//}

	//fmt.Println("\n\n---------------------- Carrinho 2 -----------------------------------\n")

	//ci1 = checkout.CartItem{1, 3, 300.00}
	//ci2 = checkout.CartItem{2, 5, 350.00}
	//ci3 = checkout.CartItem{3, 4, 120.00}
	//ci4 := checkout.CartItem{3, 2, 200.00}
	//cart = &checkout.Cart{ci1, ci2, ci3, ci4}
	//fmt.Printf("Cart: %+v\n", cart)

	//err = chs.Checkout(cart, 2)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = chs.Checkout(cart, 1)
	//if err != nil {
	//	fmt.Println(err)
	//}

	//fmt.Println("\n-------------------------- All Products ---------------------------------\n")
	//products, _ = ps.ListProducts()
	//for _, p := range products {
	//	fmt.Printf("%+v\n", p)
	//}

	//fmt.Println("\n ----------------------- All Sale Items: \n")
	//items, _ = sis.ListSaleItems()
	//for _, i := range items {
	//	fmt.Printf("%+v\n", i)
	//}

	//fmt.Println("\nSales: \n")
	//sales, _ = ss.ListSales()
	//for _, s := range sales {
	//	fmt.Printf("%+v\n", s)
	//}

}
