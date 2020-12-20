package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	_ "time"

	"github.com/gorilla/mux"

	"github.com/vitoraalmeida/silkroad/entity"
	_ "github.com/vitoraalmeida/silkroad/infra/repository"
	_ "github.com/vitoraalmeida/silkroad/usecase/category"
	_ "github.com/vitoraalmeida/silkroad/usecase/checkout"
	_ "github.com/vitoraalmeida/silkroad/usecase/customer"
	_ "github.com/vitoraalmeida/silkroad/usecase/delivery"
	_ "github.com/vitoraalmeida/silkroad/usecase/product"
	_ "github.com/vitoraalmeida/silkroad/usecase/sale"
	_ "github.com/vitoraalmeida/silkroad/usecase/saleitem"
	"github.com/vitoraalmeida/silkroad/views"
)

var (
	homeView    *views.View
	productView *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	product, _ := entity.NewProduct("Loratadina 50mg", 1, 50.00, 5, true)
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, product))
}

func product(w http.ResponseWriter, r *http.Request) {
	product, _ := entity.NewProduct("Loratadina 50mg", 1, 50.00, 5, true)
	w.Header().Set("Content-Type", "text/html")
	must(productView.Render(w, product))
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

	//// category
	//cr := repository.NewCategoryPQSL(db)
	//cs := category.NewService(cr)
	//// product
	//pr := repository.NewProductPQSL(db)
	//ps := product.NewService(pr)
	//// sale
	//sr := repository.NewSalePQSL(db)
	//ss := sale.NewService(sr)
	//// customer
	//csr := repository.NewCustomerPQSL(db)
	//css := customer.NewService(csr)
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

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/product", product)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir("views/static/"))))
	http.ListenAndServe(":3000", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
