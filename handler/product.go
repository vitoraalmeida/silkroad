package handler

import (
	_ "github.com/vitoraalmeida/silkroad/entity"
	"github.com/vitoraalmeida/silkroad/usecase/category"
	"github.com/vitoraalmeida/silkroad/usecase/product"
	"log"
	"net/http"
	"strconv"
)

type Products struct {
	l  *log.Logger
	ps *product.Service
	cs *category.Service
}

func NewProducts(l *log.Logger, ps *product.Service, cs *category.Service) *Products {
	return &Products{
		l,
		ps,
		cs,
	}
}

func (p *Products) ListProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET products")
	errorMessage := "Error reading accounts"
	data, err := p.ps.ListProducts()
	if err != nil {
		p.l.Println(errorMessage)
		panic(err)
	}
	p.l.Println(data)
}

func (p *Products) CreateProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST products")
	errorMessage := "Error creating product"
	w.Header().Set("Content-type", "text/html")

	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	p.l.Println(r.PostForm)
	name := r.PostForm["name"][0]
	category := r.PostForm["category"][0]
	price := r.PostForm["price"][0]
	stock := r.PostForm["stock"][0]

	priceFloat, err := strconv.ParseFloat(price, 64)
	if err != nil {
		p.l.Println(err)
	}
	stockInt, err := strconv.Atoi(stock)
	inputErr := p.Validate(name, category, priceFloat, uint(stockInt))
	if inputErr != "" {
		p.l.Println(errorMessage)
		p.l.Println(inputErr)
		return
	}

	cat, err := p.cs.SearchCategories(category)
	p.l.Println("passou do search categories")
	p.l.Println(cat)

	id, err := p.ps.CreateProduct(name, cat[0].ID, priceFloat, uint(stockInt), true)
	if err != nil {
		p.l.Println(err)
		return
	}
	p.l.Println("Product created id: ", id)
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
	w.WriteHeader(http.StatusCreated)
}

func (p *Products) Validate(name string, category string, price float64, stock uint) string {
	if name == "" {
		return "Product name must not be empty"
	}
	if category == "" {
		return "category name must not be empty"
	}
	if price <= 0 {
		return "Price must greater than 0"
	}
	if stock <= 0 {
		return "Price must greater than 0"
	}
	return ""
}
