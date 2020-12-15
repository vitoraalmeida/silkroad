package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"

	"github.com/vitoraalmeida/silkroad/entity"
	"github.com/vitoraalmeida/silkroad/infra/repository"
	"github.com/vitoraalmeida/silkroad/usecase/category"
	_ "github.com/vitoraalmeida/silkroad/usecase/checkout"
	"github.com/vitoraalmeida/silkroad/usecase/customer"
	"github.com/vitoraalmeida/silkroad/usecase/product"
	"github.com/vitoraalmeida/silkroad/usecase/sale"
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
	cr := repository.NewCategoryPQSL(db)
	cs := category.NewService(cr)
	// product
	pr := repository.NewProductPQSL(db)
	ps := product.NewService(pr)
	// sale
	sr := repository.NewSalePQSL(db)
	ss := sale.NewService(sr)
	// customer
	csr := repository.NewCustomerPQSL(db)
	css := customer.NewService(csr)
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
	fmt.Println("\n-------------------------- Products ---------------------------------\n")

	fmt.Println(ps.CreateProduct("O Capital Volume 1", 2, 60.00, 10, true))
	fmt.Println(ps.CreateProduct("O Capital Volume 2", 2, 70.00, 10, true))
	fmt.Println(ps.CreateProduct("O Capital Volume 3", 2, 70.00, 50, true))

	fmt.Println(ps.GetProduct(1))

	//fmt.Println(ps.CreateProduct("Capital Inicial M", 2, 30.00, 10, true))

	//fmt.Println("\n------- Search products containing 'capital' ---------------------------------\n")
	//capitals, _ := ps.SearchProducts("Capital")
	//for _, p := range capitals {
	//	fmt.Printf("%+v\n", p)
	//}

	fmt.Println("\n-------------------------- All Products ---------------------------------\n")
	products, err := ps.ListProducts()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, p := range products {
		fmt.Printf("%+v\n", p)
	}

	fmt.Println("\n-------------------------- Updating Product ------------------------------\n")
	capital, _ := ps.GetProduct(1)
	changedCapital := &entity.Product{
		ID:         1,
		Name:       "Das Kapital",
		CategoryID: 2,
		Price:      100.00,
		Stock:      5,
		Available:  true,
		CreatedAt:  capital.CreatedAt,
	}
	fmt.Printf("Updating\n\n%+v\n\nto\n\n %+v\n", capital, changedCapital)

	time.Sleep(5 * time.Second)
	err = ps.UpdateProduct(changedCapital)
	if err != nil {
		fmt.Println("Not changed")
	}

	fmt.Println("\n-------------------------- All Products ---------------------------------\n")
	products, _ = ps.ListProducts()
	for _, p := range products {
		fmt.Printf("%+v\n", p)
	}

	fmt.Println("\n-------------------------- Decrementing product --------------------------------\n")
	fmt.Println(ps.GetProduct(1))
	err = ps.DecrementProductStock(1, 5)
	err = ps.DeleteProduct(2)

	fmt.Println("\n-------------------------- All Products ---------------------------------\n")
	products, _ = ps.ListProducts()
	for _, p := range products {
		fmt.Printf("%+v\n", p)
	}

	fmt.Println("todos que contenham 'apital'")
	products, err = ps.SearchProducts("apital")
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range products {
		fmt.Printf("%v\n", v)
	}

	fmt.Println("\n\n------------------------- Customers --------------------------------------\n")
	fmt.Println(css.CreateCustomer("Vitor", "vitor@mail.com", "66666666666", "asenha"))
	fmt.Println(css.CreateCustomer("Ana", "ana@mail.com", "77777777777", "asenha"))
	fmt.Println(css.CreateCustomer("Fred", "fred@mail.com", "88888888888", "asenha"))

	fmt.Println(css.GetCustomer(1))

	fmt.Println("\n\n------------------------- All Customers --------------------------------------\n")
	customers, _ := css.ListCustomers()
	for _, p := range customers {
		fmt.Printf("%+v\n", p)
	}

	fmt.Println("\n-------------------------- Updating Customer ------------------------------\n")
	customer, _ := css.GetCustomer(1)
	changedCustomer := &entity.Customer{
		ID:        1,
		Name:      "Vitor Almeida",
		Email:     "vitor@gmail.com",
		CPF:       "66666666666",
		Password:  "asenha2",
		CreatedAt: customer.CreatedAt,
	}
	fmt.Printf("Updating\n\n%+v\n\nto\n\n %+v\n", customer, changedCustomer)

	time.Sleep(5 * time.Second)
	err = css.UpdateCustomer(changedCustomer)
	if err != nil {
		fmt.Println("Not changed")
	}

	fmt.Println("\n\n------------------------- All Customers --------------------------------------\n")
	customers, _ = css.ListCustomers()
	for _, p := range customers {
		fmt.Printf("%+v\n", p)
	}

	fmt.Println("\n\n------------------------- Delete Customer --------------------------------------\n")
	err = css.DeleteCustomer(1)

	fmt.Println("\n\n------------------------- All Customers --------------------------------------\n")
	customers, _ = css.ListCustomers()
	for _, p := range customers {
		fmt.Printf("%+v\n", p)
	}

	fmt.Println("\n\n------------------------- Sales --------------------------------------\n")
	fmt.Println(ss.CreateSale(2, 999.99))
	fmt.Println(ss.CreateSale(2, 99.99))
	fmt.Println(ss.CreateSale(3, 666.66))

	fmt.Println(ss.GetSale(1))

	fmt.Println("\n\n------------------------- All Sales --------------------------\n")
	sales, err := ss.ListSales()
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range sales {
		fmt.Printf("%v\n", v)
	}

	fmt.Println("\n\n------------------------- Delete Sale --------------------------------------\n")
	err = ss.DeleteSale(1)

	fmt.Println("\n\n------------------------- Customer 2's sales --------------------------\n")
	sales, err = ss.SearchSales(2)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range sales {
		fmt.Printf("%v\n", v)
	}

	//fmt.Println("\n\n----------------------- Carrinho 1 -----------------------------------\n")

	//ci1 := checkout.CartItem{1, 3, 300.00}
	//ci2 := checkout.CartItem{2, 5, 350.00}
	//ci3 := checkout.CartItem{3, 4, 400.00}
	//cart := &checkout.Cart{ci1, ci2, ci3}
	//fmt.Printf("Cart: %+v", cart)

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
