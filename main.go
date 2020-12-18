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

	_ "github.com/vitoraalmeida/silkroad/entity"
	_ "github.com/vitoraalmeida/silkroad/infra/repository"
	_ "github.com/vitoraalmeida/silkroad/usecase/category"
	_ "github.com/vitoraalmeida/silkroad/usecase/checkout"
	_ "github.com/vitoraalmeida/silkroad/usecase/customer"
	_ "github.com/vitoraalmeida/silkroad/usecase/delivery"
	_ "github.com/vitoraalmeida/silkroad/usecase/product"
	_ "github.com/vitoraalmeida/silkroad/usecase/sale"
	_ "github.com/vitoraalmeida/silkroad/usecase/saleitem"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func admin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, please send an email "+
		"to <a href=\"mailto:support@lenslocked.com\">"+
		"support@lenslocked.com</a>.")
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

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/admin", admin)
	http.ListenAndServe(":3000", r)
}
