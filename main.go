package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"strconv"
	_ "time"

	"github.com/gorilla/mux"

	"github.com/vitoraalmeida/silkroad/entity"
	"github.com/vitoraalmeida/silkroad/handler"

	"github.com/vitoraalmeida/silkroad/infra/repository"
	"github.com/vitoraalmeida/silkroad/usecase/category"
	_ "github.com/vitoraalmeida/silkroad/usecase/checkout"
	"github.com/vitoraalmeida/silkroad/usecase/customer"
	_ "github.com/vitoraalmeida/silkroad/usecase/delivery"
	"github.com/vitoraalmeida/silkroad/usecase/product"
	_ "github.com/vitoraalmeida/silkroad/usecase/sale"
	_ "github.com/vitoraalmeida/silkroad/usecase/saleitem"
	"github.com/vitoraalmeida/silkroad/views"
)

var (
	productService     *product.Service
	products           *[]entity.Product
	homeView           *views.View
	adminView          *views.View
	productView        *views.View
	editProductView    *views.View
	createProductView  *views.View
	createCategoryView *views.View
	signupView         *views.View
	signinView         *views.View
)

func admin(w http.ResponseWriter, r *http.Request) {
	var prods []entity.Product
	products, _ := productService.ListProducts()
	for _, v := range products {
		prods = append(prods, *v)
	}
	w.Header().Set("Content-Type", "text/html")
	must(adminView.Render(w, prods))
}

func home(w http.ResponseWriter, r *http.Request) {
	var prods []entity.Product
	products, _ := productService.ListProducts()
	for _, v := range products {
		prods = append(prods, *v)
	}
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, prods))
}

func editProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
	}
	p, _ := productService.GetProduct(uint(id))
	w.Header().Set("Content-Type", "text/html")
	must(editProductView.Render(w, p))
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(createProductView.Render(w, nil))
}

func createCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(createCategoryView.Render(w, nil))
}

func seeProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
	}
	p, _ := productService.GetProduct(uint(id))
	w.Header().Set("Content-Type", "text/html")
	must(productView.Render(w, p))
}

func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(signupView.Render(w, nil))
}

func signin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(signinView.Render(w, nil))
}

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
	l := log.New(os.Stdout, "silkroad", log.LstdFlags)

	//// category
	cr := repository.NewCategoryPQSL(db)
	cs := category.NewService(cr)
	ch := handler.NewCategories(l, cs)
	// product
	pr := repository.NewProductPQSL(db)
	ps := product.NewService(pr)
	productService = ps
	psh := handler.NewProducts(l, ps, cs)
	//// sale
	//sr := repository.NewSalePQSL(db)
	//ss := sale.NewService(sr)
	// customer
	csr := repository.NewCustomerPQSL(db)
	css := customer.NewService(csr)
	csh := handler.NewCustomers(l, css)

	// signin handler
	sih := handler.NewSignIn(l, css)

	//// saleitem
	//sir := repository.NewSaleItemPQSL(db)
	//sis := saleitem.NewService(sir)
	//// delivery
	//dr := repository.NewDeliveryPSQL(db)
	//ds := delivery.NewService(dr)
	//// checkout
	//chs := checkout.NewService(ss, sis, css, ps, ds)

	homeView = views.NewView("main", "views/home.tmpl")
	productView = views.NewView("main", "views/product.tmpl")
	signupView = views.NewView("main", "views/signup.tmpl")
	signinView = views.NewView("main", "views/signin.tmpl")
	adminView = views.NewView("main-admin", "views/home-admin.tmpl")
	editProductView = views.NewView("main-admin", "views/edit-product.tmpl")
	createProductView = views.NewView("main-admin", "views/create-product.tmpl")
	createCategoryView = views.NewView("main-admin", "views/create-category.tmpl")

	r := mux.NewRouter()
	getRouter := r.Methods("GET").Subrouter()
	getRouter.HandleFunc("/", home)
	getRouter.HandleFunc("/admin", admin)
	getRouter.HandleFunc("/admin/product/{id:[0-9]+}/update", editProduct)
	getRouter.HandleFunc("/admin/product/create", createProduct)
	getRouter.HandleFunc("/admin/category/create", createCategory)
	//getRouter.HandleFunc("/admin/products", ListProducts)
	//getRouter.HandleFunc("/admin/category", ListCategories)
	//getRouter.HandleFunc("/admin/sales", ListSales)
	getRouter.HandleFunc("/product/{id:[0-9]+}", seeProduct)
	getRouter.HandleFunc("/signup", signup)
	getRouter.HandleFunc("/signin", signin)

	postRouter := r.Methods("POST").Subrouter()
	postRouter.HandleFunc("/signup", csh.CreateCustomer)
	postRouter.HandleFunc("/signin", sih.SignIn)
	postRouter.HandleFunc("/admin/product/create", psh.CreateProduct)
	postRouter.HandleFunc("/admin/category/create", ch.CreateCategory)
	//postRouter.HandleFunc("/admin/products", CreateProduct)
	//postRouter.HandleFunc("/admin/category", CreateCategory)
	//postRouter.HandleFunc("/sale", CreateSale)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir("views/static/"))))
	http.ListenAndServe(":3000", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
