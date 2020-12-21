package handler

import (
	_ "github.com/vitoraalmeida/silkroad/entity"
	"github.com/vitoraalmeida/silkroad/usecase/category"
	"log"
	"net/http"
)

type Categories struct {
	l  *log.Logger
	cs *category.Service
}

func NewCategories(l *log.Logger, cs *category.Service) *Categories {
	return &Categories{
		l,
		cs,
	}
}

func (c *Categories) ListCategories(w http.ResponseWriter, r *http.Request) {
	c.l.Println("Handle GET categories")
	errorMessage := "Error reading categories"
	data, err := c.cs.ListCategories()
	if err != nil {
		c.l.Println(errorMessage)
		panic(err)
	}
	c.l.Println(data)
}

func (c *Categories) CreateCategory(w http.ResponseWriter, r *http.Request) {
	c.l.Println("Handle POST categories")
	errorMessage := "Error creating category"
	w.Header().Set("Content-type", "text/html")

	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	c.l.Println(r.PostForm)
	name := r.PostForm["name"][0]

	inputErr := c.Validate(name)
	if inputErr != "" {
		c.l.Println(errorMessage)
		c.l.Println(inputErr)
		return
	}

	id, err := c.cs.CreateCategory(name)
	if err != nil {
		c.l.Println(err)
		return
	}

	c.l.Println("category created id: ", id)
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (c *Categories) Validate(name string) string {
	if name == "" {
		return "Category name must not be empty"
	}
	return ""
}
